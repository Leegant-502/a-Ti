package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fffg()
	fmt.Println(Sqrt(2))

}

func fffg() {
	// 声明一个长度为 5 的整型数组
	var numbers [5]int
	// 初始化数组
	numbers = [5]int{1, 2, 3, 4, 5}
	// 简化声明和初始化
	colors := [3]string{"red", "green", "blue"}
	fmt.Println(numbers) // 输出: [1 2 3 4 5]
	fmt.Println(colors)  // 输出: [red green blue]

}
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
