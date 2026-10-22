package mq

import (
	"testing"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/config"
)

func TestBuildAMQPURL(t *testing.T) {
	tests := []struct {
		name string
		conf config.RabbitMQConf
		want string
	}{
		{
			name: "普通口令",
			conf: config.RabbitMQConf{Host: "mq", Port: "5672", Username: "user", Password: "pass"},
			want: "amqp://user:pass@mq:5672/",
		},
		{
			name: "口令含 @ : /",
			conf: config.RabbitMQConf{Host: "mq", Port: "5672", Username: "user", Password: "p@ss:word/1"},
			want: "amqp://user:p%40ss%3Aword%2F1@mq:5672/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildAMQPURL(tt.conf); got != tt.want {
				t.Errorf("buildAMQPURL() = %q, want %q", got, tt.want)
			}
		})
	}
}
