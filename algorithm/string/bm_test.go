package string

import (
	"testing"
)

func Test(t *testing.T) {
	var tts = []struct {
		a string
		b string
		i int
	}{
		{a: "abcde", b: "a", i: 0},
		{a: "abcde", b: "bcde", i: 1},
		{a: "abcde", b: "cf", i: -1},
		{a: "abcde", b: "abcde", i: 0},
		{a: "abcde", b: "abcdef", i: -1},
		{a: "abcde", b: "c", i: 2},
		{a: "abcde", b: "cd", i: 2},
	}

	for _, tt := range tts {
		if tt.i != str2(tt.a, tt.b) {
			t.Errorf("tt :%+v", tt)
		}

	}

}

func str1(a, b string) int {
	for i := 0; i < len(a); i++ {
		for j := 0; true; j++ {
			if j == len(b) {
				return i
			}
			if i+j == len(a) {
				return -1
			}
			if b[j] != a[i+j] {
				break
			}
		}
	}
	return -1
}

func str2(a, b string) int {
	for i := 0; i < len(a); i++ {
		m := 0
		for j := 0; j < len(b); j++ {
			if i+m > len(a)-1 {
				return -1
			}
			if a[i+m] == b[j] {
				m++
			} else {
				break
			}
		}
		if m == len(b) {
			return i
		}
	}
	return -1
}

//func (this BF) Index(s []byte) int {
//	p := len(this.temp)
//	q := len(s)
//	for i := 0; i+p < q; i++ {
//		if Compare(this.temp, s[i:]) {
//			return i
//		}
//	}
//	return -1
//}
