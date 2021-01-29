package template

import (
	"bytes"
	"fmt"
	"text/template"
)

type Student struct {
	ID   uint
	Name string
}

func Parse(tmpl *template.Template) string {
	stu := Student{0, "jason"}

	var bf bytes.Buffer
	err := tmpl.Execute(&bf, stu)
	if err != nil {
		panic(err)
	}
	return bf.String()
}

func abc() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

}
