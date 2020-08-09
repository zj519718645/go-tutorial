package ch9

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestHex(t *testing.T) {
	str := []byte("Hello Go!")
	fmt.Println("source length ", len(str))
	dst := make([]byte, hex.EncodedLen(len(str)))
	fmt.Println("dst length ", len(dst))
	hex.Encode(dst, str)
	fmt.Println("encoded: ", string(dst))
	src2 := []byte(string(dst))
	dst2 := make([]byte, hex.DecodedLen(len(src2)))
	fmt.Println("dst2 length", len(dst2))
	_, err := hex.Decode(dst2, src2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("encoded: ", string(dst2))

}
