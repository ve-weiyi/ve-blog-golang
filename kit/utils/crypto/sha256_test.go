package crypto

import (
	"fmt"
	"testing"
)

func TestWiki(t *testing.T) {
	//241f18af4356859f5f47b38d3f532b855b30884141c913b140a4f81ae5505dbc
	encrypted := Sha256v("学习区块链要记得记笔记", "")
	fmt.Println(encrypted)
}
