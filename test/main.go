package main

import (
	"fmt"

	"heap"
)

func main() {
	input := []int{10, 0, -13, 99, 20, 15, 34, -29, 30, 40}
	var output []int
	h := heap.NewHeap()
	for _, v := range input {
		h.Add(v)
	}

	for {
		i, ok := h.Pop()
		fmt.Println(fmt.Sprintf("Pop %v (err: %v): [%s]", i, ok, h.String()))
		output = append(output, i)
		if ok != nil {
			break
		}
	}

	fmt.Println(output)

}
