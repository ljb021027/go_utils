package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/juju/ratelimit"
)

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	bucket := ratelimit.NewBucket(1*time.Second, 1)

	for i := 0; i < 100; i++ {
		//take := bucket.Take(1)
		//fmt.Println(take)

		//available := bucket.TakeAvailable(1)
		//fmt.Println(available)

		bucket.Wait(1)
		fmt.Println(time.Now())

		bucket.WaitMaxDuration()
	}

	waitGroup.Wait()

}
