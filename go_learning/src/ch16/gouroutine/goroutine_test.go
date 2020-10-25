package gouroutine_test

import (
	"fmt"
	"testing"
	"time"
)

// 协程被调用的时候方法的调度顺序不一定是一致的
func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}
