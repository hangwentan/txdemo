package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
	"time"
)

const queueName = "task_queue"

var ctx = context.Background()

// 创建 Redis 客户端
func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如有密码，填入
		DB:       0,
	})
}

// 生产者：将消息推送到队列
func producer(client *redis.Client, message string) {
	err := client.LPush(ctx, queueName, message).Err()
	if err != nil {
		log.Fatalf("生产者推送失败: %v", err)
	}
	fmt.Printf("生产者推送消息: %s\n", message)
}

// 消费者：阻塞式从队列取出消息
func consumer(client *redis.Client) {
	for {
		// BLPOP 阻塞等待队列中的消息（避免空轮询浪费 CPU）
		result, err := client.BLPop(ctx, 5*time.Second, queueName).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Println("队列为空，等待消息...")
				continue
			}
			log.Fatalf("消费者消费失败: %v", err)
		}
		fmt.Printf("消费者处理消息: %s\n", result[1])
	}
}

func main() {
	client := newRedisClient()
	defer client.Close()

	// 生产者示例
	go func() {
		for i := 1; i <= 5; i++ {
			producer(client, fmt.Sprintf("任务-%d", i))
			time.Sleep(1 * time.Second)
		}
	}()

	// 消费者示例
	consumer(client)
}
