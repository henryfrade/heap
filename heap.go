package heap

import (
	"errors"
	"fmt"
	"strings"
)

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

func (h *heap) Pop() (int, error) {
	if len(h.array) == 0 {
		return 0, errors.New("heap is empty")
	}

	top := h.array[0]
	newLength := len(h.array) - 1
	h.array[0] = h.array[newLength]
	h.array = h.array[:newLength]
	if len(h.array) == 0 {
		return top, nil
	}

	for cI := 0; cI < len(h.array); {
		curr := h.array[cI]
		maxI, ok := h.maxChildIndex(cI)
		if ok != nil {
			break
		}
		maxChild := h.array[maxI]
		if maxChild <= curr {
			break
		}
		h.array[cI] = maxChild
		h.array[maxI] = curr
		cI = maxI
	}
	return top, nil
}

func (h *heap) maxChildIndex(parentIndex int) (maxI int, err error) {
	li := parentIndex*2 + 1
	ri := li + 1
	if li >= len(h.array) {
		return 0, errors.New("no children")
	}
	maxI = li
	if ri < len(h.array) {
		right := h.array[ri]
		if right > h.array[maxI] {
			maxI = ri
		}
	}
	return
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
