package ch9

import (
	"encoding/base64"
	"errors"
	"fmt"
	"testing"
)

func mustDecode(enc *base64.Encoding, str string) string {
	data, err := enc.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func mustEncode(enc *base64.Encoding, str string) {
	encStr := enc.EncodeToString([]byte(str))
	fmt.Println(encStr)
	decStr := mustDecode(enc, encStr)
	if decStr != str {
		panic(errors.New("unequals!"))
	}
}

func TestBase64(t *testing.T) {
	const testStr = "GO语言编程"
	mustEncode(base64.StdEncoding, testStr)
	mustEncode(base64.URLEncoding, testStr)
	mustEncode(base64.RawStdEncoding, testStr)
	mustEncode(base64.RawURLEncoding, testStr)
}
