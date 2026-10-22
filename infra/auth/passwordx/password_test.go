package passwordx

import "testing"

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{name: "正常密码", password: "123456"},
		{name: "空密码", password: "", wantErr: true},
		{name: "全空白", password: "      ", wantErr: true},
		{name: "过短", password: "12345", wantErr: true},
		{name: "过长", password: string(make([]byte, MaxPasswordLength+1)), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if tt.wantErr && err == nil {
				t.Fatal("want error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestValidatePasswordConfirm(t *testing.T) {
	if err := ValidatePasswordConfirm("123456", "123456"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := ValidatePasswordConfirm("123456", "654321"); err == nil {
		t.Fatal("want error, got nil")
	}
}
