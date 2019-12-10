package singleton

import "sync"

// lazy
var singletonObj singleton
var once sync.Once

type singleton struct {
}

func GetInstance() singleton {
	once.Do(func() {
		singletonObj = singleton{}
	})
	return singletonObj
}

// hungry
// use func init(), no need to use once.Do