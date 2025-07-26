// Package calculator 提供基本的数学计算功能
package calculator

import (
	"errors"
	"math"
)

// Add 加法运算
func Add(a, b float64) float64 {
	return a + b
}

// Subtract 减法运算
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply 乘法运算
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide 除法运算，返回结果和可能的错误
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// Power 计算 a 的 b 次方
func Power(a, b float64) float64 {
	return math.Pow(a, b)
}

// Sqrt 计算平方根
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("不能计算负数的平方根")
	}
	return math.Sqrt(a), nil
}

// Factorial 计算阶乘
func Factorial(n int) float64 {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	
	result := 1.0
	for i := 2; i <= n; i++ {
		result *= float64(i)
	}
	return result
}

// Abs 计算绝对值
func Abs(a float64) float64 {
	return math.Abs(a)
}

// Max 返回两个数中的较大值
func Max(a, b float64) float64 {
	return math.Max(a, b)
}

// Min 返回两个数中的较小值
func Min(a, b float64) float64 {
	return math.Min(a, b)
}
