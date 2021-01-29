package _go

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"unicode/utf8"
)

//range之前会拷贝，slice map 本身就是引用类型。 arr是值类型，
func TestArray(t *testing.T) {
	//v := make([]int, 4)
	//v = append(v, 1)
	//v = append(v, 2)
	//v = append(v, 3)

	v := []int{1, 2, 3} //这里是应该slice数组内部的原因

	for i, _ := range v {
		v = append(v, 4)
		fmt.Printf("%d%d ", i, v[i])
	}
}

func TestBigArr(t *testing.T) {
	//假设值都为1，这里只赋值3个
	var arr = [102400]int{1, 1, 1}
	for i, n := range &arr {
		//just ignore i and n for simplify the example
		_ = i
		_ = n
	}
}

//slice的坑 Cap够的时候 不需要扩容，append都是在原slice里，扩容后，就不影响原来的
func TestSlice(t *testing.T) {
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:2]
	fmt.Println(reflect.TypeOf(arr1).Kind())
	fmt.Println(reflect.TypeOf(slice1).Kind())

	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("slice1:", slice1)
	fmt.Println("arr1:", arr1)
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:3]
	slice2 = append(slice2, 6, 7, 8)
	fmt.Println("slice2:", slice2)
	fmt.Println("arr2:", arr2)
}

func TestForRangeRoutine(t *testing.T) {
	array := []int{1, 2, 3}
	for value := range array {
		go func() {
			fmt.Println(value)
		}()
	}
}

func TestStringLen(t *testing.T) {
	s := "世界"
	fmt.Println(len(s)) // 6

	fmt.Println(utf8.RuneCountInString(s)) // 2
}

func TestSliceAdd(t *testing.T) {
	a := []string{"1"}

	for i := 0; i < 10; i++ {
		b := a
		b = append(b, strconv.Itoa(i))
		fmt.Println(b)
	}

}

func TestAdd(t *testing.T) {
	a := make([]string, 1, 10)
	fmt.Println(fmt.Sprintf("1 %p", &a))
	add(a)
	fmt.Println(fmt.Sprintf("4 %p", &a))
	fmt.Println(a)

}

func add(a []string) {
	fmt.Println(fmt.Sprintf("2 %p", &a))
	a = append(a, "2")
	fmt.Println(a)
	fmt.Println(fmt.Sprintf("3 %p", &a))
}

func TestAdds(t *testing.T) {

	adds("1")
	arr := []string{"1", "2", "3", "4", "5"}
	strings := arr[:len(arr)-1]
	fmt.Println(strings)
}

func adds(keys ...string) {
	for i := 0; i < 10; i++ {
		tempKeys := keys
		tempKeys = append(tempKeys, "2")
		keys = append(keys, "3")
		fmt.Println(tempKeys)
	}

}

func TestSortTimeStr(t *testing.T) {
	strings := []string{
		"2020-11-24 20:43:51.395",
		"2020-11-23 20:43:51.295",
		"2020-11-24 20:43:51.295",
		"2020-11-23 01:43:51.295",
	}

	sort.Strings(strings)

	fmt.Println(strings)
}
