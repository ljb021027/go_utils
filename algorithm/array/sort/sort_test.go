package sort

import (
	"fmt"
	"testing"
)

func TestQsort(t *testing.T) {
	arr := []int{7, 3, 2, 5, 4}
	//qsort(arr, 0, len(arr)-1)

	quick1(arr)
	fmt.Println(arr)
}

func TestHeap(t *testing.T) {
	a := []int{5, 2, 10, 7, 3, 11, 99, 8}
	heapSort(a)
	fmt.Println(a)
	b := []int{1, 100}
	for i := 0; i < len(b); i++ {
	}
}

func TestFor(t *testing.T) {
	m := make(map[int]int)
	m[1] = 1
	fmt.Println(m)
}

func quick1(arr []int) {
	quick1_(arr, 0, len(arr)-1)
}

func quick1_(arr []int, s, e int) {
	if s >= e {
		return
	}
	i := par(arr, s, e)
	quick1_(arr, s, i-1)
	quick1_(arr, i+1, e)
}

func par(arr []int, s, e int) int {
	partition := arr[e]
	i := s
	j := s
	for ; i < e; i++ {
		if arr[i] > partition {
			//swap i,j
			arr[i], arr[j] = arr[j], arr[i]
			//temp := arr[j]
			//arr[j] = arr[i]
			//arr[i] = temp
			j++
		}
	}
	//swap e,j
	arr[e] = arr[j]
	arr[j] = partition
	return j
}
