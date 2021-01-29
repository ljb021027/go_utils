package string

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/gorilla/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		src    string
		expect string
	}{
		{"helloworld", "helloworld"},
		{"hello_world", "hello_world"},
		{"HELLOWORLD", "h_e_l_l_o_w_o_r_l_d"},
		{"Helloworld", "helloworld"},
		{"HelloWorld", "hello_world"},
		{"HElloWorld", "hello_world"},
	}
	for _, test := range tests {
		actual := snakeString(test.src)
		if actual != test.expect {
			t.Errorf("expect:%s actual:%s", test.expect, actual)
		}
	}
}

var h = "HelloWorld"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		snakeString(h)
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toLowerSnakeCase(h)
	}
}

func TestQuery(t *testing.T) {
	d, err := decimal.NewFromString("5.4999")
	if err != nil {
		return
	}
	part := d.RoundBank(0).IntPart()
	fmt.Println(part)
}

func TestMap(t *testing.T) {
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)

}

func Test1(t *testing.T) {
	fmt.Println(60%30 == 0)

}

func Test2(t *testing.T) {
	encoder := schema.NewEncoder()
	peoples := []*Person{
		{
			Name:   "",
			Age:    0,
			Emails: nil,
			Extra:  nil,
		},
	}
	values := url.Values{}
	err := encoder.Encode(peoples, values)
	assert.NoError(t, err)

}
