package main

import "fmt"

func main() {
	t := T{a: 1}
	t1 := Foo(t)
	t2 := Foo1(&t)
	fmt.Printf("t: %+v, t1: %+v, t2: %+v\n", t, t1, t2)
	t.a = 2
	fmt.Printf("t: %+v, t1: %+v, t2: %+v\n", t, t1, t2)

	t3 := &T{a: 1}
	t4 := Foo(*t3)
	t5 := Foo1(t3)
	fmt.Printf("t: %+v, t1: %+v, t2: %+v\n", t3, t4, t5)
	t3.a = 2
	fmt.Printf("t: %+v, t1: %+v, t2: %+v\n", t3, t4, t5)
}

// T ...
type T struct {
	a int
}

// Foo 参数 call by value
func Foo(t T) *T {
	t1 := &t
	return t1
}

// Foo1 参数 call by reference
func Foo1(t *T) *T {
	t1 := t
	return t1
}
