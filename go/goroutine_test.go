package _go

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 2; i++ {
		go func(index int) {
			for {
				if index == 1 {
					fmt.Println(index)
				}
			}
		}(i)
	}
	time.Sleep(time.Second)
	//1.14以前，由于调度模型，这里永远无法执行
	//1.14以后，会在gc的时候，发送信号抢占上面的goroutine
	fmt.Println("end")

}
