package mailtemplate

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	cases := []struct {
		name    string
		content string
		params  map[string]string
		want    string
		wantErr bool
	}{
		{
			name:    "replace all placeholders",
			content: "您的验证码是：{code}，{time}分钟内有效",
			params:  map[string]string{"code": "123456", "time": "15"},
			want:    "您的验证码是：123456，15分钟内有效",
		},
		{
			name:    "no placeholders",
			content: "系统通知",
			params:  nil,
			want:    "系统通知",
		},
		{
			name:    "missing param",
			content: "您的验证码是：{code}，{site_name}",
			params:  map[string]string{"code": "123456"},
			wantErr: true,
		},
		{
			name:    "repeated missing param reported once",
			content: "{link}{link}",
			params:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Render(c.content, c.params)
			if (err != nil) != c.wantErr {
				t.Fatalf("Render() err = %v, wantErr %v", err, c.wantErr)
			}
			if err != nil {
				if !strings.Contains(err.Error(), "missing template params") {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}
			if got != c.want {
				t.Errorf("Render() = %q, want %q", got, c.want)
			}
		})
	}
}
