package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{4, 2, 8}
	sort.Ints(ints)
	fmt.Println(ints)

	fmt.Println("sorted:", sort.IntsAreSorted(ints))
}
