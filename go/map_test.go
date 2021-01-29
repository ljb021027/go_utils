package _go

import (
	"fmt"
	"testing"
)

func TestRange1(t *testing.T) {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	for i, _ := range m {
		m[4] = 4
		fmt.Printf("%d%d ", i, m[i]) //11 22 33 44
	}
}
