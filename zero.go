package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	/**
	 * 变量声明时没有显示地赋值将被赋予"零值"
	 * 数值类型零值为0
	 * 布尔类型零值为false
	 * 字符类型零值为""
	 */
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
