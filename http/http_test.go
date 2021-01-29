package main

import (
	"fmt"
	"testing"
)

type List struct {
	val  int
	next *List
}

func Test(t *testing.T) {
	list := List{val: 1, next: &List{val: 2, next: &List{val: 3, next: &List{val: 4, next: nil}}}}
	result := reverse(list)
	fmt.Println(result)
}

func reverse(head List) *List {
	var result *List
	cur := &head
	for cur != nil {
		temp := cur.next
		cur.next = result
		result = cur
		cur = temp
	}
	return result
}
