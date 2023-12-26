package crypto

import (
	"fmt"
	"testing"
)

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
