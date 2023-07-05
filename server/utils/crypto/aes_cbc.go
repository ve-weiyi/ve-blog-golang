package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

/**
2.Cipher Block Chaining(CBC)  密码分组链接模式
明文被加密前要与前面的密文进行异或运算后再加密，因此只要选择不同的初始向量，相同的密文加密后会形成不同的密文，这是目前应用最广泛的模式。
这种模式是先将明文切分成若干小段，然后每一小段与初始块或者上一段的密文段进行异或运算后，再与密钥进行加密。
*/

func AESEncryptCBC(plaintext []byte, key []byte) (ciphertext []byte) {
	// AES加密算法的加密块必须是16字节(128bit)，所以不足部分需要填充，常用的填充算法是PKCS7。
	plaintext = pkcs7Padding(plaintext, aes.BlockSize)
	// 创建密文接收区
	ciphertext = make([]byte, len(plaintext))
	//加密向量,取密钥前16位
	iv := key[:aes.BlockSize]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//执行加密
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func AESDecryptCBC(ciphertext []byte, key []byte) (plaintext []byte) {
	// 创建明文接收区
	plaintext = make([]byte, len(ciphertext))
	//加密向量,取密钥前16位
	iv := key[:aes.BlockSize]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//执行解密
	blockMode.CryptBlocks(plaintext, ciphertext)
	//去填充
	plaintext = pkcs7UnPadding(plaintext)

	return plaintext
}
