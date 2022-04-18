package main

import "fmt"

func main() {
	n := 123
	p := &n
	q := &p
	m := 10000
	p2 := &m
	q2 := &p2
	fmt.Printf("p value:%d, address%p\n", *p, p)
	fmt.Printf("q value:%d, address%p\n", **q, *q)
	fmt.Printf("p2 value:%d, address%p\n", *p2, p2)
	fmt.Printf("q2 value:%d, address%p\n", **q2, *q2)
	pb := p
	p = p2
	p2 = pb
	fmt.Printf("p value:%d, address%p\n", *p, p)
	fmt.Printf("p2 value:%d, address%p\n", *p2, p2)
}
