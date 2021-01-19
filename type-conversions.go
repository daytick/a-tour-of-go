package main

import(
	"fmt"
	"math"
)

func main(){
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // 赋值给不同类型变量需要显示的类型转换
	var z uint = uint(f) // T(v) 将 v 转换成 T 类型
	fmt.Println(x, y, z)
}
