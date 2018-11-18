package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 最简单的方式
	data1 := []byte("hello\ngo\n")
	ioutil.WriteFile("/tmp/dat1", data1, 0644)

	// 更细粒度的写入操作
	f, _ := os.Create("/tmp/dat2")
	defer f.Close()

	n2, _ := f.Write([]byte("hello world"))
	fmt.Printf("wrote %d bytes\n", n2)

	n3, _ := f.WriteString("write string")
	fmt.Printf("wrote %d bytes\n", n3)

	// 将缓冲区的信息写入磁盘
	f.Sync()

	// bufio 提供带缓冲的写入器
	writer := bufio.NewWriter(f)
	n4, _ := writer.WriteString("bufio write String\n")
	fmt.Printf("wrote %d bytes\n", n4)

	writer.Flush()
}
