// Go语言接口详解
package main

import "fmt"

// 1. 基本接口定义
type Speaker interface {
	Speak() string
}

// 2. 多方法接口
type Animal interface {
	Speak() string
	Move() string
}

// 3. 接口组合
type Walker interface {
	Walk() string
}

type Runner interface {
	Run() string
}

type Athlete interface {
	Walker
	Runner
	Train() string
}

// 4. 实现接口的结构体
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s说: 汪汪!", d.Name)
}

func (d Dog) Move() string {
	return fmt.Sprintf("%s在跑步", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s说: 喵喵!", c.Name)
}

func (c Cat) Move() string {
	return fmt.Sprintf("%s在悄悄走动", c.Name)
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return fmt.Sprintf("%s说: 你好!", h.Name)
}

func (h Human) Walk() string {
	return fmt.Sprintf("%s在走路", h.Name)
}

func (h Human) Run() string {
	return fmt.Sprintf("%s在跑步", h.Name)
}

func (h Human) Train() string {
	return fmt.Sprintf("%s在训练", h.Name)
}

// 5. 空接口
func describe(i interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", i, i)
}

// 6. 类型断言
func checkType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("这是字符串: %s, 长度: %d\n", v, len(v))
	case int:
		fmt.Printf("这是整数: %d\n", v)
	case Dog:
		fmt.Printf("这是狗狗: %s\n", v.Name)
	case Cat:
		fmt.Printf("这是猫咪: %s\n", v.Name)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

// 7. 接口值
func printSpeaker(s Speaker) {
	if s == nil {
		fmt.Println("Speaker接口值为nil")
		return
	}
	fmt.Printf("Speaker说: %s\n", s.Speak())
}

// 8. 多态示例
func makeAllSpeak(speakers []Speaker) {
	for i, speaker := range speakers {
		fmt.Printf("%d. %s\n", i+1, speaker.Speak())
	}
}

func makeAllMove(animals []Animal) {
	for _, animal := range animals {
		fmt.Printf("🗣️  %s\n", animal.Speak())
		fmt.Printf("🚶 %s\n", animal.Move())
		fmt.Println()
	}
}

func main() {
	fmt.Println("=== Go语言接口 ===")

	// 1. 基本接口使用
	fmt.Println("\n1. 基本接口使用:")

	dog := Dog{Name: "旺财"}
	cat := Cat{Name: "咪咪"}

	// 接口变量可以保存任何实现了该接口的类型
	var speaker Speaker

	speaker = dog
	fmt.Println(speaker.Speak())

	speaker = cat
	fmt.Println(speaker.Speak())

	// 2. 多态演示
	fmt.Println("\n2. 多态演示:")
	speakers := []Speaker{
		Dog{Name: "小黄"},
		Cat{Name: "小白"},
		Dog{Name: "大黑"},
	}
	makeAllSpeak(speakers)

	// 3. 多方法接口
	fmt.Println("\n3. 多方法接口:")
	animals := []Animal{
		Dog{Name: "阿黄"},
		Cat{Name: "小花"},
	}
	makeAllMove(animals)

	// 4. 接口组合
	fmt.Println("\n4. 接口组合:")
	human := Human{Name: "小明"}

	var athlete Athlete = human
	fmt.Println(athlete.Walk())
	fmt.Println(athlete.Run())
	fmt.Println(athlete.Train())

	// 5. 空接口
	fmt.Println("\n5. 空接口:")
	describe(42)
	describe("Hello")
	describe(3.14)
	describe(Dog{Name: "Buddy"})
	describe([]int{1, 2, 3})

	// 6. 类型断言
	fmt.Println("\n6. 类型断言:")
	var empty interface{}

	empty = "Go语言"
	checkType(empty)

	empty = 42
	checkType(empty)

	empty = Dog{Name: "Lucky"}
	checkType(empty)

	empty = 3.14
	checkType(empty)

	// 7. 安全的类型断言
	fmt.Println("\n7. 安全的类型断言:")
	var i interface{} = "hello"

	// 不安全的断言（如果类型不匹配会panic）
	// s := i.(string)

	// 安全的断言
	if s, ok := i.(string); ok {
		fmt.Printf("成功断言为字符串: %s\n", s)
	} else {
		fmt.Println("断言失败")
	}

	if n, ok := i.(int); ok {
		fmt.Printf("成功断言为整数: %d\n", n)
	} else {
		fmt.Println("不是整数类型")
	}

	// 8. 接口值和nil
	fmt.Println("\n8. 接口值和nil:")
	var speaker1 Speaker
	printSpeaker(speaker1) // nil接口

	var speaker2 Speaker = Dog{Name: "Buddy"}
	printSpeaker(speaker2) // 非nil接口

	// 9. 接口的动态类型和动态值
	fmt.Println("\n9. 接口的内部结构:")
	var a interface{}
	fmt.Printf("空接口: 值=%v, 类型=%T\n", a, a)

	a = 42
	fmt.Printf("赋值后: 值=%v, 类型=%T\n", a, a)

	a = "hello"
	fmt.Printf("再次赋值: 值=%v, 类型=%T\n", a, a)

	// 10. 接口比较
	fmt.Println("\n10. 接口比较:")
	var s1, s2 Speaker

	s1 = Dog{Name: "A"}
	s2 = Dog{Name: "A"}

	fmt.Printf("s1 == s2: %t\n", s1 == s2) // 相同类型和值，返回true

	s2 = Dog{Name: "B"}
	fmt.Printf("s1 == s2: %t\n", s1 == s2) // 相同类型但不同值，返回false

	s2 = Cat{Name: "A"}
	fmt.Printf("s1 == s2: %t\n", s1 == s2) // 不同类型，返回false
}
