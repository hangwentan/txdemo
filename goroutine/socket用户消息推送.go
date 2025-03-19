package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Message 定义了消息的结构体
type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

// Client 表示一个连接的客户端
type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
	userID string
}

// Hub 用于管理所有客户端连接，并负责消息转发
type Hub struct {
	// clients 存储所有连接，key 为用户ID
	clients map[string]*Client
	mu      sync.RWMutex
	// broadcast 用于接收所有需要转发的消息
	broadcast chan Message
}

func newHub() *Hub {
	return &Hub{
		clients:   make(map[string]*Client),
		broadcast: make(chan Message),
	}
}

// run 在后台循环处理消息转发
func (h *Hub) run() {
	for {
		msg := <-h.broadcast
		h.mu.RLock()
		if client, ok := h.clients[msg.To]; ok {
			// 将消息转换为字节
			b, err := json.Marshal(msg)
			if err != nil {
				log.Println("json marshal error:", err)
				h.mu.RUnlock()
				continue
			}
			// 将消息发送到目标客户端的发送队列中
			select {
			case client.send <- b:
			default:
				log.Println("client send channel full, dropping message")
			}
		} else {
			log.Printf("目标用户 %s 不在线\n", msg.To)
		}
		h.mu.RUnlock()
	}
}

// readPump 负责读取客户端发送的消息，并将其转发到 Hub
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister(c)
		c.conn.Close()
	}()
	// 设置读取限制和心跳检测
	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("unmarshal error:", err)
			continue
		}
		// 自动设置消息的发送者为当前客户端
		msg.From = c.userID
		// 将消息交由 Hub 转发
		c.hub.broadcast <- msg
	}
}

// writePump 负责将 Hub 推送的消息写入到客户端连接中
func (c *Client) writePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				// 当发送通道被关闭，退出连接
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("write error:", err)
				return
			}
		case <-ticker.C:
			// 发送 ping 保持连接活跃
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// register 将新连接注册到 Hub 中
func (h *Hub) register(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[client.userID] = client
}

// unregister 将断开的连接从 Hub 中移除
func (h *Hub) unregister(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.clients[client.userID]; ok {
		delete(h.clients, client.userID)
		close(client.send)
	}
}

var upgrader1 = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有来源，生产环境下请根据需求进行限制
	CheckOrigin: func(r *http.Request) bool { return true },
}

// serveWs 处理 WebSocket 连接请求
// 客户端连接时需要在 URL 中传递自己的 user 参数，例如：/ws?user=user1
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")
	if userID == "" {
		http.Error(w, "缺少 user 参数", http.StatusBadRequest)
		return
	}

	conn, err := upgrader1.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	client := &Client{
		hub:    hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		userID: userID,
	}
	hub.register(client)

	// 分别启动读和写的 goroutine
	go client.writePump()
	go client.readPump()
}

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Println("服务器启动，监听 :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
