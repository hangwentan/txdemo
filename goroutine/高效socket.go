package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket 升级器
var upgrader2 = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有跨域请求
	},
}

// 连接管理
type Connection struct {
	Conn *websocket.Conn
	Send chan []byte // 发送消息通道
}

var (
	clients   sync.Map            // 存储所有 WebSocket 连接
	broadcast = make(chan []byte) // 消息广播通道
)

// 处理 WebSocket 连接
func handleConnection(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 WebSocket
	conn, err := upgrader2.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	clientID := fmt.Sprintf("%p", conn) // 以连接地址作为 ID
	fmt.Println("clientID:", clientID)
	client := &Connection{Conn: conn, Send: make(chan []byte, 256)}
	clients.Store(clientID, client)

	// 监听读取消息
	go readMessages(client, clientID)

	// 监听发送消息
	go writeMessages(client, clientID)

	// 处理心跳检测
	go heartbeatCheck(client, clientID)
}

// 读取客户端消息
func readMessages(client *Connection, clientID string) {
	defer func() {
		clients.Delete(clientID)
		client.Conn.Close()
	}()

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", msg)

		// 将消息发送到广播通道
		broadcast <- msg
	}
}

// 发送消息
func writeMessages(client *Connection, clientID string) {
	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write error:", err)
			clients.Delete(clientID)
			client.Conn.Close()
			break
		}
	}
}

// 监听广播消息并转发给所有客户端
func broadcastMessages() {
	for {
		msg := <-broadcast
		clients.Range(func(key, value interface{}) bool {
			client := value.(*Connection)
			select {
			case client.Send <- msg:
			default:
				log.Println("Send channel full, closing connection:", key)
				clients.Delete(key)
				client.Conn.Close()
			}
			return true
		})
	}
}

// 心跳检测
func heartbeatCheck(client *Connection, clientID string) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			log.Println("Ping failed, closing connection:", clientID)
			clients.Delete(clientID)
			client.Conn.Close()
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)

	go broadcastMessages() // 启动广播协程

	log.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
