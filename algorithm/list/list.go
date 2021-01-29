package main

import (
	"fmt"
	"strings"
)

type List struct {
	val  string
	next *List
}

func (l *List) toString() string {
	if l == nil {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(l.val)
	next := l.next
	for next != nil {
		sb.WriteString(">")
		sb.WriteString(next.val)
		next = next.next
	}
	return sb.String()
}

func main() {
	list := &List{
		val: "0",
		next: &List{
			val: "1",
			next: &List{
				val: "2",
				next: &List{
					val:  "3",
					next: nil,
				},
			},
		},
	}

	fmt.Println(list.toString())
	l := reverse4(list)
	fmt.Println(l.toString())
}
func reverse2(list *List) *List {
	var result *List
	cur := list
	for cur != nil {
		temp := cur.next
		cur.next = result
		result = cur
		cur = temp
	}
	return result
}

func reverse3(list *List) *List {
	var result *List
	for list != nil {
		temp := list.next
		list.next = result
		result = list
		list = temp
	}
	return result
}

func reverse4(list *List) *List {
	if list.next == nil {
		return list
	}
	l := reverse4(list.next)
	list.next.next = list
	list.next = nil
	return l
}
