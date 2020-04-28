package main

import (
	"errors"
	"fmt"
)
func main() {
	h := NewHeap()
}

type heap struct {
	array []int
}

func NewHeap() *heap {
	return &heap{
		array: make([]int, 0, 100),
	}
}
