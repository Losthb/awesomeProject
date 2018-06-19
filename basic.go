package main

import (
	"fmt"
	"math"
)

//函数外定义变量，函数外必须用var 不可以短定义（包内变量）
var (
	aa = 3
	ss = "zdy"
	bb = true
)

//main函数入口
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	consts()
	enums()
	enums1()
}

//空值定义变量
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q", a, s)
}

//赋值定义变量
func variableInitiaValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//组定义变量
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//短定义变量：
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

/*
常量定义
const filename = "abc.txt"
const 数值可作为各种类型使用
const a,b = 3,4
var c int = int(math.Sqrt(a*a + b*b))
 */
func consts() {
	const (
		filename = "test.txt"
		a        = 3
		b        = 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

/*
常量枚举类型
 */
func enums() {
	const (
		c      = iota
		_
		python
		golang
		nodejs
	)
	fmt.Println(c, python, golang, nodejs)
}

//iota为自增变量，记住自增变量从0开始
func enums1() {
	const (
		mb = 1 << (10 * iota)
		gb
		tb
		pb
	)
	fmt.Print(mb, gb, tb, pb)
}
