package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// 执行一个简单的命令(没有参数或者和输入, 只有打印)
	cmd := exec.Command("date")
	out, _ := cmd.Output()
	fmt.Println("> date")
	fmt.Println(string(out))

	// 一个稍微复杂的例子
	cmd2 := exec.Command("grep", "hello")
	grepIn, _ := cmd2.StdinPipe()
	grepOut, _ := cmd2.StdoutPipe()
	defer grepOut.Close()

	cmd2.Start()
	grepIn.Write([]byte("hello world\ngoodbyte grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	cmd2.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 通过明确的参数和命令, 可以使用bash -c ...
	cmd3 := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, _ := cmd3.Output()
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
