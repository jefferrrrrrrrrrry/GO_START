// Go语言包管理示例
package main

import (
	"math"
	"strings"
	"time"
	
	// 导入同项目中的包
	"go-learn/08_packages/utils"
	"go-learn/08_packages/calculator"
)

// 别名导入
import (
	str "strings" // 给strings包起别名
	. "fmt"      // 点导入，可以直接使用Println而不需要fmt.
)

func main() {
	Println("=== Go语言包管理 ===")
	
	// 1. 标准库包的使用
	Println("\n1. 标准库包:")
	
	// math包
	Printf("π = %.2f\n", math.Pi)
	Printf("√16 = %.0f\n", math.Sqrt(16))
	Printf("2^10 = %.0f\n", math.Pow(2, 10))
	
	// strings包
	text := "Hello, Go World!"
	Println("原文本:", text)
	Println("转大写:", strings.ToUpper(text))
	Println("转小写:", strings.ToLower(text))
	Println("包含'Go':", strings.Contains(text, "Go"))
	Println("替换:", strings.Replace(text, "World", "Gopher", -1))
	
	// 使用别名
	words := str.Split(text, " ")
	Println("分割后:", words)
	
	// time包
	now := time.Now()
	Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))
	
	// 2. 自定义包的使用
	Println("\n2. 自定义包:")
	
	// 使用utils包
	Println("工具函数演示:")
	Println(utils.Greeting("小明"))
	
	numbers := []int{1, 2, 3, 4, 5}
	Printf("数组 %v 的最大值: %d\n", numbers, utils.Max(numbers))
	Printf("数组 %v 的最小值: %d\n", numbers, utils.Min(numbers))
	Printf("数组 %v 的平均值: %.2f\n", numbers, utils.Average(numbers))
	
	// 使用calculator包
	Println("\n计算器演示:")
	Printf("10 + 5 = %.2f\n", calculator.Add(10, 5))
	Printf("10 - 5 = %.2f\n", calculator.Subtract(10, 5))
	Printf("10 × 5 = %.2f\n", calculator.Multiply(10, 5))
	
	result, err := calculator.Divide(10, 5)
	if err != nil {
		Printf("错误: %s\n", err)
	} else {
		Printf("10 ÷ 5 = %.2f\n", result)
	}
	
	// 测试除零错误
	_, err = calculator.Divide(10, 0)
	if err != nil {
		Printf("捕获到错误: %s\n", err)
	}
	
	// 高级计算
	Printf("5! = %.0f\n", calculator.Factorial(5))
	Printf("2^8 = %.0f\n", calculator.Power(2, 8))
}
