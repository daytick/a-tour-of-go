package main

import "fmt"

func main(){
	sum := 1
	for ; sum < 1000; { // 可以只保留条件语句
		sum += sum
	}
	fmt.Println(sum)
}
