package main

import "fmt"

func main(){
	sum := 0
	// Go 循环只有 for，且不含括号
	for i := 0; i < 10; i++ { // 初始化语句、条件表达式、每轮循环结束语句
		sum += i
	}
	fmt.Println(sum)
}
