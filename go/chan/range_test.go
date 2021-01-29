package main

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(3*time.Second)
		c1<-1
		c2<-1
	}()


	for {
		select {
		case <-c1:
			fmt.Println("c1")
		case <-c2:
			fmt.Println("c2")
		}
		fmt.Println("ee")
	}


}
