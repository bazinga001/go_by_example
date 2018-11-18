package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 有时候, 我们希望go能智能地处理Unix信号
// 比如当服务器收到`SIGTERM`时, 可以自动关机
// 或者命令行工具在收到`SIGINT`时, 可以停止处理输入信息

// go 通过通过来处理信号

func main() {
	// go 通过向通道发送`os.Signal`来进行信号通知
	// 我们还需要创建一个通道来接收这些通知
	sigs := make(chan os.Signal, 1)
	// 同时还需要创建一个用于在程序可以结束时进行通知的通道
	done := make(chan bool, 1)

	// `singal.Notify'注册这个给定的通道用于接收特定信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting singal")
	<-done
	fmt.Println("exiting")
}
