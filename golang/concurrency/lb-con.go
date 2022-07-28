package main

import (
	"fmt"
	"sync"
)

func main() {
	ips := []int{1, 2, 3}
	lb := NewLB(ips)
	wg := &sync.WaitGroup{}
	for j := 0; j < 3; j++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Printf("%v\n", lb.Next())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

type LB struct {
	ips   []int
	count int
	size  int
	m     *sync.Mutex
}

func NewLB(ips []int) *LB {
	return &LB{
		ips:   ips,
		count: 0,
		size:  len(ips),
		m:     &sync.Mutex{},
	}
}

func (lb *LB) Next() int {
	lb.m.Lock()
	defer lb.m.Unlock()

	lb.count += 1
	if lb.count == lb.size {
		lb.count = 0
	}
	return lb.ips[lb.count]
}
