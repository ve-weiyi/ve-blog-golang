package mailtemplate

import (
	"strings"
	"testing"

	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
)

func TestCodeEmail(t *testing.T) {
	html, err := CodeEmail(enums.CodeSceneRegister, "123456", "15")
	if err != nil {
		t.Fatalf("CodeEmail() error = %v", err)
	}

	for _, want := range []string{"注册验证码", "123456", "15 分钟", "<!DOCTYPE html>"} {
		if !strings.Contains(html, want) {
			t.Errorf("rendered email missing %q", want)
		}
	}
	if strings.Contains(html, "{") {
		t.Errorf("rendered email contains an unresolved placeholder: %s", html)
	}
}

func TestCodeSceneText(t *testing.T) {
	cases := []struct {
		scene string
		title string
	}{
		{enums.CodeSceneRegister, "注册验证码"},
		{enums.CodeSceneEmailLogin, "登录验证码"},
		{enums.CodeSceneResetPassword, "重置密码验证码"},
		{enums.CodeSceneBind, "绑定验证码"},
		{"unknown", "验证码"},
	}

	for _, c := range cases {
		title, intro := CodeSceneText(c.scene)
		if title != c.title {
			t.Errorf("CodeSceneText(%q) title = %q, want %q", c.scene, title, c.title)
		}
		if intro == "" {
			t.Errorf("CodeSceneText(%q) intro is empty", c.scene)
		}
	}
}
