package util

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"sync"
	"testing"
	"time"
)

//etcd分布式锁1
func getLock(path string) sync.Locker {
	config := clientv3.Config{
		Endpoints:   []string{"10.10.134.30:2379", "10.10.134.31:2379", "10.10.134.32:2379"},
		DialTimeout: time.Duration(10) * time.Second,
	}
	client, _ := clientv3.New(config)
	session, err := concurrency.NewSession(client, concurrency.WithTTL(1))
	if err != nil {
		fmt.Println(err)
	}
	return concurrency.NewLocker(session, path)
}

func TestLock1(t *testing.T) {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			fmt.Println(time.Now())
		}
	}()
	lock := getLock("com/testkety")
	lock.Lock()
	fmt.Println(1)
	time.Sleep(30 * time.Second)
	fmt.Println(2)
	defer lock.Unlock()
}

func TestLock2(t *testing.T) {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			fmt.Println(time.Now())
		}
	}()
	lock := getLock("com/testkety")
	lock.Lock()
	fmt.Println(3)
	defer lock.Unlock()

}
