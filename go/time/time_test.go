package time

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"

	parse, err := time.ParseInLocation(timeLayout, "2020-11-18 14:34:46.903", time.UTC)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parse)

	parse2, err := time.ParseInLocation(timeLayout, "2020-11-18 14:34:46.903", time.UTC)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parse2)
}
