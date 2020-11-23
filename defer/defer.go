package main

/*
defer 关键字收到的函数在执行时就已经确定了，
*/

import (
	"fmt"
)

func main() {
	fmt.Println(Foo())
	fmt.Println(Foo1())
	Bar()
	Bar1()
	Bar2()
}

// Foo 返回 10
// defer 关键字收到的是一个函数闭包，内部引用了 Foo 函数的返回值 "r"
// Foo 结尾返回的结果 5 表示 r = 5,在函数返回之前调用了 defer，
// 此时 r = 5, 执行 defer 的方法后， r = 10.
func Foo() (r int) {
	defer func() {
		r += 5
	}()
	return 5
}

// Foo1 返回 5
// defer 关键字接收到的函数需要参数 r，并把 Foo1 函数的返回值 r 传入。
// (Foo1 的返回值 r 和 defer 闭包中用到的 r 不是同一个)
// 此时 Foo1 的返回值 r 没有明确赋值，所以为 int 的默认值 0，所以传入
// defer 函数中的 r 的值为 0. 最后 Foo1 函数返回之前调用了 defer 的
// 匿名函数修改了匿名函数内部的 r 的值为 6，但是并没有改变 Foo1 函数的返回值。
func Foo1() (r int) {
	defer func(r int) {
		r += 5
	}(r)
	return 5
}

// T ...
type T struct {
	a int
}

// Bar ...
func Bar() {
	var t T
	defer fmt.Println(t)
	t = T{a: 1}
}

// Bar1 ...
func Bar1() {
	var t T
	defer func() {
		fmt.Println(t)
	}()
	t = T{a: 1}
}

// Bar2 ...
func Bar2() {
	var t *T
	defer fmt.Println(t)
	t = &T{a: 1}
}
