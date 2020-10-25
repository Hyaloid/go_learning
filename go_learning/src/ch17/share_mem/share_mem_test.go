package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 不太安全，因为有线程的竞争
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 安全状况
func TestCouterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func TestCh(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 1)
	go sum(s[:len(s)/2], c) // 17
	time.Sleep(time.Second * 1)
	go sum(s[len(s)/2:], c) // -5
	go sum(s[0:2], c) // 9
	x := <-c
	time.Sleep(time.Second * 1)
	y := <-c
	t.Log(x, y)
}
