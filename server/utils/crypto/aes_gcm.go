package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

/*
AES  GCM 加密
key:加密key
plaintext：加密明文
ciphertext:解密返回字节字符串[ 整型以十六进制方式显示]
noncetext: 当前的mac
*/

func AESEncryptGCM(plaintext string, key string) (ciphertext, noncetext string) {
	plainbyte := []byte(plaintext)
	keybyte := []byte(key)

	block, _ := aes.NewCipher(keybyte)
	aesgcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, aesgcm.NonceSize())
	// 由于存在重复的风险，请勿使用给定密钥使用超过2^32个随机值。
	io.ReadFull(rand.Reader, nonce)

	//out := aesgcm.Seal(nonce, nonce, data, nil) nonce拼接在密文上
	cipherbyte := aesgcm.Seal(nil, nonce, plainbyte, nil)
	ciphertext = fmt.Sprintf("%x\n", cipherbyte)
	noncetext = fmt.Sprintf("%x\n", nonce)
	return
}

func AESDecryptGCM(ciphertext, key, noncetext string) (plaintext string) {
	cipherbyte, _ := hex.DecodeString(ciphertext)
	nonce, _ := hex.DecodeString(noncetext)
	keybyte := []byte(key)

	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainbyte, err := aesgcm.Open(nil, nonce, cipherbyte, nil)
	if err != nil {
		panic(err.Error())
	}

	plaintext = string(plainbyte[:])
	return plaintext
}
