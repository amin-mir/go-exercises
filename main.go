package main

import (
	"fmt"

	"github.com/amin-mir/go-exercises/slice"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(slice.Reverse(a), a)

	slice.ReverseInPlace(a)
	fmt.Println(a)
}
