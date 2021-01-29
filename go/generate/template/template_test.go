package template

import (
	"fmt"
	"testing"
	"text/template"
)

func Benchmark(b *testing.B) {
	tmpl, err := template.New("test").Parse("The name for student {{.ID}} is {{.Name}}")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		Parse(tmpl)
	}
}

func Test1(t *testing.T) {
	tmpl, err := template.New("test1").Parse("The name for student {{.ID}} is {{.Name}}")
	if err != nil {
		panic(err)
	}
	parse := Parse(tmpl)
	fmt.Println(parse)
}

func Test(t *testing.T) {
	abc()
}
