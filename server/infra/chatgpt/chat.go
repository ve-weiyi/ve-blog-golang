package chatgpt

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/httpx"
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
func (s *AIChatGPT) Chat(req []*ChatMessage) (resp *ChatResponse, err error) {
	content := ChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: req,
	}

	res, err := httpx.NewClient(
		httpx.WithHeader("Content-Type", "application/json"),
		httpx.WithHeader("Authorization", "Bearer "+s.ApiKey),
		httpx.WithBodyObject(content),
	).DoRequest("POST", s.Url)
	if err != nil {
		return nil, err
	}

	resp = &ChatResponse{}
	err = json.Unmarshal([]byte(res), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AIChatGPT) CosRole(act string) (resp *ChatResponse, err error) {

	prompt, err := s.getRole(act)
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

	res, err := httpx.NewClient(
		httpx.WithHeader("Content-Type", "application/json"),
		httpx.WithHeader("Authorization", "Bearer "+s.ApiKey),
		httpx.WithBodyObject(content),
	).DoRequest("POST", s.Url)
	if err != nil {
		return nil, err
	}

	resp = &ChatResponse{}
	err = json.Unmarshal([]byte(res), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AIChatGPT) getRole(act string) (string, error) {
	roles, err := s.readModelJSON("./prompts-zh.json")
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

func (s *AIChatGPT) readModelJSON(filepath string) ([]*ChatRole, error) {
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
