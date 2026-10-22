package tokenx

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

// graceLabel 为从 refresh token 派生密钥提供域分隔。
const graceLabel = "tokenx-grace-v1"

// sealGrace 用旧 refresh token 派生的密钥加密轮换结果。
func sealGrace(refreshToken string, plaintext []byte) ([]byte, error) {
	aead, err := newGraceAEAD(refreshToken)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("tokenx: generate grace nonce: %w", err)
	}
	return aead.Seal(nonce, nonce, plaintext, nil), nil
}

// openGrace 用旧 refresh token 派生的密钥解密轮换结果。
func openGrace(refreshToken string, ciphertext []byte) ([]byte, error) {
	aead, err := newGraceAEAD(refreshToken)
	if err != nil {
		return nil, err
	}

	nonceSize := aead.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("tokenx: invalid grace payload")
	}
	nonce, payload := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aead.Open(nil, nonce, payload, nil)
	if err != nil {
		return nil, errors.New("tokenx: invalid grace payload")
	}
	return plaintext, nil
}

func newGraceAEAD(refreshToken string) (cipher.AEAD, error) {
	if refreshToken == "" {
		return nil, errors.New("tokenx: grace token is required")
	}

	mac := hmac.New(sha256.New, []byte(refreshToken))
	mac.Write([]byte(graceLabel))

	block, err := aes.NewCipher(mac.Sum(nil))
	if err != nil {
		return nil, fmt.Errorf("tokenx: create aes cipher: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("tokenx: create gcm: %w", err)
	}
	return aead, nil
}
