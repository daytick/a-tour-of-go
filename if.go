package main

import(
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 { // 类似 for 语句，if 语句也不用括号
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main(){
	fmt.Println(sqrt(2), sqrt(-4))
}
