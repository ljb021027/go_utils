package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

//32位系统访问粒度是4字节（bytes），64位系统的是8字节
//cpu cache line 64字节
type Origin struct {
	a uint64
	b uint64
}

type WithPadding struct {
	a uint64
	_ [56]byte
	b uint64
	_ [56]byte
}

var num = 1000 * 1000

func OriginParallel() {
	var v Origin

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < num; i++ {
			atomic.AddUint64(&v.a, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < num; i++ {
			atomic.AddUint64(&v.b, 1)
		}
		wg.Done()
	}()

	wg.Wait()
	_ = v.a + v.b
}

func WithPaddingParallel() {
	var v WithPadding

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < num; i++ {
			atomic.AddUint64(&v.a, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < num; i++ {
			atomic.AddUint64(&v.b, 1)
		}
		wg.Done()
	}()

	wg.Wait()
	_ = v.a + v.b
}

func main() {

	fmt.Printf("arrange fields to reduce size:\n"+
		"T1 align: %d, size: %d\n"+
		"T2 align: %d, size: %d\n",
		unsafe.Alignof(Origin{}), unsafe.Sizeof(Origin{}),
		unsafe.Alignof(WithPadding{}), unsafe.Sizeof(WithPadding{}))

	var b time.Time

	b = time.Now()
	OriginParallel()
	fmt.Printf("OriginParallel. Cost=%+v.\n", time.Now().Sub(b))

	b = time.Now()
	WithPaddingParallel()
	fmt.Printf("WithPaddingParallel. Cost=%+v.\n", time.Now().Sub(b))
}
