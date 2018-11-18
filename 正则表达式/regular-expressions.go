package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 编译一个优化的正则表达式结构体
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringIndex("peach punch pinch", -1))

	// 限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 上面的MatchString也可以使用下面这种方式
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile 只返回一个参数
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper))

}
