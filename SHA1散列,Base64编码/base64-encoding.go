package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "this is test string"

	encode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encode)

	decode, _ := base64.StdEncoding.DecodeString(encode)
	fmt.Println(string(decode))

	// 使用url兼容的base64格式进行编码解码
	uencode := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uencode)
	udecode, _ := base64.URLEncoding.DecodeString(uencode)
	fmt.Println(string(udecode))
}
