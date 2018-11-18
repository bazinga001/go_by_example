package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 将文件读取到内存
	// 这里有个小细节, 不能用"~"表示家目录
	data, err := ioutil.ReadFile("/home/chen/.bashrc")
	fmt.Println(err)
	fmt.Println(string(data))

	// 对文件怎么读, 读到哪里进行更多的控制
	f, _ := os.Open("/home/chen/ss.sh")
	defer f.Close()
	// 最多只读取16字节
	buffer := make([]byte, 16)
	n, _ := f.Read(buffer)
	fmt.Printf("%d bytes: %s\n", n, string(buffer))

	// Seek()的两个参数, 第一个表示offset, 第二个参数表示起始位置
	buffer2 := make([]byte, 8)
	ret, _ := f.Seek(6, os.SEEK_SET)
	n2, _ := f.Read(buffer2)
	fmt.Printf("%d bytes @ %d: %s\n", n2, ret, string(buffer2))

	// 将文件指针放到文件的开头
	f.Seek(0, 0)

	// bufio 实现带缓冲的读取, 性能更好, 并且提供了很多附加读取函数
	reader := bufio.NewReader(f)
	buf3, _ := reader.Peek(5)
	fmt.Printf("5 bytes: %s\n", string(buf3))

}
