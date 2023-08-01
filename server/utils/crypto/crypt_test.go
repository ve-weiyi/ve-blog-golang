package crypto

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// 一个字节Byte由8位bit。8*64=512
var Bit256 = "012345678912345.012345678912345."
var Bit512 = "012345678912345.012345678912345." + "012345678912345.012345678912345."
var Bit1024 = Bit512 + Bit512
var plaintext = Bit512 + Bit512

// 组成32位16进制 128bit x 4
var ciphertext = ""

func TestBcrypt(t *testing.T) {

	//$2a$10$D3LbggNCcxz95XOr5CkLdeDadDc22xoSISMSADvM3p2BCmO49x1Yu
	//$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
	//\__/\/ \____________________/\_____________________________/
	//Alg Cost      Salt 128bits               Hash 192bits
	ciphertext = BcryptHash(plaintext)
	log.Println(BcryptHash(Bit512))
	log.Println(BcryptCheck(plaintext, ciphertext))
	log.Println(plaintext)
	log.Println(ciphertext)

	//6.60s
	//for i := 0; i < 10; i++ {
	//	bcrypt.GenerateFromPassword([]byte(Bit256), bcrypt.DefaultCost)
	//}
}

func TestMd5(t *testing.T) {
	//1bdf247646854ad6d841ba6b0cd376fe
	log.Println(Md5v(plaintext, "123"))
}

func TestSha256(t *testing.T) {
	log.Println(Sha256v(plaintext, "123"))

}

func BenchmarkMd5v(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 494.4 ns/op
		Md5v(plaintext, "123")
	}
}

func BenchmarkSha256v(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 249.4 ns/op
		Sha256v(plaintext, "123")
	}
}

func BenchmarkBcrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// cost=4:1075624 ns/op    cost=10:65876130 ns/op
		bcrypt.GenerateFromPassword([]byte(Bit256), bcrypt.DefaultCost)
	}
}

func BitsX10(src string) string {
	return src + src + src + src + src + src + src + src + src + src
}

func TestBase64(t *testing.T) {
	s := "{  }"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	//base64.StdEncoding.EncodeToString([]byte(s))
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	fmt.Println("base64.NewEncoding(encodeStd).EncodeToString")
	fmt.Println(s64)

	s64_std, _ := base64.StdEncoding.DecodeString(s64)
	fmt.Println("base64.StdEncoding.EncodeToString")
	fmt.Println(string(s64_std))
}

func BenchmarkECC(b *testing.B) {
	prvKey, err := genPrivateKey()
	if err != nil {
		fmt.Println(err)
	}
	pubKey := prvKey.PublicKey
	plain := "我们没什么不同"
	for i := 0; i < b.N; i++ {
		//     cost=102245 ns/op --0.0001s
		cipher, _ := ECCEncrypt(plain, &pubKey)
		plain, err = ECCDecrypt(cipher, prvKey)
	}
}

func TestECC(t *testing.T) {
	prvKey, err := genPrivateKey()
	if err != nil {
		fmt.Println(err)
	}
	pubKey := prvKey.PublicKey
	plain := Bit1024
	for i := 0; i < 10000; i++ {
		//     cost=102245 ns/op --0.0001s
		cipher, _ := ECCEncrypt(plain, &pubKey)
		plain, err = ECCDecrypt(cipher, prvKey)
	}
	//prvKey, err := genPrivateKey()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//pubKey := prvKey.PublicKey
	//plain := "我们没什么不同"
	//fmt.Printf("明文：%s\n", plain)
	//cipher, err := ECCEncrypt(plain, &pubKey)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("密文：%v\n", cipher)
	//plain, err = ECCDecrypt(cipher, prvKey)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("明文：%s\n", plain)
}
