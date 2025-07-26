// Go语言结构体与方法
package main

import (
	"fmt"
	"unsafe"
	"encoding/json"
)

// 1. 基本结构体定义
type Student struct {
	Name   string
	Age    int
	Grade  float64
	School string
}

// 2. 嵌入字段的结构体
type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Employee struct {
	Name     string
	Position string
	Salary   float64
	Address  // 嵌入字段
}

// 3. 方法定义

// 值接收者方法 - 不会修改原结构体
func (s Student) introduce() {
	fmt.Printf("我是%s，%d岁，成绩%.1f分，就读于%s\n",
		s.Name, s.Age, s.Grade, s.School)
}

// 指针接收者方法 - 可以修改原结构体
func (s *Student) updateGrade(newGrade float64) {
	s.Grade = newGrade
	fmt.Printf("%s的成绩更新为%.1f分\n", s.Name, s.Grade)
}
// 计算方法
func (s Student) isExcellent() bool {
	return s.Grade >= 90
}

// Employee的方法
func (e Employee) getFullInfo() string {
	return fmt.Sprintf("%s - %s, 薪资: %.0f, 地址: %s, %s %s",
		e.Name, e.Position, e.Salary, e.Street, e.City, e.ZipCode)
}

func (e *Employee) raiseSalary(amount float64) {
	e.Salary += amount
	fmt.Printf("%s的薪资增加%.0f，新薪资: %.0f\n", e.Name, amount, e.Salary)
}

// 4. 构造函数模式
func NewStudent(name string, age int, school string) *Student {
	return &Student{
		Name:   name,
		Age:    age,
		Grade:  0.0, // 默认成绩
		School: school,
	}
}

func NewEmployee(name, position string, salary float64, address Address) *Employee {
	return &Employee{
		Name:     name,
		Position: position,
		Salary:   salary,
		Address:  address,
	}
}

// 5. 结构体比较
type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	fmt.Println("=== Go语言结构体与方法 ===")

	// 1. 创建结构体实例
	fmt.Println("\n1. 创建结构体:")

	// 方式1: 字面量
	student1 := Student{
		Name:   "小明",
		Age:    18,
		Grade:  85.5,
		School: "清华大学",
	}

	// 方式2: 指定字段名
	student2 := Student{
		Name:   "小红",
		Age:    19,
		School: "北京大学",
		// Grade会是零值0.0
	}

	// 方式3: 使用构造函数
	student3 := NewStudent("小刚", 20, "复旦大学")
	student3.Grade = 92.0

	student1.introduce()
	student2.introduce()
	student3.introduce()

	// 2. 方法调用
	fmt.Println("\n2. 方法调用:")
	student1.updateGrade(88.0)

	if student3.isExcellent() {
		fmt.Printf("%s是优秀学生！\n", student3.Name)
	}

	// 3. 嵌入字段
	fmt.Println("\n3. 嵌入字段:")
	address := Address{
		Street:  "中关村大街1号",
		City:    "北京",
		ZipCode: "100080",
	}

	employee := NewEmployee("王工程师", "软件工程师", 15000, address)
	fmt.Println(employee.getFullInfo())

	// 直接访问嵌入字段
	fmt.Printf("员工城市: %s\n", employee.City)
	employee.raiseSalary(3000)

	// 4. 结构体指针
	fmt.Println("\n4. 结构体指针:")
	studentPtr := &student1
	fmt.Printf("通过指针访问: %s, 成绩: %.1f\n",
		studentPtr.Name, studentPtr.Grade)

	// Go会自动解引用
	studentPtr.updateGrade(95.0)
	// 5. 匿名结构体
	fmt.Println("\n5. 匿名结构体:")
	car := struct {
		Brand string
		Model string
		Year  int
	}{
		Brand: "Tesla",
		Model: "Model 3",
		Year:  2023,
	}
	fmt.Printf("汽车: %s %s (%d年)\n", car.Brand, car.Model, car.Year)

	// 6. 结构体比较
	fmt.Println("\n6. 结构体比较:")
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	p3 := Point{X: 2, Y: 3}

	fmt.Printf("p1: %s\n", p1)
	fmt.Printf("p2: %s\n", p2)
	fmt.Printf("p3: %s\n", p3)
	fmt.Printf("p1 == p2: %t\n", p1 == p2)
	fmt.Printf("p1 == p3: %t\n", p1 == p3)

	// 7. 结构体标签（常用于JSON）
	fmt.Println("\n7. 结构体标签:")
	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email,omitempty"`
	}

	user := User{
		ID:       1,
		Username: "gopher",
		Email:    "",
	}
	fmt.Printf("用户: %+v\n", user)
	//json
	json, err := json.Marshal(user)
	if err != nil {
		fmt.Println("JSON序列化错误:", err)
		return
	}
	fmt.Printf("用户JSON: %s\n", json)

	// 8. 空结构体
	fmt.Println("\n8. 空结构体:")
	type Empty struct{}

	var empty Empty
	fmt.Printf("空结构体大小: %d 字节\n", unsafe.Sizeof(empty))

	// 空结构体常用于信号传递
	done := make(chan struct{})
	go func() {
		fmt.Println("goroutine完成工作")
		done <- struct{}{} // 发送信号
	}()
	<-done // 等待信号
	fmt.Println("主程序继续执行")
}
