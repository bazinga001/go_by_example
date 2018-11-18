package main

import (
	"flag"
	"fmt"
)

var p = fmt.Println

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("number", 42, "an int")

	boolPtr := flag.Bool("fork", false, "a bool")

	var value string
	flag.StringVar(&value, "svar", "bar", "a string val")

	// 在所有的标志声明完成后, 使用Parse()来执行
	flag.Parse()

	p("word:", *wordPtr)
	p("number:", *numPtr)
	p("fork:", *boolPtr)
	p("svar:", value)
	p("tail:", flag.Args())
}

// 使用 go run xx.go -h 查看帮助
