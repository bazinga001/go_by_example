package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 获取命令的路径
	path, _ := exec.LookPath("ls") // "bin/ls"

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	// 实际上执行的是 `/bin/ls -a -l -h`, 需要传递环境参数
	syscall.Exec(path, args, env)
}

// go不提供经典的fork()
// 运行go协程, 生成进程, 执行进程覆盖了fork的大多数使用场景
