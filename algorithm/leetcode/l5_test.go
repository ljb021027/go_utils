package leetcode

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	str := "bb"
	fmt.Println(str[0:1])
	fmt.Println(str[1:2])

}

func Test(t *testing.T) {
	fmt.Println(longestPalindrome("eeeeeeeee"))

	//fmt.Println(longestPalindrome("dddddddd"))
	//fmt.Println(longestPalindrome("tattarrattat"))
	//fmt.Println(longestPalindrome("ababababa"))
	//fmt.Println(longestPalindrome("cbbbbbd"))
	//
	//fmt.Println(longestPalindrome("aaaa"))
	//fmt.Println(longestPalindrome("bb"))
	//
	//fmt.Println(longestPalindrome("ab"))
	//
	//fmt.Println(longestPalindrome("cbbbbbd"))
	//
	//fmt.Println(longestPalindrome("abcdcba"))

}

func longestPalindrome(s string) string {
	srcLen := len(s)
	if srcLen <= 1 {
		return s
	}
	result := s[:1]
	for i := 0; i < srcLen; i++ {
		max := len(result)
		//2*(i+1) 为当前位置能得出的最长
		if max > 2*(i+1) {
			break
		}
		b1 := true
		b2 := true
		//从中心向两边，每次判断一个字符
		temp := ""
		for j := 1; j <= srcLen-i && (b1 || b2); j++ {
			if max < 2 && i+1 < srcLen && s[i:i+1] == s[i+1:i+2] {
				temp = s[i : i+2]
				break
			}
			if i-j < 0 {
				break
			}
			//以i为中心
			if b1 && i+j < srcLen && s[i-j:i-j+1] == s[i+j:i+j+1] {
				if 2*j+1 > max && 2*j+1 > len(temp) {
					temp = s[i-j : i+j+1]
				}
			} else {
				b1 = false
			}
			//以i-1和i的中间 为中心
			if b2 && s[i-j:i-j+1] == s[i+j-1:i+j] {
				if 2*j > max && 2*j > len(temp) {
					temp = s[i-j : i+j]
				}
			} else {
				b2 = false
			}
		}
		if len(temp) > max {
			result = temp
		}

	}

	return result

}
