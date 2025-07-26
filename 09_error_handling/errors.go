// Go语言错误处理详解
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// 1. 自定义错误类型
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("验证错误: 字段 '%s' %s", e.Field, e.Message)
}

// 2. 多种创建错误的方式
func createErrors() {
	fmt.Println("=== 创建错误的方式 ===")

	// 方式1: errors.New()
	err1 := errors.New("这是一个简单错误")
	fmt.Printf("errors.New(): %v\n", err1)

	// 方式2: fmt.Errorf()
	user := "admin"
	err2 := fmt.Errorf("用户 %s 不存在", user)
	fmt.Printf("fmt.Errorf(): %v\n", err2)

	// 方式3: 自定义错误类型
	err3 := ValidationError{
		Field:   "email",
		Message: "格式不正确",
	}
	fmt.Printf("自定义错误: %v\n", err3)
}

// 3. 返回错误的函数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

func parseInt(s string) (int, error) {
	if s == "" {
		return 0, ValidationError{
			Field:   "input",
			Message: "不能为空",
		}
	}

	result, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("无法解析 '%s' 为整数: %w", s, err)
	}

	return result, nil
}

// 4. 错误处理的不同模式
func demonstrateErrorHandling() {
	fmt.Println("\n=== 错误处理模式 ===")

	// 模式1: 基本错误检查
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("10 ÷ 2 = %.2f\n", result)

	// 模式2: 错误链式检查
	if result, err := divide(10, 0); err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("结果: %.2f\n", result)
	}

	// 模式3: 多个可能的错误
	num, err := parseInt("123")
	if err != nil {
		// 类型断言检查具体错误类型
		if ve, ok := err.(ValidationError); ok {
			fmt.Printf("验证错误: %s\n", ve.Message)
		} else {
			fmt.Printf("解析错误: %v\n", err)
		}
		return
	}
	fmt.Printf("解析结果: %d\n", num)

	// 测试空字符串
	_, err = parseInt("")
	if err != nil {
		fmt.Printf("空字符串错误: %v\n", err)
	}

	// 测试无效字符串
	_, err = parseInt("abc")
	if err != nil {
		fmt.Printf("无效字符串错误: %v\n", err)
	}
}

// 5. panic和recover
func demonstratePanicRecover() {
	fmt.Println("\n=== Panic和Recover ===")

	// recover必须在defer函数中调用
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到panic: %v\n", r)
		}
	}()

	fmt.Println("正常执行...")

	// 模拟一个会导致panic的情况
	slice := []int{1, 2, 3}
	fmt.Printf("访问有效索引: %d\n", slice[1])

	// 这行会导致panic
	fmt.Printf("访问无效索引: %d\n", slice[10])

	// 这行不会执行
	fmt.Println("这行不会被打印")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("安全函数捕获到panic: %v\n", r)
		}
	}()

	panic("模拟的panic")
}

// 6. 文件操作中的错误处理
func fileOperationExample() {
	fmt.Println("\n=== 文件操作错误处理 ===")

	// 尝试读取不存在的文件
	content, err := os.ReadFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("读取文件错误: %v\n", err)

		// 检查具体的错误类型
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}
	} else {
		fmt.Printf("文件内容: %s\n", content)
	}

	// 创建文件并写入内容
	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Printf("创建文件错误: %v\n", err)
		return
	}
	defer func() {
		// 确保文件被关闭
		if err := file.Close(); err != nil {
			fmt.Printf("关闭文件错误: %v\n", err)
		}
		// 删除临时文件
		if err := os.Remove("temp.txt"); err != nil {
			fmt.Printf("删除文件错误: %v\n", err)
		}
	}()

	_, err = file.WriteString("Hello, Go错误处理!")
	if err != nil {
		fmt.Printf("写入文件错误: %v\n", err)
		return
	}

	fmt.Println("文件操作成功完成")
}

// 7. 错误包装和解包
func errorWrappingExample() {
	fmt.Println("\n=== 错误包装 ===")

	// 创建一个基础错误
	baseErr := errors.New("数据库连接失败")

	// 包装错误
	wrappedErr := fmt.Errorf("无法获取用户信息: %w", baseErr)

	fmt.Printf("包装的错误: %v\n", wrappedErr)

	// 解包错误
	if errors.Is(wrappedErr, baseErr) {
		fmt.Println("错误链中包含数据库连接错误")
	}

	// 使用errors.Unwrap
	unwrapped := errors.Unwrap(wrappedErr)
	fmt.Printf("解包后的错误: %v\n", unwrapped)
}

// 8. 最佳实践示例
func bestPracticeExample() {
	fmt.Println("\n=== 错误处理最佳实践 ===")

	// 1. 错误信息应该清晰具体
	validateUser := func(username string) error {
		if username == "" {
			return errors.New("用户名不能为空")
		}
		if len(username) < 3 {
			return fmt.Errorf("用户名长度不能少于3个字符，当前长度: %d", len(username))
		}
		return nil
	}

	// 2. 使用哨兵错误
	var (
		ErrUserNotFound = errors.New("用户未找到")
		ErrInvalidInput = errors.New("输入无效")
	)

	findUser := func(id int) error {
		if id <= 0 {
			return ErrInvalidInput
		}
		if id == 999 {
			return ErrUserNotFound
		}
		return nil
	}

	// 测试
	if err := validateUser("ab"); err != nil {
		fmt.Printf("用户验证失败: %v\n", err)
	}

	err := findUser(999)
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("处理用户不存在的情况")
	}

	// 3. 优雅降级
	getValue := func() (string, error) {
		// 模拟可能失败的操作
		return "", errors.New("服务暂时不可用")
	}

	value, err := getValue()
	if err != nil {
		fmt.Printf("获取值失败: %v，使用默认值\n", err)
		value = "默认值"
	}
	fmt.Printf("最终值: %s\n", value)
}

func main() {
	fmt.Println("=== Go语言错误处理 ===")

	createErrors()
	demonstrateErrorHandling()
	demonstratePanicRecover()

	// 演示安全的panic处理
	fmt.Println("\n--- 安全的panic处理 ---")
	safeFunction()
	fmt.Println("主程序继续执行")

	fileOperationExample()
	errorWrappingExample()
	bestPracticeExample()
}
