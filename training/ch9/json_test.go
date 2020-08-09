package ch9

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	source_data := make(map[string]interface{})
	source_data["name"] = "The Go programming Language"
	source_data["url"] = "https://Golang.org/"
	json_data, err := json.Marshal(source_data)
	if err != nil {
		panic(err)
	}
	fmt.Println(source_data)
	fmt.Println(string(json_data))

	decode_data := make(map[string]interface{})
	err = json.Unmarshal(json_data, &decode_data)
	if err != nil {
		panic(err)
	}
	fmt.Println(decode_data)
}
