package main

import (
	"fmt"
	"os"
)

func main() {
	// 当使用`os.Exit`时, defer 将不会执行
	defer fmt.Println("???")

	os.Exit(3)
}
