package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func AESEncryptCFB(plaintext []byte, key []byte) (ciphertext []byte) {
	// AES加密算法的加密块必须是16字节(128bit)，所以不足部分需要填充，常用的填充算法是PKCS7。
	plaintext = pkcs7Padding(plaintext, aes.BlockSize)
	// 创建密文接收区
	ciphertext = make([]byte, len(plaintext))
	//加密向量,取密钥前16位
	iv := key[:aes.BlockSize]

	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCFBEncrypter(block, iv)
	blockMode.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}

func AESDecryptCFB(ciphertext []byte, key []byte) (plaintext []byte) {
	// 创建明文接收区
	plaintext = make([]byte, len(ciphertext))
	//加密向量,取密钥前16位
	iv := key[:aes.BlockSize]

	block, _ := aes.NewCipher([]byte(key))
	blockMode := cipher.NewCFBDecrypter(block, iv)
	blockMode.XORKeyStream(plaintext, ciphertext)

	plaintext = pkcs7UnPadding(plaintext)
	return plaintext
}
