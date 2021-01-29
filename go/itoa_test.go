package _go

import (
	"fmt"
	"testing"
)

const (
	a1 int = 1001 + iota // from 1001
	b1
	c1
	d1
)
const (
	a2 = 2001 + iota
	b2
	c2
)

func TestItoa(t *testing.T) {
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Println("==")
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)

}
