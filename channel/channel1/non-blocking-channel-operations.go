package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <- messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no messagereveived")
	}	

	msg2 := "hi~"
	select {
	case messages <- msg2:
		fmt.Println("send message:", msg2)
	default:
		fmt.Println("no message sent")
	}

	// if add `default' condition, the struct became non-blocked
	select {
	case msg := <- messages:
		fmt.Println("reveived message", msg)
	case sig := <- signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
