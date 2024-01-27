package _go

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	var a int
	a, b := 1, 2
	var c = addNum(a, b)
	if b != 0 {
		a := 32 //在此代码块中重声明，那表示在此代码块中，对a重新进行了声明。此会覆盖此代码块外的a=2的声明。并且此时a的作用域，也仅在此代码块中。
		fmt.Println(a)
	}
	fmt.Println(a) //此处的a，是外层的变量a。 a=1。
	fmt.Println(c)
}

func addNum(a, b int) int {
	return a + b
}
