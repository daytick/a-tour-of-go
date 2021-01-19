package main

import "fmt"

func main(){
	sum := 1
	for sum < 1000 { // 只使用条件语句，for 就相当于 while
		sum += sum
	}
	fmt.Println(sum)
}
