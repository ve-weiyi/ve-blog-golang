package svc

import (
	"log"
	"testing"
)

func Test_Captcha(t *testing.T) {
	sv := NewTestServiceContext()
	key := "test"
	captcha, err := sv.CaptchaHolder.GetCodeCaptcha(key)
	if err != nil {
		log.Fatal(err)
	}

	t.Log(captcha)
	if sv.CaptchaHolder.VerifyCaptcha(key, captcha) {
		t.Log("verify success")
	}

}
