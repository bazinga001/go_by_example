package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	// 通常下面的内容更有用
	fmt.Println(args[1:])
}
