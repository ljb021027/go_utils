package format

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/gorilla/schema"
	"github.com/stretchr/testify/assert"
)

type Abc struct {
	Name  string
	Names []string
}

func (Abc) Method1(a string, b int) {

}

func TestValue(t *testing.T) {
	abcs := []Abc{
		{
			Name: "ljb",
		},
	}
	values, err := query.Values(abcs)
	if err != nil {
		assert.NoError(t, err)
	}
	fmt.Println(values)

	abc := Abc{
		Names: []string{"ljb"},
	}
	values, err = query.Values(abc)
	if err != nil {
		assert.NoError(t, err)
	}
	fmt.Println(values)

	encoder := schema.NewEncoder()
	form := url.Values{}
	err = encoder.Encode(abcs, form)
	assert.NoError(t, err)
	fmt.Println(form)
}

var opt2 = Options{
	Query:   "2",
	ShowAll: false,
	Page:    222,
}

func Test(t *testing.T) {

	//opt := Options{
	//	Query:   "2",
	//	ShowAll: false,
	//	Page:    222,
	//	Array:   []string{"1","2"},
	//	ArrayA:  []A{
	//		{
	//			Name: "aaa",
	//			Age:  222,
	//		},
	//	},
	//}
	//v, _ := query.Values(opt)
	//fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
	abc := Abc{}
	of := reflect.TypeOf(abc)
	numMethod := of.NumMethod()
	fmt.Println(numMethod)
	method := of.Method(0)
	method.Func.Len()
}

func Test2(t *testing.T) {
	encoder := schema.NewEncoder()
	form := url.Values{}
	err := encoder.Encode(opt2, form)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(form.Encode())
}

func Benchmark_gorilla(b *testing.B) {
	encoder := schema.NewEncoder()
	for i := 0; i < b.N; i++ {
		form := url.Values{}
		err := encoder.Encode(opt2, form)
		if err != nil {
			return
		}

	}
}

func Benchmark_google(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := query.Values(opt2)
		if err != nil {
			return
		}
	}
}
