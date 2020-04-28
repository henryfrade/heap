package main

import (
	"fmt"
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
