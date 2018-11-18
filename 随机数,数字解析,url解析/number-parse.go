package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 0 表示自动推断字符串所表示的数字的进制
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Atoi 基础的10进制整数转换函数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
}
