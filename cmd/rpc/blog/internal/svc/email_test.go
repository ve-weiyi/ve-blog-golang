package svc

import (
	"testing"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/config"
)

func TestToEmailConfigMapsSSL(t *testing.T) {
	in := config.EmailConf{
		Host:     "smtp.example.com",
		Port:     465,
		Username: "user",
		Password: "pass",
		Nickname: "blog",
		BCC:      []string{"bcc@example.com"},
		IsSsl:    true,
	}

	got := toEmailConfig(in)

	if !got.SSL {
		t.Error("SSL = false, want true")
	}
	if got.Host != in.Host || got.Port != in.Port || got.Username != in.Username ||
		got.Password != in.Password || got.Nickname != in.Nickname {
		t.Errorf("config fields mismatch: %+v", got)
	}
	if len(got.BCC) != 1 || got.BCC[0] != "bcc@example.com" {
		t.Errorf("BCC = %v, want [bcc@example.com]", got.BCC)
	}
}

func TestNewEmailDeliverRejectsImplicitTLSMismatch(t *testing.T) {
	if _, err := NewEmailDeliver(config.EmailConf{Port: 465, IsSsl: false}); err == nil {
		t.Fatal("port 465 with is_ssl=false should be rejected")
	}
}
