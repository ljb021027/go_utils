package tool

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	aliResult, err := ReadCsv("/Users/liujiabei/Documents/xl/doc/sms/v3/07-29/events.csv.1", ",")
	if err != nil {
		fmt.Println(err)
		return
	}
	arMap := make(map[string][]string, 0)
	for _, ar := range aliResult {
		arMap[ar[0]] = ar
	}

	localResult, err := ReadCsv("/Users/liujiabei/Documents/xl/doc/sms/v3/07-29/ali.csv", ",")
	if err != nil {
		fmt.Println(err)
		return
	}
	succ := 0
	for _, lr := range localResult {
		if lr[8] != "发送成功" {
			continue
		}
		succ++
		_, ok := arMap[lr[0]]
		if !ok {
			fmt.Println(lr)
		}
	}
	fmt.Println(succ)
}

func Test2(t *testing.T) {
	localResult, err := ReadCsv("/Users/liujiabei/Documents/xl/doc/sms/v2/a.txt.2", "	")
	if err != nil {
		fmt.Println(err)
		return
	}
	localMap := make(map[string][]string, 0)
	for _, ar := range localResult {
		localMap[ar[0]] = ar
	}

	arsoftResult, err := ReadCsv("/Users/liujiabei/Documents/xl/doc/sms/v2/7月4日明细.csv", ",")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ar := range arsoftResult {
		_, ok := localMap[ar[4]]
		if !ok {
			fmt.Println(ar)
		}
	}
}
