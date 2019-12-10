package singleton

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var WG sync.WaitGroup

func ITestSingleton(t *testing.T) {
	singletonObjA := GetInstance()
	singletonObjB := GetInstance()
	fmt.Printf("%p\n", &singletonObjA)
	fmt.Printf("%p\n", &singletonObjB)
	if &singletonObjA != &singletonObjB {
		t.Errorf("Singleton address not equal with one: %p, another: %p", &singletonObjA, &singletonObjB)
	}
}

func TestSingletonParallel(t *testing.T) {
	runtime.GOMAXPROCS(4)
	WG.Add(2)
	var aAddr, bAddr singleton
	go func(addr *singleton) {
		time.Sleep(3 * time.Second)
		singletonObj := GetInstance()
		addr = &singletonObj
		WG.Done()
	}(&aAddr)
	go func(addr *singleton) {
		time.Sleep(3 * time.Second)
		singletonObj := GetInstance()
		addr = &singletonObj
		WG.Done()
	}(&bAddr)
	WG.Wait()
	fmt.Printf("%p\n", &aAddr)
	fmt.Printf("%p\n", &bAddr)
	if aAddr != bAddr {
		t.Errorf("Singleton address not equal with one: %p, another: %p", &aAddr, &bAddr)
	}
}
