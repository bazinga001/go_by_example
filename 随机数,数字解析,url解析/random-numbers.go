package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// [0, n]
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	// [0.0, 1.0]
	fmt.Println(rand.Float64())

	// 要让伪随机数生成器有确定性, 可以给他一个明确的种子
	// 有确定性意味着, 如果用同样的种子, 生成的序列是相同的
	seed := rand.NewSource(42)
	r := rand.New(seed)
	fmt.Println(r.Intn(100))
	fmt.Println(r.Intn(100))
}
