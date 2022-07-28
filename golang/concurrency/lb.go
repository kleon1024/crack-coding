package main

import "fmt"

func main() {
	ips := []int{1, 2, 3}
	lb := NewLB(ips)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", lb.Next())
	}
}

type LB struct {
	ips   []int
	count int
	size  int
}

func NewLB(ips []int) *LB {
	return &LB{
		ips:   ips,
		count: 0,
		size:  len(ips),
	}
}

func (lb *LB) Next() int {
	lb.count += 1
	lb.count %= lb.size
	return lb.ips[lb.count]
}

func (lb *LB) Next1() int {
	lb.count += 1
	if lb.count == lb.size {
		lb.count = 0
	}
	return lb.ips[lb.count]
}
