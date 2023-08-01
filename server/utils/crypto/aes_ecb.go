package crypto

import (
	"crypto/aes"
	"errors"
)

//golang crypt包的AES加密函数的使用 https://www.jianshu.com/p/47e8c137ecd4

var ErrorAESEncrypt = errors.New("加密字符串错误！")

/**
1、Electronic Code Book(ECB) 电子密码本模式
ECB模式是最早采用和最简单的模式，相同的明文将永远加密成相同的密文。无初始向量，容易受到密码本重放攻击，一般情况下很少用.
它将加密的数据分成若干组，每组的大小跟加密密钥长度相同，然后每组都用相同的密钥进行加密。
*/

func AesEncryptECB(plaintext []byte, key []byte) (ciphertext []byte) {
	// AES加密算法的加密块必须是16字节(128bit)，所以不足部分需要填充，常用的填充算法是PKCS7。
	plaintext = pkcs7Padding(plaintext, aes.BlockSize)
	// 创建密文接收区
	ciphertext = make([]byte, len(plaintext))

	block, _ := aes.NewCipher(key)
	// 分组分块加密
	for bs, be := 0, block.BlockSize(); bs < len(plaintext); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(ciphertext[bs:be], plaintext[bs:be])
	}

	return ciphertext
}

func AesDecryptECB(ciphertext []byte, key []byte) (plaintext []byte) {
	// 创建明文接收区
	plaintext = make([]byte, len(ciphertext))

	block, _ := aes.NewCipher(key)
	// 分组分块解密
	for bs, be := 0, block.BlockSize(); bs < len(ciphertext); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(plaintext[bs:be], ciphertext[bs:be])
	}

	plaintext = pkcs7UnPadding(plaintext)
	return plaintext
}
