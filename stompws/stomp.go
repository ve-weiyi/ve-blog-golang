package stompws

import (
	"fmt"
	"strings"
)

var (
	PongStomp = StompFrame{
		Command: "",
		Headers: nil,
		Body:    "\n",
	}
)

// StompFrame 表示一个 STOMP 帧
type StompFrame struct {
	Command string
	Headers map[string]string
	Body    string
}

func (f *StompFrame) Serialize() string {
	var builder strings.Builder

	// 写入命令行
	builder.WriteString(f.Command)
	builder.WriteString("\n")

	// 写入头部
	for key, value := range f.Headers {
		builder.WriteString(fmt.Sprintf("%s:%s\n", key, value))
	}

	// 空行分隔头部和正文
	builder.WriteString("\n")

	// 写入正文
	if f.Body != "" {
		builder.WriteString(f.Body)
	}

	// STOMP 帧以 null 字符结尾
	builder.WriteByte(0)

	return builder.String()
}

// parseStompFrame 解析 STOMP 帧
func parseStompFrame(data string) (*StompFrame, error) {
	frame := &StompFrame{
		Headers: make(map[string]string),
	}

	lines := splitLines(data)
	if len(lines) == 0 {
		return nil, ErrInvalidFrame
	}

	frame.Command = lines[0]

	i := 1
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		parts := strings.SplitN(lines[i], ":", 2)
		if len(parts) != 2 {
			return nil, ErrInvalidFrame
		}
		frame.Headers[parts[0]] = parts[1]
	}

	i++
	if i < len(lines) {
		frame.Body = strings.Join(lines[i:], "\n")
		frame.Body = strings.TrimSuffix(frame.Body, "\x00")
	}

	return frame, nil
}

func splitLines(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}
