package main

import "fmt"

func main(){
	v := 42 // 在声明变量时没有指定显示类型时，将从右侧的值推断类型
	fmt.Printf("v is of type %T\n", v)
	// 对于没有指定类型的数值常量，将根据其精度推断类型
}
