package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 这个url示例包含:
	// 认证信息, 主机名, 端口号, 路径, 查询参数, 片段
	s := "postgres://user:pass@host.com:3456/path?key1=value1#fff"

	u, _ := url.Parse(s)

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	fmt.Println(u.Host)

	fmt.Println(u.Path)

	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
}
