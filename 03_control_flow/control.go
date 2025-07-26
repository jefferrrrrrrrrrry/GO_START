// Go语言控制流
package main

import "fmt"

func main() {
	fmt.Println("=== Go语言控制流 ===")
	
	// 1. if条件语句
	fmt.Println("\n1. if条件语句:")
	age := 20
	
	if age >= 18 {
		fmt.Println("已成年")
	} else if age >= 13 {
		fmt.Println("青少年")
	} else {
		fmt.Println("儿童")
	}
	
	// if语句中可以包含初始化语句
	if score := 85; score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("需要努力")
	}
	
	// 2. switch语句
	fmt.Println("\n2. switch语句:")
	day := "Monday"
	
	switch day {
	case "Monday":
		fmt.Println("周一，新的开始！")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("工作日")
	case "Friday":
		fmt.Println("周五，快周末了！")
	case "Saturday", "Sunday":
		fmt.Println("周末，休息时间！")
	default:
		fmt.Println("未知的日期")
	}
	
	// switch无表达式（类似if-else链）
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("冰点以下")
	case temperature < 20:
		fmt.Println("凉爽")
	case temperature < 30:
		fmt.Println("温暖")
	default:
		fmt.Println("炎热")
	}
	
	// 3. for循环
	fmt.Println("\n3. for循环:")
	
	// 标准for循环
	fmt.Print("计数: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// for range遍历切片
	fruits := []string{"苹果", "香蕉", "橙子"}
	fmt.Println("水果列表:")
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index, fruit)
	}
	
	// for range遍历映射
	scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}
	fmt.Println("成绩单:")
	for name, score := range scores {
		fmt.Printf("%s: %d分\n", name, score)
	}
	
	// while风格的for循环
	fmt.Print("倒计时: ")
	count := 5
	for count > 0 {
		fmt.Printf("%d ", count)
		count--
	}
	fmt.Println("发射！")
	
	// 无限循环（需要break跳出）
	fmt.Print("寻找偶数: ")
	number := 1
	for {
		if number%2 == 0 {
			fmt.Printf("找到偶数: %d\n", number)
			break
		}
		number++
		if number > 10 { // 防止无限循环
			break
		}
	}
	
	// 4. 跳转语句
	fmt.Println("\n4. 跳转语句:")
	
	// continue - 跳过当前迭代
	fmt.Print("奇数: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 跳过偶数
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// break - 跳出循环
	fmt.Print("寻找5: ")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Printf("找到了%d！", i)
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// 5. 标签和goto（谨慎使用）
	fmt.Println("\n5. 嵌套循环和标签:")
	
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("在(%d,%d)处跳出外层循环\n", i, j)
				break outer
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
}
