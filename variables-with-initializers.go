package main

import "fmt"

var i, j int = 1, 2 // var声明同时初始化

func main(){
	var c, python, java = true, false, "no!" // 变量类型可由声明时初始化值确定
	fmt.Println(i, j, c, python, java)
}
