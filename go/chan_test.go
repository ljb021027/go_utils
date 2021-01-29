package _go

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//chan 的 close
func TestChanClose(t *testing.T) {
	c := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			i, ok := <-c
			if !ok {
				return
			}
			time.Sleep(1 * time.Second)
			fmt.Println(i, ok)
		}
	}()

	fmt.Println("start push")
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	fmt.Println("close ")
	wg.Wait()
}

func TestBuffer(t *testing.T) {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(count)
}

type Record struct {
	Name string
}

func TestChan(t *testing.T) {
	recordsChan := make(chan int, 10)
	for i := 0; i < 5; i++ {
		recordsChan <- i
	}
	go func() {
		for i := 0; i < 10; i++ {
			i2 := <-recordsChan
			fmt.Println(i2)
			time.Sleep(1 * time.Second)
		}
	}()
	fmt.Println(fmt.Sprintf("1 %s", time.Now().String()))
	close(recordsChan)
	fmt.Println(fmt.Sprintf("2 %s", time.Now().String()))

	time.Sleep(10 * time.Second)
}
