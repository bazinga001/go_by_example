package main

import (
	"fmt"
	"strings"
)

// 函数也可以作为变量, 可以用于简写某些常用方法
var p = fmt.Println

func main() {
	p("Contain:", strings.Contains("test", "es"))
	p("Count:", strings.Count("test", "t"))
	p("HasPreifx:", strings.HasPrefix("test", "te"))
	p("HasSuffix:", strings.HasSuffix("test", "st"))
	p("Index:", strings.Index("test", "e"))
	p("Join", strings.Join([]string{"aaa", "bbb"}, "--"))
	p("Repeat:", strings.Repeat("a", 5))
	p("Replace:", strings.Replace("foo", "o", "x", -1))
	p("Split:", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:", strings.ToLower("TEST"))
	p("ToUpper:", strings.ToUpper("test"))
	p()

	p("Len:", len("hello"))
	p("Char:", "hello"[3])
}
