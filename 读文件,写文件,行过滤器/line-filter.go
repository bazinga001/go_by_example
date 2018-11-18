package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// NewScanner 直接读取一行
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fmt.Println(strings.ToUpper(sc.Text()))
	}

	// 检查Scan的错误, 文件结束标识符不会当做错误
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}
