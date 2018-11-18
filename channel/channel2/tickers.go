package main

import (
	"fmt"
	"time"
)

func main() {
	// in chinese: da dian qi
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at:", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
