package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	var result interface{}
	result = "3"
	_, ok := result.(string)
	fmt.Println(ok)
	fmt.Println(result == "3")

}

func TestLock(t *testing.T) {
	lock := sync.RWMutex{}
	lock.RLock()
	lock.RLock()

	defer lock.RUnlock()
	lock.Lock()
	defer lock.Unlock()

	fmt.Println(11)

}
