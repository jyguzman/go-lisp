package main

import "fmt"

type Cons struct {
	car any
	cdr *Cons
}

func NewCons(car any, cdr *Cons) *Cons {
	return &Cons{car, cdr}
}

func (c *Cons) print() {
	ptr := c
	for ptr != nil {
		if ptr.cdr != nil {
			fmt.Printf("%v -> ", ptr.car)
		} else {
			fmt.Printf("%v", ptr.car)
		}
		ptr = ptr.cdr
	}
}

func cons(car any, cdr *Cons) *Cons {
	return &Cons{car, cdr}
}
