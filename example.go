package main

import "fmt"

/**
 * @Author: mf
 * @email: 18539271635@163.com
 * @Date: 2021/8/9 5:49 下午
 * @Desc:
 */

type Student struct {
	Age int
}

func foo(s int) *Student {
	a := new(Student)
	a.Age = s
	return a
}

func main() {
	a := foo(10)
	b := a.Age + 4
	c := b + 2
	fmt.Println(c)
}