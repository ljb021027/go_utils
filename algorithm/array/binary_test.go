package array

import (
	"fmt"
	"math"
	"testing"
)

var arr = []int{1, 2, 5, 9, 19, 33, 44}

func TestBinaryFind(t *testing.T) {
	index := binaryLoop(arr, 1)
	fmt.Println(index)
}

func binaryLoop(a []int, b int) int {
	left := 0
	right := len(a) - 1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] > b {
			right = mid - 1
		} else if a[mid] < b {
			left = mid + 1
		} else {
			return mid
		}
	}
	if arr[right] == b {
		return right
	} else {
		absLeft := math.Abs(float64(arr[left] - b))
		absRight := math.Abs(float64(arr[right] - b))
		if absLeft >= absRight {
			return right
		} else {
			return left
		}
		return -1
	}
}
