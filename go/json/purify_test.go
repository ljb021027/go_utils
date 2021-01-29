package json

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

var b = []byte(`{
		"string":"haha",
		"bool": true,
		"int": 123,
		"map":{
			"string": "gaga",
			"float": 0.99999999,
			"arr":[1,2]
		},
		"arr":[
			"a",
			"b"
		],
		"arr_map":[
			{
				"string":"haha",
				"int":123
			},
			{
				"string":"haha",
				"int":123
			}
		]
	}`)

func TestJsonPurify(t *testing.T) {

	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		t.Error(err.Error())
		return
	}
	for k, v := range m {
		v = purify(v)
		switch k {
		case "string":
			s := "haha"
			if v != s {
				t.Errorf("expect:%s actual:%s", s, v)
			}
		case "bool":
			s := "true"
			if v != s {
				t.Errorf("expect:%s actual:%s", s, v)
			}
		case "int":
			s := "123"
			if v != s {
				t.Errorf("expect:%s actual:%s", s, v)
			}
		case "map":
			//map遍历无序
			s := "string=gaga&float=0.99999999"
			if v != s {
				t.Errorf("expect:%s actual:%s", s, v)
			}
		case "arr":
			s := "[a,b]"
			if v != s {
				t.Errorf("expect:%s actual:%s", s, v)
			}
		case "arr_map":
			s := "[int=123&string=haha,int=123&string=haha]"
			if v != s {
				t.Errorf("expect:%s actual:%s", s, v)
			}
		}
		fmt.Println(fmt.Sprintf("%s:%s", k, v))
	}

}

func Test(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestField(t *testing.T) {
	people := People{
		Name:    "111",
		Age:     0,
		Address: "222",
	}

	marshal, _ := json.Marshal(people)
	fmt.Println(string(marshal))
}

func TestOrderJson(t *testing.T) {

	om := OrderedMap1{
		"age":  "20",
		"name": "John",
	}
	fmt.Println(om.ToJson("name", "age"))

}
func TestOrderJson2(t *testing.T) {
	dict := map[string]interface{}{
		"orderedMap": OrderedMap2{
			{"name", "John"},
			{"age", 20},
		},
	}
	dump, err := json.Marshal(dict)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", dump)

	marshal, _ := json.Marshal(dict)
	fmt.Println(string(marshal))

}

func TestOrderJson3(t *testing.T) {
	m := make(map[string]string, 0)
	bytes := []byte("{\n    \"app_id\":\"123123\",\n    \"sig\":\"nE0bXpOKm4BPqnjiV9zKT+G2Hlk=\",\n    \"source\":[\n        {\n            \"score_source\":\"task_product\",\n            \"source_quantity\":1,\n            \"source_value\":\"product_item_001\"\n        },\n        {\n            \"score_source\":\"amt\",\n            \"source_quantity\":10.1,\n            \"source_value\":\"CNY\"\n        }\n    ],\n    \"user_id\":\"077135621231379\",\n    \"user_token\":\"st_dummy\"\n}")
	err := json.Unmarshal(bytes, &m)
	assert.NoError(t, err)
}
