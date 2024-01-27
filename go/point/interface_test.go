package point

import (
	"fmt"
	"testing"
)

type Table interface {
	GetDbInfo() Table
}

type OrderInfo struct {
	Name string
}

func (o OrderInfo) GetDbInfo() Table {
	return o
}

func Print(t Table) {
	info := t.GetDbInfo()
	orderInfo := info.(OrderInfo)
	fmt.Println(orderInfo.Name)
}

func TestInterface(t *testing.T) {
	info := OrderInfo{
		Name: "ljb",
	}

	Print(info)
}
