package _go

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

//sync.WaitGroup timeout
func Test(t *testing.T) {
	wg := sync.WaitGroup{}
	c := 10
	wg.Add(c)
	go func() {
		for i := 0; i < c; i++ {
			time.Sleep(5 * time.Second)
			wg.Done()
		}
	}()

	since := time.Now()
	timer := time.NewTimer(2 * time.Second)
	endC := make(chan struct{})
	go func() {
		<-timer.C
		endC <- struct{}{}
	}()
	go func() {
		wg.Wait()
		endC <- struct{}{}

	}()

	<-endC
	fmt.Println(time.Now().Sub(since))
}

func TestContext(t *testing.T) {
	timeout, _ := context.WithTimeout(context.Background(), 2*time.Second)
	for {
		select {
		case <-timeout.Done():
			return
		default:
			time.Sleep(5 * time.Second)
		}
	}

	fmt.Println("end")
}

func TestVar(t *testing.T) {
	var a []int
	b := []int{}

	fmt.Println(a)
	fmt.Println(b)

	a = append(a, 1)
	b = append(b, 1)
	fmt.Println(a)
	fmt.Println(b)
}

func Test1(t *testing.T) {
	a := []int{1, 2, 3}
	a[0], a[1] = a[1], a[0]
	fmt.Println(a) //[2 1 3]

}
