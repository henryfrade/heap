package main

import (
	"fmt"
	"math"
	"testing"
)

func TestHeap_Add(t *testing.T) {
	t.Run("One element to empty heap", func(t *testing.T) {
		h := NewHeap()
		if len(h.array) != 0 {
			t.Fail()
		}
		h.Add(5)
		if len(h.array) != 1 || h.array[0] != 5 {
			t.Fail()
		}
	})
	t.Run("One element that is bigger than rest", func(t *testing.T) {
		h := NewHeap()
		h.Add(1)
		h.Add(-1)
		h.Add(5)
		if len(h.array) != 3 || h.array[0] != 5 {
			t.Fail()
		}
	})
	t.Run("One element that is smaller than the rest", func(t *testing.T) {
		h := NewHeap()
		h.Add(100)
		h.Add(200)
		h.Add(5)
		if len(h.array) != 3 || h.array[2] != 5 {
			t.Fail()
		}
	})
}

func TestHeap_Pop(t *testing.T) {
	t.Run("empty heap", func(t *testing.T) {
		h := NewHeap()
		_, ok := h.Pop()
		if ok == nil {
			t.Fail()
		}
	})
	t.Run("single element", func(t *testing.T) {
		h := NewHeap()
		h.Add(5)
		if len(h.array) != 1 {
			t.Fail()
		}
		v, ok := h.Pop()
		if v != 5 || ok != nil {
			t.Fail()
		}
	})
}

func TestHeap_String(t *testing.T) {
	t.Run("Empty Heap", func(t *testing.T) {
		h := NewHeap()
		if h.String() != "[]" {
			t.Fail()
		}
	})

	t.Run("One element", func(t *testing.T) {
		h := NewHeap()
		h.Add(5)
		if h.String() != "[5,]" {
			t.Fail()
		}
	})

	t.Run("A few elements", func(t *testing.T) {
		h := NewHeap()
		h.Add(5)
		h.Add(7)
		if h.String() != "[7,5,]" {
			fmt.Println(h.String())
			t.Fail()
		}
	})
}

func TestHeap(t *testing.T) {
	t.Run("sort array", func(t *testing.T) {
		input := []int{10, 0, -13, 99, 20, 15, 34, -29, 30, 40}
		h := NewHeap()
		for _, v := range input {
			h.Add(v)
		}
		prev := math.MinInt32
		var output []int
		for i := 0; i < len(input); i++ {
			v, ok := h.Pop()
			if v < prev || ok != nil {
				t.Fail()
			}
			output = append(output, v)
		}
		if len(h.array) > 0 {
			t.Fail()
		}
		if fmt.Sprintf("%v",output) != "[99 40 34 30 20 15 10 0 -13 -29]" {
			t.Fail()
		}
	})
}
