package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/IBM/sarama"
	_ "github.com/go-sql-driver/mysql"
)

// Redis 配置
var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// Kafka 配置
var kafkaBrokers = []string{"localhost:9092"}
var kafkaTopic = "order_topic"

// MySQL 配置
var db *sql.DB

// 令牌桶限流
type TokenBucket struct {
	rate   int
	bucket int
	mu     sync.Mutex
	ticker *time.Ticker
}

func NewTokenBucket(rate int) *TokenBucket {
	tb := &TokenBucket{
		rate:   rate,
		bucket: rate,
		ticker: time.NewTicker(time.Second),
	}
	go tb.refill()
	return tb
}

func (tb *TokenBucket) refill() {
	for range tb.ticker.C {
		tb.mu.Lock()
		tb.bucket = tb.rate
		tb.mu.Unlock()
	}
}

func (tb *TokenBucket) TryAcquire() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	if tb.bucket > 0 {
		tb.bucket--
		return true
	}
	return false
}

var tb = NewTokenBucket(100) // 每秒最多 100 个请求

// 预加载库存
func LoadStockToRedis(productID string, stock int) {
	err := rdb.Set(ctx, "stock:"+productID, stock, 0).Err()
	if err != nil {
		log.Fatalf("Failed to load stock: %v", err)
	}
}

// 扣减 Redis 库存
func DeductStock(productID string) bool {
	stock, err := rdb.Decr(ctx, "stock:"+productID).Result()
	if err != nil {
		log.Printf("Decr error: %v", err)
		return false
	}
	if stock < 0 { // 防止超卖
		rdb.Incr(ctx, "stock:"+productID)
		return false
	}
	return true
}

// 发送订单到 Kafka
func SendOrderToKafka(userID, productID string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(kafkaBrokers, config)
	if err != nil {
		log.Println("Failed to create Kafka producer:", err)
		return
	}
	defer producer.Close()

	message := fmt.Sprintf("%s:%s", userID, productID)
	msg := &sarama.ProducerMessage{
		Topic: kafkaTopic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Println("Failed to send order to Kafka:", err)
	}
}

// 处理秒杀请求
func seckillHandler(w http.ResponseWriter, r *http.Request) {
	if !tb.TryAcquire() {
		http.Error(w, "请求过多，限流", http.StatusTooManyRequests)
		return
	}

	userID := r.URL.Query().Get("userID")
	productID := r.URL.Query().Get("productID")
	if userID == "" || productID == "" {
		http.Error(w, "参数错误", http.StatusBadRequest)
		return
	}

	// Redis 预扣库存
	if !DeductStock(productID) {
		http.Error(w, "库存不足", http.StatusGone)
		return
	}

	// 发送订单到 Kafka
	SendOrderToKafka(userID, productID)

	w.Write([]byte("秒杀成功，订单处理中"))
}

// 订单消费者，监听 Kafka 并存储到数据库
func consumeOrders() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(kafkaBrokers, config)
	if err != nil {
		log.Fatal("Failed to create Kafka consumer:", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Failed to consume Kafka partition:", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		orderData := string(msg.Value)
		parts := strings.Split(orderData, ":")
		if len(parts) != 2 {
			log.Println("Invalid order format:", orderData)
			continue
		}
		userID, productID := parts[0], parts[1]

		// 处理订单
		SaveOrder(userID, productID)
	}
}

// 将订单存入 MySQL
func SaveOrder(userID, productID string) {
	_, err := db.Exec("INSERT INTO orders (user_id, product_id) VALUES (?, ?)", userID, productID)
	if err != nil {
		log.Println("Failed to save order:", err)
	} else {
		log.Printf("订单已保存：用户 %s 购买了商品 %s\n", userID, productID)
	}
}

// 初始化 MySQL 连接
func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/seckill")
	if err != nil {
		log.Fatal(err)
	}

	// 创建订单表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id VARCHAR(50),
		product_id VARCHAR(50),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	LoadStockToRedis("1001", 10) // 商品 ID 1001，库存 10 个

	// 启动 Kafka 消费者
	go consumeOrders()

	http.HandleFunc("/seckill", seckillHandler)
	log.Println("秒杀系统启动，访问 http://localhost:8080/seckill?userID=1&productID=1001")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
