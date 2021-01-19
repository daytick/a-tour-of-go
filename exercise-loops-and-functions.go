package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for ; math.Abs(z - attempt(x, z)) > 1e-15; { // 值停止变化或变化微小时停止循环
		z = attempt(x, z)
		fmt.Println(z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z := 1.0 // 初始化一个浮点型变量
	for i := 0; i < 10; i++ {
		z = attempt(x, z)
		fmt.Println(z)
	}
	return z
}

// 牛顿法求平方根
func attempt(x, z float64) float64 {
	return z - (z*z - x) / (2*z)
}

func main(){
	fmt.Println(Sqrt(2))
}
