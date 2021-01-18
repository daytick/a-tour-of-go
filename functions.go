package main

import "fmt"

//类型放在变量后面，详见 https://blog.golang.org/declaration-syntax
func add(x int, y int) int{
	return x + y
}

func main(){
	fmt.Println(add(42, 13))
}
