package singleton

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock     sync.Mutex
	instance *Singleton
)

func GetInstance() (*Singleton) {
	fmt.Println("enter get instance")
	if instance == nil {
		// 为测试方便，使用sleep方法让出cpu
		time.Sleep(time.Second)
		fmt.Println("enter instance == nil")
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("init obj")
			instance = &Singleton{}
		}
	}
	return instance
}
