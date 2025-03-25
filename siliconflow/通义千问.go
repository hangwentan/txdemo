package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func deepSeek() {
	url := "https://api.siliconflow.cn/v1/chat/completions"

	payload := strings.NewReader("{\n  \"model\": \"Qwen/QwQ-32B\",\n  \"messages\": [\n    {\n      \"role\": \"user\",\n      \"content\": \"go语言常用的包工具有哪些\"\n    }\n  ],\n  \"stream\": false,\n  \"max_tokens\": 512,\n  \"stop\": null,\n  \"temperature\": 0.7,\n  \"top_p\": 0.7,\n  \"top_k\": 50,\n  \"frequency_penalty\": 0.5,\n  \"n\": 1,\n  \"response_format\": {\n    \"type\": \"text\"\n  },\n  \"tools\": [\n    {\n      \"type\": \"function\",\n      \"function\": {\n        \"description\": \"<string>\",\n        \"name\": \"<string>\",\n        \"parameters\": {},\n        \"strict\": false\n      }\n    }\n  ]\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "Bearer <token>")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func main() {
	deepSeek()
}
