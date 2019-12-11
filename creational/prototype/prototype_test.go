package prototype

import (
	"fmt"
	"testing"
)

func TestProtoType(t *testing.T) {
	sheep := &Sheep{Name: "mie", Age: 3}
	sheep1 := sheep.Clone().(*Sheep)
	fmt.Printf("original sheep addr: %p\n", sheep)
	fmt.Printf("cloned sheep addr: %p\n", sheep1)
	if sheep == sheep1 {
		t.Error("sheep are the same")
	}
}
