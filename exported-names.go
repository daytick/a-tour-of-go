package main

import(
	"fmt"
	"math"
)

func main(){
	// fmt.Println(math.pi) 只能引用导入的包中的已导出名称
	fmt.Println(math.Pi) // 导出名称首字母均为大写
}
