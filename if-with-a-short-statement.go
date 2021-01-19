package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // if 语句也可以在条件语句前执行一条语句
		return v
	}
	return lim
}

func main(){
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
}
