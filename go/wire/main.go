package wire

import (
	"fmt"
)

type IHello interface {
	Hello(name string)
}

type A struct {
}

func (*A) Hello(name string) {
	fmt.Println("hello " + name + ", i am a")
}

type D struct {
}

func (*D) Hello(name string) {
	fmt.Println("hello " + name + ", i am d")
}

type B struct {
	IHello
}

func (*B) Hello(name string) {
	fmt.Println("hello " + name + ", i am b")
}

type C struct {
	IHello
}

func main() {
	name := "Lee"
	a := A{}
	a.Hello(name) //hello Lee, i am a

	b := B{&A{}}
	b.Hello(name) //hello Lee, i am b

	b.IHello.Hello(name) //hello Lee, i am a

	c := C{&A{}}
	c.Hello(name) //hello Lee, i am a

	c.IHello = &D{}
	c.Hello(name) //hello Lee, i am d
}
