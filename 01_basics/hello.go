// Go语言基础语法
package main

import "fmt"

func main() {
	// 1. Hello World
	fmt.Println("Hello, Go!")

	// 2. 变量声明的几种方式
	// 方式1: var关键字
	var name string = "Go语言"
	var age int = 14

	// 方式2: 类型推断
	var language = "Go"

	// 方式3: 短声明（只能在函数内使用）
	version := "1.21"

	fmt.Printf("语言: %s, 年龄: %d年, 当前语言: %s, 版本: %s\n", name, age, language, version)

	// 3. 常量
	const pi = 3.14159
	const greeting = "你好"

	fmt.Printf("%s, π = %.2f\n", greeting, pi)

	// 4. 多重赋值
	x, y := 10, 20
	fmt.Printf("x = %d, y = %d\n", x, y)

	// 5. 交换变量
	x, y = y, x
	fmt.Printf("交换后: x = %d, y = %d\n", x, y)
}
