package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	var mut *sync.RWMutex = &sync.RWMutex{}
	mut.RLock()
	var score = []int{0}
	mut.RUnlock()
	wg.Add(3)
	go func(wg *sync.WaitGroup,m *sync.RWMutex) {
		fmt.Println("One goroutine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg,mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two goroutine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg,mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup,m *sync.RWMutex) {
		// time.Sleep(1 * time.Millisecond)
		fmt.Println("Three goroutine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()
	fmt.Println(score)
}
