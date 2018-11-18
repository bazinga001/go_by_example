package main

import (
	"fmt"
	"time"
)

func main() {
	var p = fmt.Println
	now := time.Now()
	p(now)

	// 注意, 时间总是和位置相关的(比如时区)
	then := time.Date(2019, 1, 30, 4, 35, 11, 11111111, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
