package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	p := Point{2, 3}
	// %v 打印结构体
	fmt.Printf("%v\n", p)

	// %+v 打印结构体包含结构体字段名
	fmt.Printf("%+v\n", p)

	// %T 打印值的类型
	fmt.Printf("%T\n", p)

	// %t bool值
	fmt.Printf("%t\n", true)

	// %b 二进制格式
	fmt.Printf("%b\n", 10)

	// %d %c %x %o %s %e %p %f

	// %q 带双引号输出
	fmt.Printf("%q\n", "\"string\"")

	// %x 使用base-16编码输出
	fmt.Printf("%x\n", "hex this")

	// Sprintf()返回一个字符串
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Fprintf()格式化并输出到 io.Writers (不是 os.Stdout)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
