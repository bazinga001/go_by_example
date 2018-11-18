package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("FOO", "123456")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("PATH:", os.Getenv("PATH"))

	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
