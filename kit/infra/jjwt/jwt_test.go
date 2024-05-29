package jjwt

import (
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

func TestJwt(t *testing.T) {
	jj := NewJwtToken([]byte("test"))

	tk, err := jj.CreateToken(
		TokenExt{
			Uid:      0,
			Username: "121",
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "blog",
		})

	log.Println(tk, err)

	token, err := jj.ParserToken(tk)
	if err != nil {
		return
	}

	log.Println(jsonconv.ObjectToJsonIndent(token))
}
