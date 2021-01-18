package main

import "fmt"

func add(x, y int) int { // 同类型变量最后一个声明类型即可
	 return x + y
}

func main(){
	fmt.Println(add(42, 13))
}
