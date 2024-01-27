package point

import (
	"fmt"
	"testing"
)

var p1 People
var p2 *People

func Test(t *testing.T) {
	people := People{name: "ljb"}
	fmt.Printf("变量的地址: %x %p\n", &people, &people)
	people.SetName("ljb2")
	fmt.Println(people.name)
	Abc(people)

	var p3 People
	var p4 *People

	fmt.Printf("变量的地址: %x %p\n", &p1, &p1)
	fmt.Printf("变量的地址: %x %p\n", &p2, &p2)
	fmt.Printf("变量的地址: %x %p\n", &p3, &p3)
	fmt.Printf("变量的地址: %x %p\n", &p4, &p4)

	fmt.Println(64 << 1)
}

type People struct {
	name string
}

func (p *People) SetName(name string) {
	fmt.Printf("SetName 方法内变量的地址: %x %p\n", p, &p)

	p = &People{name: name}
}

func (p *People) SetName2(name string) {
	*p = People{name: name}
}

func Abc(p People) {
	fmt.Printf("Abc 方法内变量的地址: %x %p\n", &p, &p)
}
