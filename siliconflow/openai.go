package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ChatMessage 定义了消息格式
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIRequest 定义了请求 OpenAI API 的请求体结构
type OpenAIRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

// OpenAIResponse 定义了从 OpenAI API 返回的数据结构
type OpenAIResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

// ChatLog 用于记录聊天记录（后端搭建 - 第3部分）
type ChatLog struct {
	ID        uint      `gorm:"primaryKey"`
	UserMsg   string    // 用户输入
	Reply     string    // 系统回复
	CreatedAt time.Time // 时间戳
}

// FAQItem 定义 FAQ 知识库中的问题与答案（增强能力 - 第4部分）
type FAQItem struct {
	Question string
	Answer   string
}

// 模拟 FAQ 知识库
var faqDB = []FAQItem{
	{
		Question: "如何重置密码",
		Answer:   "您可以点击登录页面的“忘记密码”链接，然后按照提示重置您的密码。",
	},
	{
		Question: "如何查询订单状态",
		Answer:   "请在订单详情页面查看订单状态，或联系在线客服获取帮助。",
	},
}

// searchFAQ 检索 FAQ 知识库，简单使用字符串匹配
func searchFAQ(query string) (string, bool) {
	query = strings.ToLower(query)
	for _, item := range faqDB {
		if strings.Contains(strings.ToLower(item.Question), query) {
			return item.Answer, true
		}
	}
	return "", false
}

var db *gorm.DB

// initDB 初始化数据库（这里使用 SQLite 进行示例）
func initDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("chatlog.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	// 自动迁移 ChatLog 表
	return db.AutoMigrate(&ChatLog{})
}

// chatHandler 处理 /chat 请求
func chatHandler(c *gin.Context) {
	// 定义请求结构，接收前端传来的消息
	var req struct {
		Message string `json:"message"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式错误"})
		return
	}

	// 先尝试在 FAQ 知识库中查找答案（增强能力 - 第4部分）
	if answer, found := searchFAQ(req.Message); found {
		// 记录聊天日志
		chatLog := ChatLog{
			UserMsg: req.Message,
			Reply:   answer,
		}
		db.Create(&chatLog)
		c.JSON(http.StatusOK, gin.H{"reply": answer, "source": "FAQ知识库"})
		return
	}

	// 从环境变量中获取 OpenAI API Key
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "缺少 OPENAI_API_KEY"})
		return
	}

	// 构造 OpenAI API 请求体
	openaiRequest := OpenAIRequest{
		Model: "gpt-3.5-turbo", // 可根据需要选择其他模型
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: req.Message,
			},
		},
	}

	requestBody, err := json.Marshal(openaiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求序列化失败"})
		return
	}

	// 调用 OpenAI API
	client := &http.Client{}
	reqOpenAI, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建 OpenAI 请求失败"})
		return
	}
	reqOpenAI.Header.Set("Content-Type", "application/json")
	reqOpenAI.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(reqOpenAI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发送 OpenAI 请求失败"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取 OpenAI 响应失败"})
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": fmt.Sprintf("OpenAI API 返回错误: %s", string(body))})
		return
	}

	// 解析 OpenAI API 返回的数据
	var openaiResponse OpenAIResponse
	if err := json.Unmarshal(body, &openaiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解析 OpenAI 响应失败"})
		return
	}

	if len(openaiResponse.Choices) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未收到 OpenAI 回复"})
		return
	}

	reply := openaiResponse.Choices[0].Message.Content

	// 记录聊天日志（后端搭建 - 第3部分）
	chatLog := ChatLog{
		UserMsg: req.Message,
		Reply:   reply,
	}
	db.Create(&chatLog)

	// 返回回复给前端
	c.JSON(http.StatusOK, gin.H{"reply": reply, "source": "OpenAI"})
}

func main() {
	// 初始化数据库
	if err := initDB(); err != nil {
		panic(fmt.Sprintf("数据库初始化失败: %v", err))
	}

	// 创建 Gin 路由引擎
	r := gin.Default()
	// 定义 /chat 路由
	r.POST("/chat", chatHandler)
	// 启动服务（监听 8080 端口）
	r.Run(":8080")
}
