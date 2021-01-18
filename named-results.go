package main

import "fmt"

func split(sum int) (x, y int) { // 返回值可以命名
	x = sum * 4 / 9
	y = sum - x
	return // 裸返回（无参数return）返回命名的返回值
}

func main(){
	fmt.Println(split(17))
}
