package leetcode

import (
	"fmt"
	"testing"
)

func TestL4(t *testing.T) {
	//nums1 := []int{1, 3, 4, 9}
	//nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//nums1 := []int{1, 2}
	//nums2 := []int{3, 4}
	nums1 := []int{1, 3}
	nums2 := []int{2}
	n := findMedianSortedArrays(nums1, nums2)
	fmt.Println(n)
}

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 0 {
		//偶数
		return (f1(nums1, nums2, totalLen/2) + f1(nums1, nums2, totalLen/2+1)) / 2
	} else {
		//奇数
		return f1(nums1, nums2, totalLen/2+1)
	}
}

//nums1 和 nums2 中第k小的
func f1(nums1 []int, nums2 []int, k int) float64 {
	i := 0
	j := 0

	for k > 0 {
		//nums1 没有了
		if i == len(nums1) {
			return float64(nums2[j+k-1])
		}
		//nums2 没有了
		if j == len(nums2) {
			return float64(nums1[i+k-1])
		}
		//退出
		if k == 1 {
			if nums1[i] <= nums2[j] {
				return float64(nums1[i])
			} else {
				return float64(nums2[j])
			}
		}
		s := k / 2
		//超过长度，取最后一个
		index1 := i + s - 1
		if i+s > len(nums1) {
			index1 = len(nums1) - 1
		}
		index2 := j + s - 1
		if j+s > len(nums2) {
			index2 = len(nums2) - 1
		}

		if nums1[index1] <= nums2[index2] {
			s = index1 - i + 1
			i = index1 + 1
		} else {
			s = index2 - j + 1
			j = index2 + 1
		}
		//剩余k最小
		k = k - s
	}
	return 0

}
