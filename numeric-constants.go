package main

import "fmt"

const(
	Big = 1 << 100 // 无类型常量根据上下文需要变成不同类型
	Small = Big >> 99 // 指定类型的常量举例：const v float64 = 1
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main(){
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
