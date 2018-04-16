package main

import "fmt"

//常量
const PI = 3.1415926
const s string = "hello, world!"
const index int = 100
const intvalue = 3 / 4

const (
	L = 1
	M = 2
	H = 3
)

const (
	Sunday    = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Color int

const (
	RED    Color = iota
	ORANGE
	YELLOW
	GREEN
	BLUE
	INDIGO
	VIOLET
)

//变量
var a int
var b bool
var str string

var (
	a_   int
	b_   bool
	str_ string
)

func main() {
	fmt.Println(s)
	fmt.Println(PI)
	fmt.Println(index)
	fmt.Println(intvalue)
}
