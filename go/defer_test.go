package _go

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	fmt.Println(*a()) //2
	fmt.Println(b())  //2
	fmt.Println(c())  //0

}

func a() *int {
	var i int
	defer func(i *int) {
		*i++
		fmt.Println("a defer2:", *i) // 打印结果为 a defer2: 2
	}(&i)
	defer func(i *int) {
		*i++
		fmt.Println("a defer1:", *i) // 打印结果为 a defer1: 1
	}(&i)
	return &i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}

func c() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}
