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
)

func Test_B_1(t *testing.T) {
	plaintext := []byte("791422171@qq.co23422sadamaaa11") // 待加密的数据
	key := []byte("9876787656785679")                     // 加密的密钥
	log.Println("原文：", string(plaintext))

	log.Println("------------------ CBC模式 --------------------")
	encrypted := AESEncryptCBC(plaintext, key)
	//e08ace47cf7ddddfa3cc968874dfd5c81ffff653883522219dcdcdf44b3fe7d6
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	//4IrOR8993d+jzJaIdN/VyB//9lOINSIhnc3N9Es/59Y=
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AESDecryptCBC(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}

func Test_B_2(t *testing.T) {
	plaintext := []byte("791422171@qq.co23422sadamaaa1") // 待加密的数据
	key := []byte("1234567.1234567.1234567.")            // 加密的密钥
	log.Println("原文：", string(plaintext))
	//ypxEb/wBULTW+FkKkrhGBA==
	log.Println("------------------ ECB模式 --------------------")
	encrypted := AesEncryptECB(plaintext, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))

}

func Test_B_3(t *testing.T) {
	plaintext := []byte("791422171@qq.comasdsda") // 待加密的数据
	key := []byte("9876787656785679")             // 加密的密钥
	log.Println("原文：", string(plaintext))

	log.Println("------------------ CFB模式 --------------------")
	encrypted := AESEncryptCFB(plaintext, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AESDecryptCFB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}

func TestAESGCM(t *testing.T) {
	log.Println("原文：", string(plaintext))

	log.Println("------------------ CFB模式 --------------------")
	encrypted, mac := AESEncryptGCM(plaintext, string(key))
	log.Println("密文(hex)：", hex.EncodeToString([]byte(encrypted)))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString([]byte(encrypted)))
	decrypted := AESDecryptGCM(encrypted, string(key), mac)
	log.Println("解密结果：", decrypted)
}

var key = []byte("1234567.1234567.1234567.") // 加密的密钥

func BenchmarkAES(b *testing.B) {

	plaintext = BitsX10(BitsX10(BitsX10(BitsX10(Bit1024))))
	for i := 0; i < b.N; i++ {
		// 24bits		1024	  1024*10*10	1024*10*10*10*10
		// 	cbc		:808.4 ns/op   24871 ns/op   2354281 ns/op
		//  gcm		:1642 ns/op    40181 ns/op   4575092 ns/op
		//encrypted := AESEncryptCBC([]byte(plaintext), key)
		//AESDecryptCBC(encrypted, key)

		encrypted, mac := AESEncryptGCM(plaintext, string(key))
		AESDecryptGCM(encrypted, string(key), mac)
	}
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
