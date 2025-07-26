// Package utils 提供常用的工具函数
package utils

import "fmt"

// Greeting 生成问候语
func Greeting(name string) string {
	return fmt.Sprintf("你好, %s! 欢迎使用Go语言!", name)
}

// Max 返回整数切片中的最大值
func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// Min 返回整数切片中的最小值
func Min(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

// Average 计算整数切片的平均值
func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// IsEven 判断数字是否为偶数
func IsEven(n int) bool {
	return n%2 == 0
}

// Reverse 反转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
