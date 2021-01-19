package main

import "fmt"

const Pi = "3.14" // 使用 const 声明常量

func main(){
	const World = "世界" // 常量可以是数值、字符和布尔类型
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true // 不能使用 := 声明
	fmt.Println("Go rules?", Truth)
}
