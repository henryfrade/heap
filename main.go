package main

import (
	"errors"
	"fmt"
)
func main() {
	input := []int{10, 0, -13, 99, 20, 15, 34, -29, 30, 40}
	h := NewHeap()
	for _, v := range input {
		h.Add(v)
	}
}

type heap struct {
	array []int
}

func NewHeap() *heap {
	return &heap{
		array: make([]int, 0, 100),
	}
}

func (h *heap) Add(i int) {
	h.array = append(h.array, i)
	//fmt.Println(fmt.Sprintf("Add %v: [%s]", i, h.String()))
	curr := len(h.array) - 1
	for curr > 0 {
		papa := (curr - 1) / 2
		if h.array[papa] < h.array[curr] {
			t := h.array[papa]
			h.array[papa] = h.array[curr]
			h.array[curr] = t
		}
		curr = papa
	}
}

func (h *heap) String() string {
	var buffer strings.Builder
	buffer.Grow(len(h.array) * 2 + 2)
	buffer.WriteString("[")
	for _, v := range h.array {
		buffer.WriteString(fmt.Sprintf("%v,", v))
	}
	buffer.WriteString("]")
	return buffer.String()
}
