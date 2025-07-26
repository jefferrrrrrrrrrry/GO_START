// Go语言数据类型详解
package main

import "fmt"

func main() {
	fmt.Println("=== Go语言数据类型 ===")
	
	// 1. 基本数据类型
	fmt.Println("\n1. 基本数据类型:")
	
	// 整型
	var a int = 42
	var b int8 = 127        // -128 到 127
	var c uint = 42         // 无符号整型
	var d int64 = 9223372036854775807
	
	fmt.Printf("int: %d, int8: %d, uint: %d, int64: %d\n", a, b, c, d)
	
	// 浮点型
	var e float32 = 3.14
	var f float64 = 3.141592653589793
	
	fmt.Printf("float32: %.2f, float64: %.15f\n", e, f)
	
	// 布尔型
	var g bool = true
	var h bool = false
	
	fmt.Printf("bool1: %t, bool2: %t\n", g, h)
	
	// 字符串
	var i string = "Hello, 世界"
	fmt.Printf("string: %s, 长度: %d\n", i, len(i))
	
	// 2. 复合数据类型
	fmt.Println("\n2. 复合数据类型:")
	
	// 数组 - 固定长度
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("数组: %v\n", arr)
	
	// 切片 - 动态数组
	var slice []int = []int{1, 2, 3, 4, 5}
	slice = append(slice, 6) // 添加元素
	fmt.Printf("切片: %v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))
	
	// 映射 - 类似字典
	var m map[string]int = make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3
	fmt.Printf("映射: %v\n", m)
	
	// 3. 指针
	fmt.Println("\n3. 指针:")
	var num int = 42
	var ptr *int = &num  // 获取地址
	fmt.Printf("值: %d, 地址: %p, 指针指向的值: %d\n", num, ptr, *ptr)
	
	// 4. 类型转换
	fmt.Println("\n4. 类型转换:")
	var x int = 42
	var y float64 = float64(x)  // 显式转换
	var z string = fmt.Sprintf("%d", x)  // 转为字符串
	
	fmt.Printf("int: %d, float64: %.1f, string: %s\n", x, y, z)
	
	// 5. 零值
	fmt.Println("\n5. 零值:")
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	var zeroSlice []int
	
	fmt.Printf("int零值: %d, float零值: %.1f, bool零值: %t\n", zeroInt, zeroFloat, zeroBool)
	fmt.Printf("string零值: '%s', slice零值: %v\n", zeroString, zeroSlice)
}
