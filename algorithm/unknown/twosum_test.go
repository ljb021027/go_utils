package unknown

import (
	"fmt"
	"testing"
)

func TestTwosum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	sum := twoSum(nums, 9)
	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		index1, ok := m[target-n]
		if ok {
			return []int{index1, i}
		}
		m[n] = i
	}
	return nil

}
