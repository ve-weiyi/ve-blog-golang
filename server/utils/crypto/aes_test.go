package crypto

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
	"time"
)

// 对于使用 AES（Advanced Encryption Standard）算法的 CBC 模式，密钥的长度可以是 128 比特（16 字节）、192 比特（24 字节）或 256 比特（32 字节）。
func Test_AES(t *testing.T) {

	fmt.Println(time.Now().Unix())
	plaintext := []byte("791422171@qq.comasdsda")     // 待加密的数据
	key := []byte("1234567.1234567.1234567.1234567.") // 加密的密钥
	iv := []byte("1234567.1234567.")                  // 初始化向量
	aes := AesCFB

	log.Println("原文：", string(plaintext))
	log.Println("------------------ CBC模式 --------------------")
	encrypted := aes.AESEncrypt(plaintext, key, iv...)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := aes.AESDecrypt(encrypted, key, iv...)
	log.Println("解密结果：", string(decrypted))
}

func BenchmarkRSA(b *testing.B) {

	key, _ := rsa.GenerateKey(rand.Reader, RSA_KEY_SIZE)
	pub := &key.PublicKey
	priv := key

	//明文
	plaintext := []byte(Bit512)

	for i := 0; i < b.N; i++ {
		//key-1024bits data-512bits  time-222050 ns/op
		ciphertext, _ := rsa.EncryptOAEP(md5.New(), rand.Reader, pub, plaintext, nil)
		plaintext, _ = rsa.DecryptOAEP(md5.New(), rand.Reader, priv, ciphertext, nil)
	}
}

func TestRSA(t *testing.T) {

	key, _ := rsa.GenerateKey(rand.Reader, RSA_KEY_SIZE)
	pub := &key.PublicKey
	priv := key

	//明文
	plaintext := []byte(Bit512)

	//加密生成密文
	fmt.Printf("%q\n加密:\n", plaintext)
	ciphertext, e := rsa.EncryptOAEP(md5.New(), rand.Reader, pub, plaintext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%x\n", ciphertext)

	//解密得到明文
	fmt.Printf("解密:\n")
	plaintext, e = rsa.DecryptOAEP(md5.New(), rand.Reader, priv, ciphertext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%q\n", plaintext)
}
