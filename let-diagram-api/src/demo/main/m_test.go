package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestA(t *testing.T) {
	dict := sync.Map{}
	dict.Store(1, 1)
	v, stored := dict.LoadOrStore(1, 2)
	v2, stored2 := dict.LoadOrStore(2, 2)
	fmt.Println(v, stored)    // 1 true
	fmt.Println(v2, stored2)  // 2 false
	fmt.Println(dict.Load(1)) // 1
	fmt.Println(dict.Load(2)) // 2
}
