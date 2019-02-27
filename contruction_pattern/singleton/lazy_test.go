package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	t.Run("test instance", func(t *testing.T) {
		got1 := GetInstance()
		got2 := GetInstance()
		if got1 != got2 {
			t.Errorf("Singleton GetInstance returns difference object!!!")
		}
	})

}

func TestGetInstanceAsync(t *testing.T) {
	t.Run("test get instance async", func(t *testing.T) {
		var channel1 = make(chan *Singleton)
		var channel2 = make(chan *Singleton)
		go AsyncGetInstance(channel1)
		go AsyncGetInstance(channel2)
		got1 := <-channel1
		got2 := <-channel2
		if got1 != got2 {
			t.Errorf("Singleton GetInstance returns difference object!!!")
		}
	})
}

func AsyncGetInstance(channel chan *Singleton) () {
	channel <- GetInstance()
	return
}
