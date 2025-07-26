// Go语言实战练习
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 练习1: 学生管理系统
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade float64
}

type StudentManager struct {
	students []Student
	nextID   int
}

func NewStudentManager() *StudentManager {
	return &StudentManager{
		students: make([]Student, 0),
		nextID:   1,
	}
}

func (sm *StudentManager) AddStudent(name string, age int, grade float64) {
	student := Student{
		ID:    sm.nextID,
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	sm.students = append(sm.students, student)
	sm.nextID++
	fmt.Printf("添加学生成功: %s (ID: %d)\n", name, student.ID)
}

func (sm *StudentManager) FindStudent(id int) (*Student, error) {
	for i := range sm.students {
		if sm.students[i].ID == id {
			return &sm.students[i], nil
		}
	}
	return nil, fmt.Errorf("未找到ID为%d的学生", id)
}

func (sm *StudentManager) ListAllStudents() {
	if len(sm.students) == 0 {
		fmt.Println("暂无学生信息")
		return
	}
	
	fmt.Println("\n=== 学生列表 ===")
	fmt.Printf("%-4s %-10s %-4s %-6s\n", "ID", "姓名", "年龄", "成绩")
	fmt.Println(strings.Repeat("-", 30))
	
	for _, student := range sm.students {
		fmt.Printf("%-4d %-10s %-4d %-6.1f\n", 
			student.ID, student.Name, student.Age, student.Grade)
	}
}

func (sm *StudentManager) GetAverageGrade() float64 {
	if len(sm.students) == 0 {
		return 0
	}
	
	total := 0.0
	for _, student := range sm.students {
		total += student.Grade
	}
	return total / float64(len(sm.students))
}

// 练习2: 简单的并发下载器
func downloadFile(url string, id int) {
	fmt.Printf("开始下载文件 %d: %s\n", id, url)
	
	// 模拟下载时间
	downloadTime := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(downloadTime)
	
	fmt.Printf("文件 %d 下载完成，耗时: %v\n", id, downloadTime)
}

func concurrentDownloader(urls []string) {
	fmt.Println("\n=== 并发下载器演示 ===")
	
	start := time.Now()
	done := make(chan bool, len(urls))
	
	// 启动并发下载
	for i, url := range urls {
		go func(url string, id int) {
			downloadFile(url, id)
			done <- true
		}(url, i+1)
	}
	
	// 等待所有下载完成
	for i := 0; i < len(urls); i++ {
		<-done
	}
	
	fmt.Printf("所有下载完成，总耗时: %v\n", time.Since(start))
}

// 练习3: 简单的计算器
type Calculator struct{}

func (c Calculator) Calculate(expression string) (float64, error) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return 0, fmt.Errorf("表达式格式错误，应为: 数字 操作符 数字")
	}
	
	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("第一个数字格式错误: %v", err)
	}
	
	operator := parts[1]
	
	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("第二个数字格式错误: %v", err)
	}
	
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("除数不能为零")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("不支持的操作符: %s", operator)
	}
}

// 练习4: 猜数字游戏
func guessNumberGame() {
	fmt.Println("\n=== 猜数字游戏 ===")
	fmt.Println("我想了一个1-100之间的数字，你来猜猜看！")
	
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 0
	maxAttempts := 7
	
	reader := bufio.NewReader(os.Stdin)
	
	for attempts < maxAttempts {
		fmt.Printf("第%d次猜测 (剩余%d次): ", attempts+1, maxAttempts-attempts)
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误:", err)
			continue
		}
		
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("请输入有效的数字！")
			continue
		}
		
		attempts++
		
		if guess == target {
			fmt.Printf("🎉 恭喜你！猜对了！数字就是 %d\n", target)
			fmt.Printf("你用了 %d 次猜测\n", attempts)
			return
		} else if guess < target {
			fmt.Println("太小了！再试试更大的数字")
		} else {
			fmt.Println("太大了！再试试更小的数字")
		}
	}
	
	fmt.Printf("😢 游戏结束！正确答案是 %d\n", target)
}

// 练习5: 文件处理器
type FileProcessor struct{}

func (fp FileProcessor) CountWords(content string) map[string]int {
	wordCount := make(map[string]int)
	
	// 简单的分词（按空格分割）
	words := strings.Fields(strings.ToLower(content))
	
	for _, word := range words {
		// 去除标点符号
		word = strings.Trim(word, ".,!?;:")
		if word != "" {
			wordCount[word]++
		}
	}
	
	return wordCount
}

func (fp FileProcessor) FindMostFrequentWord(wordCount map[string]int) (string, int) {
	if len(wordCount) == 0 {
		return "", 0
	}
	
	maxWord := ""
	maxCount := 0
	
	for word, count := range wordCount {
		if count > maxCount {
			maxWord = word
			maxCount = count
		}
	}
	
	return maxWord, maxCount
}

func demonstrateFileProcessor() {
	fmt.Println("\n=== 文件处理器演示 ===")
	
	text := `Go语言是Google开发的开源编程语言。Go语言简洁、高效、并发。
	Go语言适合构建网络服务。许多公司使用Go语言开发微服务。
	学习Go语言让编程变得更加有趣。`
	
	fp := FileProcessor{}
	wordCount := fp.CountWords(text)
	
	fmt.Println("词频统计:")
	for word, count := range wordCount {
		fmt.Printf("  %s: %d\n", word, count)
	}
	
	mostFrequent, count := fp.FindMostFrequentWord(wordCount)
	fmt.Printf("\n出现最多的词: '%s' (出现 %d 次)\n", mostFrequent, count)
}

func interactiveMenu() {
	fmt.Println("\n=== 交互式菜单 ===")
	fmt.Println("选择你想要尝试的练习:")
	fmt.Println("1. 学生管理系统")
	fmt.Println("2. 并发下载器")
	fmt.Println("3. 计算器")
	fmt.Println("4. 猜数字游戏")
	fmt.Println("5. 文件处理器")
	fmt.Println("0. 退出")
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("\n请选择 (0-5): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误:", err)
			continue
		}
		
		choice := strings.TrimSpace(input)
		
		switch choice {
		case "1":
			demonstrateStudentManager()
		case "2":
			urls := []string{
				"https://example.com/file1.zip",
				"https://example.com/file2.zip", 
				"https://example.com/file3.zip",
				"https://example.com/file4.zip",
			}
			concurrentDownloader(urls)
		case "3":
			demonstrateCalculator()
		case "4":
			guessNumberGame()
		case "5":
			demonstrateFileProcessor()
		case "0":
			fmt.Println("感谢使用！再见!")
			return
		default:
			fmt.Println("无效选择，请输入 0-5")
		}
	}
}

func demonstrateStudentManager() {
	fmt.Println("\n=== 学生管理系统演示 ===")
	
	sm := NewStudentManager()
	
	// 添加一些示例学生
	sm.AddStudent("张三", 20, 85.5)
	sm.AddStudent("李四", 19, 92.0)
	sm.AddStudent("王五", 21, 78.5)
	sm.AddStudent("赵六", 20, 88.0)
	
	// 显示所有学生
	sm.ListAllStudents()
	
	// 查找学生
	student, err := sm.FindStudent(2)
	if err != nil {
		fmt.Printf("查找失败: %v\n", err)
	} else {
		fmt.Printf("\n找到学生: %s, 年龄: %d, 成绩: %.1f\n", 
			student.Name, student.Age, student.Grade)
	}
	
	// 计算平均成绩
	avgGrade := sm.GetAverageGrade()
	fmt.Printf("\n班级平均成绩: %.1f\n", avgGrade)
}

func demonstrateCalculator() {
	fmt.Println("\n=== 计算器演示 ===")
	
	calc := Calculator{}
	expressions := []string{
		"10 + 5",
		"20 - 8",
		"6 * 7",
		"15 / 3",
		"10 / 0", // 错误示例
		"abc + 5", // 错误示例
	}
	
	for _, expr := range expressions {
		result, err := calc.Calculate(expr)
		if err != nil {
			fmt.Printf("%s = 错误: %v\n", expr, err)
		} else {
			fmt.Printf("%s = %.2f\n", expr, result)
		}
	}
}

func main() {
	fmt.Println("🚀 欢迎来到Go语言实战练习！")
	
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	
	// 启动交互式菜单
	interactiveMenu()
}
