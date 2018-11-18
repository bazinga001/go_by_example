package main

import (
	"fmt"
	"time"
)

// 时间的格式化和解析

func main() {
	var p = fmt.Println

	now := time.Now()
	// 按照RFC3339进行格式化
	p(now.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	p(now.Format("Mon Jan 2 15:04:05 2006"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

}
