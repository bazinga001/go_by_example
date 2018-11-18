package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// Marshal 编码 (转化为json格式)
	// 基本类型
	fmt.Println(Must(json.Marshal(true)))
	fmt.Println(Must(json.Marshal(1)))
	fmt.Println(Must(json.Marshal(2.34)))
	fmt.Println(Must(json.Marshal("gopher")))
	fmt.Println(Must(json.Marshal([]string{"apple", "banaba", "pear"})))
	fmt.Println(Must(json.Marshal(map[string]int{"apple": 5, "lettuce": 7})))

	// 自定义类型
	resp1 := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	fmt.Println(Must(json.Marshal(resp1)))

	resp2 := &Response2{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear"},
	}
	fmt.Println(Must(json.Marshal(resp2)))

	// Unmarshal 解码 (json字符串转化为基本类型或者自定义类型数据)
	// 需要提供一个变量来接收
	data := []byte(`{"num":3, "strs":["aaa", "bbb"]}`)
	var value map[string]interface{}
	json.Unmarshal(data, &value)
	fmt.Println(value)

	fmt.Println(value["num"].(float64))

	// 访问这个字符比较复杂
	fmt.Println(value["strs"].([]interface{})[0].(string))

	// NewEncoder 直接将json编码输出到io.Writer流中
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(map[string]int{"apple": 5, "lettcue": 7})
}

func Must(b []byte, err error) string {
	if err != nil {
		panic(err)
	}
	return string(b)
}
