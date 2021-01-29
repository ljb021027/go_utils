package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type People struct {
	Name    string
	Age     int `json:"age,omitempty"`
	Address string
}

// purify
func purify(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case float64:
		return fmt.Sprintf("%v", vv)
	case bool:
		return strconv.FormatBool(vv)
	case []interface{}:
		// [a,b,c]
		var sb strings.Builder
		sb.WriteString("[")
		for _, o := range vv {
			sb.WriteString(purify(o))
			sb.WriteString(",")
		}
		return sb.String()[:sb.Len()-1] + "]"
	case map[string]interface{}:
		// a=1&b=2
		var sb strings.Builder
		for k, v := range vv {
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(purify(v))
			sb.WriteString("&")
		}
		return sb.String()[:sb.Len()-1]
	default:
		return fmt.Sprintf("%v", vv)
	}
}

var JSONAPI = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()

// purify
func purify3(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(vv)
	case []interface{}, map[string]interface{}:
		marshal, err := json.Marshal(vv)
		if err != nil {
			return ""
		}
		return string(marshal)
	default:
		return fmt.Sprintf("%v", vv)
	}

}

func purify2(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(vv)
	case []interface{}:
		sb := strings.Builder{}
		sb.WriteString("[")
		for i := range vv {
			sb.WriteString("\"")
			sb.WriteString(purify2(vv[i]))
			sb.WriteString("\"")
			sb.WriteString(",")
		}
		return sb.String()[:sb.Len()-1] + "]"
	case map[string]interface{}:
		sb := strings.Builder{}
		sb.WriteString("{")
		for k, v := range vv {
			sb.WriteString("\"")
			sb.WriteString(k)
			sb.WriteString("\":")
			sb.WriteString(purify2(v))
			sb.WriteString("\"")
		}
		return sb.String()
	case OrderedMap:
		json, err := vv.MarshalJSON()
		if err != nil {
			return ""
		}
		return string(json)
	default:
		return fmt.Sprintf("%v", vv)
	}
}

type OrderedMap1 map[string]string

func (om OrderedMap1) ToJson(order ...string) string {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{', '\n'})
	l := len(order)
	for i, k := range order {
		fmt.Fprintf(buf, "\t\"%s\": \"%v\"", k, om[k])
		if i < l-1 {
			buf.WriteByte(',')
		}
		buf.WriteByte('\n')
	}
	buf.Write([]byte{'}', '\n'})
	return buf.String()
}

type KeyVal struct {
	Key string
	Val interface{}
}

// Define an ordered map
type OrderedMap2 []KeyVal

// Implement the json.Marshaler interface
func (omap OrderedMap2) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{")
	for i, kv := range omap {
		if i != 0 {
			buf.WriteString(",")
		}
		// marshal key
		key, err := json.Marshal(kv.Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteString(":")
		// marshal value
		val, err := json.Marshal(kv.Val)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}

	buf.WriteString("}")
	return buf.Bytes(), nil
}
