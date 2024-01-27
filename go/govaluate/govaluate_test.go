package govaluate

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/Knetic/govaluate"
)

func Test(t *testing.T) {
	functions := map[string]govaluate.ExpressionFunction{
		"str2num": func(args ...interface{}) (interface{}, error) {
			argsStr, ok := args[0].(string)
			if !ok {
				return nil, errors.New("args[0] is not string")
			}
			argsNum, err := strconv.ParseFloat(argsStr, 10)
			if err != nil {
				return nil, err
			}
			return argsNum, nil
		},
		"nowUnixMilli": func(args ...interface{}) (interface{}, error) {
			return float64(time.Now().UnixMilli()), nil
		},
	}

	expr, err := govaluate.
		NewEvaluableExpressionWithFunctions("str2num() && str2num(total_count) > 5 && nowUnixMilli()-str2num(window_end)<5000000", functions)
	if err != nil {
		log.Fatal("syntax error:", err)
	}
	metaValue := `{
    "window_end":"1699326367534",
    "total_count":"6",
    "window_start":"1699326067534",
    "numerator_count":"0",
    "window_size_interval":"5m0s",
    "address":{
        "city":"6"
    }
}`
	metaMap := make(map[string]interface{})
	json.Unmarshal([]byte(metaValue), &metaMap)
	result, err := expr.Evaluate(metaMap)
	if err != nil {
		log.Fatal("evaluate error:", err)
	}
	fmt.Println(result)

}
