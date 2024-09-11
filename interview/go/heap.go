package main

import (
	"fmt"
	"sort"
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 和 Pop 使用指针接收器，因为它们会修改切片的长度。
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
func main() {
	h := &IntHeap{2, 2, 2, 2, 1, 3, 5}
	heap.Init(h)
	heap.Push(h, 0)
	fmt.Println("heap", h)

	fmt.Println("len", h.Len())
	sort.Ints(*h)
	fmt.Println("sort", h)

	// Search 返回满足条件的最小索引, 如果没有满足条件的元素，返回 len(*h)
	fmt.Println(sort.Search(len(*h), func(i int) bool {
		return (*h)[i] >= 2 
	}))
	fmt.Println(sort.Search(len(*h), func(i int) bool {
		return (*h)[i] > 2 
	}))
	// 从小到大排列时，条件设定为小于等于必定返回 0 或 len
	fmt.Println("<2\t", sort.Search(len(*h), func(i int) bool {
		return (*h)[i] < 2 
	}))
	fmt.Println("<=2\t", sort.Search(len(*h), func(i int) bool {
		return (*h)[i] <= 2 
	}))
	fmt.Println("<=100\t", sort.Search(len(*h), func(i int) bool {
		return (*h)[i] <= 100 
	}))
	// 找不到返回 len
	fmt.Println("<=5\t", sort.Search(len(*h), func(i int) bool {
		return (*h)[i] < -1
	}))
}