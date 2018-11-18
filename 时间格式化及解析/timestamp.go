package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unix := now.Unix()
	nano := now.UnixNano()
	fmt.Println("now:", now)
	fmt.Println("unix:", unix)
	fmt.Println("nano:", nano)

	// time.Unix(整数, 纳秒数)
	fmt.Println(time.Unix(unix, 0))
	fmt.Println(time.Unix(0, nano))
}
