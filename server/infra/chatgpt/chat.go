package chatgpt

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/https"
)

type AiMessage struct {
	Role string `json:"role"`
	Msg  []*ChatMessage
}

type ChatGPT interface {
	Chat(request *ChatRequest) (*ChatResponse, error)
}

type AIChatGPT struct {
	ChatGPT
	Url    string `json:"url"`
	ApiKey string `json:"apiKey"`
	Model  string `json:"model"`
}

func NewAIChatGPT() *AIChatGPT {
	return &AIChatGPT{
		Url:    "https://api.openai.com/v1/chat/completions",
		ApiKey: global.CONFIG.ChatGPT.ApiKey,
		Model:  "gpt-3.5-turbo",
	}
}

func (s *AIChatGPT) ReadModelJSON(filepath string) ([]*ChatRole, error) {
	// 读取 JSON 文件内容
	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Failed to read JSON file:", err)
		return nil, err
	}

	// 解析 JSON 数据到结构体
	var roles []*ChatRole
	err = json.Unmarshal(jsonData, &roles)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return nil, err
	}

	return roles, nil
}

func (s *AIChatGPT) GetRole(act string) (string, error) {
	roles, err := s.ReadModelJSON("./prompts-zh.json")
	if err != nil {
		return "", err
	}

	for _, role := range roles {
		if role.Act == act {
			return role.Prompt, nil
		}
	}

	return "", fmt.Errorf("not found role")
}

func (s *AIChatGPT) CosRole(act string) (resp *ChatResponse, err error) {

	prompt, err := s.GetRole(act)
	if err != nil {
		return nil, err
	}

	content := ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: append([]*ChatMessage{{
			Role:    "user",
			Content: prompt,
		}}),
	}

	res, code := https.NewHttpBuilder(s.Url).
		AddHeader("Content-Type", "application/json").
		AddHeader("Authorization", "Bearer "+s.ApiKey).
		AddBody(content).
		Post()

	global.LOG.Println(code)
	global.LOG.Println(res)

	resp = &ChatResponse{}
	err = json.Unmarshal([]byte(res), resp)
	if err != nil {
		global.LOG.Println(err)
		return nil, err
	}

	return resp, nil
}

func (s *AIChatGPT) Chat(req []*ChatMessage) (resp *ChatResponse, err error) {
	content := ChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: req,
	}

	res, code := https.NewHttpBuilder(s.Url).
		AddHeader("Content-Type", "application/json").
		AddHeader("Authorization", "Bearer "+s.ApiKey).
		AddBody(content).
		Post()

	global.LOG.Println(code)
	global.LOG.Println(res)

	resp = &ChatResponse{}
	err = json.Unmarshal([]byte(res), resp)
	if err != nil {
		global.LOG.Println(err)
		return nil, err
	}

	return resp, nil
}
