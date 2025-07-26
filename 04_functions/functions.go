// Go语言函数详解
package main

import "fmt"

// 1. 基本函数定义
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 2. 带返回值的函数
func add(a, b int) int {
	return a + b
}

// 3. 多返回值函数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// 4. 命名返回值
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 等同于 return area, perimeter
}

// 5. 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 6. 高阶函数 - 函数作为参数
func calculate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// 7. 闭包
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 8. 递归函数
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 9. 方法（属于某个类型的函数）
type Person struct {
	Name string
	Age  int
}

// 值接收者方法
func (p Person) introduce() {
	fmt.Printf("我是%s，今年%d岁\n", p.Name, p.Age)
}

// 指针接收者方法（可以修改结构体）
func (p *Person) birthday() {
	p.Age++
	fmt.Printf("%s生日快乐！现在%d岁了\n", p.Name, p.Age)
}

func main() {
	fmt.Println("=== Go语言函数 ===")
	
	// 1. 基本函数调用
	fmt.Println("\n1. 基本函数:")
	greet("小明")
	
	// 2. 带返回值的函数
	fmt.Println("\n2. 带返回值的函数:")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	// 3. 多返回值函数
	fmt.Println("\n3. 多返回值函数:")
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Printf("错误: %s\n", err)
	} else {
		fmt.Printf("10 ÷ 3 = %.2f\n", quotient)
	}
	
	// 处理除零错误
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("错误: %s\n", err)
	}
	
	// 4. 命名返回值
	fmt.Println("\n4. 命名返回值:")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("矩形面积: %.1f, 周长: %.1f\n", area, perimeter)
	
	// 5. 可变参数函数
	fmt.Println("\n5. 可变参数函数:")
	fmt.Printf("sum(1,2,3) = %d\n", sum(1, 2, 3))
	fmt.Printf("sum(1,2,3,4,5) = %d\n", sum(1, 2, 3, 4, 5))
	
	// 传递切片给可变参数函数
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("sum(切片) = %d\n", sum(numbers...))
	
	// 6. 高阶函数
	fmt.Println("\n6. 高阶函数:")
	multiply := func(a, b int) int { return a * b }
	subtract := func(a, b int) int { return a - b }
	
	fmt.Printf("calculate(10, 5, multiply) = %d\n", calculate(10, 5, multiply))
	fmt.Printf("calculate(10, 5, subtract) = %d\n", calculate(10, 5, subtract))
	
	// 7. 闭包
	fmt.Println("\n7. 闭包:")
	counter := makeCounter()
	fmt.Printf("第1次调用counter: %d\n", counter())
	fmt.Printf("第2次调用counter: %d\n", counter())
	fmt.Printf("第3次调用counter: %d\n", counter())
	
	// 8. 递归函数
	fmt.Println("\n8. 递归函数:")
	fmt.Printf("5! = %d\n", factorial(5))
	fmt.Printf("8! = %d\n", factorial(8))
	
	// 9. 方法
	fmt.Println("\n9. 方法:")
	person := Person{Name: "张三", Age: 25}
	person.introduce()
	person.birthday()
	person.introduce()
	
	// 10. 匿名函数
	fmt.Println("\n10. 匿名函数:")
	func(message string) {
		fmt.Printf("匿名函数说: %s\n", message)
	}("你好，Go!")
	
	// 11. defer语句（延迟执行）
	fmt.Println("\n11. defer语句:")
	defer fmt.Println("这行会在main函数结束时执行")
	defer fmt.Println("defer按栈的顺序执行（后进先出）")
	fmt.Println("正常执行的代码")
}
