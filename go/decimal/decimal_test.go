package decimal

import (
	"fmt"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	value := float64(9999.123456789123456789)
	//s := decimal.NewFromFloat(value).String()
	//fmt.Println(s)

	fmt.Println(strconv.FormatFloat(value, 'f', -1, 64))

}
