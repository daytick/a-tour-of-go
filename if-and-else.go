package main

import(
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // if 短语句中声明的变量在 else 中可见
	}
	return lim
}

func main(){
	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))
}
