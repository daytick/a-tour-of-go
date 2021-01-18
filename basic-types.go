package main

import(
	"fmt"
	"math/cmplx"
)

var(
	ToBe bool     = false
	MaxInt uint64 = 1 << 64 - 1
	z complex128  = cmplx.Sqrt(-5 + 12i) // 复数
)

func main(){
	fmt.Printf("Type %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value %v\n", z, z)
}
/**
	bool // 布尔

	string // 字符串

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr // 整型，uintptr 是能存指针的整型
  // int，uint和uintptr类型在32位系统上通常为32位宽，在64位系统上为64位宽。

	byte // uint8别称

	rune // int32别称，代表一个 Unicode 码点

	float32 float64 // 浮点数

	complex64 complex128 // 复数
 */
