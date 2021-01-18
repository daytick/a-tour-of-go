package main

import "fmt"

func swap(x, y string) (string, string) { // 函数可以有任意个返回值
	return y, x
}

func main(){
	a, b := swap("hello", "world") // := 等价于 声明并赋值
	fmt.Println(a, b)
}
