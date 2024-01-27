package json

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMap_Sort(t *testing.T) {
	s := `
{
  "b": 2,
  "a": 1,
  "c": 3
}
`
	o := New()
	//json.Unmarshal([]byte(s), &o)
	err := o.UnmarshalJSON([]byte(s))
	assert.NoError(t, err)
	json, err := o.MarshalJSON()
	assert.NoError(t, err)
	fmt.Println(string(json))

}

func TestOrderMap2(t *testing.T) {
	s := `{
    "app_id":"123123",
    "sig":"nE0bXpOKm4BPqnjiV9zKT+G2Hlk=",
    "source":[
        {
		 "source_quantity":1,
            "score_source":"task_product",
           
            "source_value":"product_item_001"
        },
        {
            "score_source":"amt",
            "source_quantity":10.1,
            "source_value":"CNY"
        }
    ],
    "user_id":"077135621231379",
    "user_token":"st_dummy",
    "arr":[
        1,
        2
    ],
    "arrStr":[
        "1",
        "2"
    ]
}`

	o := New()
	o.SetEscapeHTML(false)

	err := o.UnmarshalJSON([]byte(s))
	assert.NoError(t, err)

	inputMap := make(url.Values)

	keys := o.Keys()
	for i := range keys {
		key := keys[i]
		get, b := o.Get(key)
		if !b {
			panic("!!")
		}
		inputMap.Set(key, purify3(get))
	}

	fmt.Println(inputMap)
}

func TestMap(t *testing.T) {
	s := `{"a":1,"b":2}`

	result := make(map[int]int)
	m := make(map[string]int)
	json.Unmarshal([]byte(s), &m)
	for i := 0; i < 100; i++ {
		for _, v := range m {
			result[v]++
			break
		}

	}

	fmt.Println(result)

}

func TestMap1(t *testing.T) {
	m := make(map[string]string, 0)
	s, ok := m["1"]
	fmt.Println(s)
	fmt.Println(ok)

	fmt.Println(m)
}
