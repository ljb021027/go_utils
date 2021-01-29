package unknown

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreesum(t *testing.T) {
	a := []int{-1, 0, 1, 2, -1, -4}
	i := three(a, 0)
	fmt.Println(i)
}

func three(a []int, target int) [][]int {
	result := make([][]int, 0)
	if a == nil || len(a) < 3 {
		return result
	}
	//排序之后好去重，
	sort.Ints(a)

	//先固定i，然后双指针left right
	for i := 0; i < len(a); i++ {
		//
		if a[i] > target {
			break
		}
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		left := i + 1
		right := len(a) - 1
		for left < right {
			if a[i]+a[left]+a[right] == target {
				result = append(result, []int{a[i], a[left], a[right]})
				left++
				right--
				//去重
				for left < right && a[left-1] == a[left] {
					left++
				}
				//去重
				for left < right && a[right+1] == a[right] {
					right--
				}
			} else if a[i]+a[left]+a[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
