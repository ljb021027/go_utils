package _map

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

type People struct {
	Name string
	Age  int
}

func Test(t *testing.T) {
	m := make(map[int]int)
	go func() {
		for i := 0; i < 100; i++ {
			i2 := m[i]
			fmt.Println(i2)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			m[i] = i
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestStructTO(t *testing.T) {
	people := People{
		Name: "ljb",
		Age:  18,
	}
	name := structs.Name(people)

	fmt.Println(name)
}

func TestMapStructure(t *testing.T) {
	people := People{
		Name: "ljb",
		Age:  18,
	}
	m := make(map[string]interface{}, 0)
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &m,
	})
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(people)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}

func TestJson(t *testing.T) {
	strChan := make(chan int, 10)
	tchan(strChan)
	fmt.Println(strChan)

	a, ok := <-strChan
	fmt.Println(a)
	fmt.Println(ok)

	b, ok := <-strChan
	fmt.Println(b)
	fmt.Println(ok)

	fmt.Println(strChan)
}

func tchan(strChan chan int) {
	defer close(strChan)
	for i := 10; i < 20; i++ {
		strChan <- i
	}
	a, ok := <-strChan
	fmt.Println(a)
	fmt.Println(ok)

}

func TestJson1(t *testing.T) {
	strjson := []byte(`
{
    "a":null
}`)
	m := make(map[string]interface{})
	err := json.Unmarshal(strjson, &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}

func TestTime(t *testing.T) {
	now := time.Now()
	format := now.Format("2006-01-02 15:04:05Z-0700")
	fmt.Println(format)
	split := strings.Split(format, "Z")
	fmt.Println(split[0])
	fmt.Println(split[1])

}

func TestDefer(t *testing.T) {

	for i := 0; i < 2; i++ {
		defer func() {
			fmt.Println(1)
		}()
	}
	fmt.Println(2)

}

func TestII(t *testing.T) {
	str := "123456789abcdefg"
	fmt.Println(str[15:16])
}

func TestFor(t *testing.T) {

	//m := map[int]int{
	//	1:1,
	//	2:2,
	//	3:3,
	//	4:4,
	//}

	//m:=map[int]Peo{
	//	1:{Name: "ljb1"},
	//	2:{Name: "ljb2"},
	//	3:{Name: "ljb3"},
	//
	//}
	//
	//for k,v := range m {
	//	fmt.Println(k)
	//	fmt.Println(v)
	//
	//}

	var m1 = []int{1, 2, 3}
	for i := range m1 {
		go func() {
			fmt.Print(i)
		}()
	}
	//block main 1ms to wait goroutine finished
	time.Sleep(10 * time.Millisecond)

}

type Peo struct {
	Name string
}

func Test11(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	//m:=100
	n := 10
	for i := 0; i < n; i++ {
		i2 := rand.Intn(100)
		fmt.Println(i2)
	}

}

func Test3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	m := float64(100)
	n := 10
	rest := m
	result := float64(0)
	for i := 0; i < n-1; i++ {
		min := rest / (float64(n-i-1) * 2)
		max := (rest * 3) / (float64(n-i-1) * 2)
		if min < 0.01 {
			min = 0.01
		}
		if max > 30 {
			max = 30
		}
		if (rest-max)/float64(n-i-1) < 0.01 {
			max = rest - float64(n-i-1)*0.01
		}
		count := min + rand.Float64()*(max-min)
		fmt.Println(count)
		result = result + count
		rest -= count
	}
	fmt.Println(rest)
	fmt.Println(result)
}
