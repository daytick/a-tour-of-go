package main

import "fmt"

func main(){
	var i, j int = 1, 2
	k := 3 // 在函数内部，可用 := 替代隐式 var 声明（函数外部不能使用）  
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
