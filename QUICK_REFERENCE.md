# Go语法快速参考卡片

## 🚀 快速开始
```bash
# 运行单个示例
go run 01_basics/hello.go

# 运行所有示例
./run_all.sh

# 运行指定章节
./run_chapter.sh 5
```

## 📝 核心语法速查

### 变量声明
```go
var name string = "Go"     // 完整声明
var age = 25              // 类型推断
count := 10               // 短声明（函数内）
const pi = 3.14           // 常量
```

### 基本类型
```go
bool, string
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
```

### 控制结构
```go
// if语句
if age >= 18 {
    fmt.Println("成年")
}

// switch语句
switch day {
case "Monday":
    fmt.Println("周一")
default:
    fmt.Println("其他")
}

// for循环
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// range循环
for index, value := range slice {
    fmt.Printf("%d: %v\n", index, value)
}
```

### 函数
```go
// 基本函数
func add(a, b int) int {
    return a + b
}

// 多返回值
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除零错误")
    }
    return a / b, nil
}

// 可变参数
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

### 结构体和方法
```go
// 定义结构体
type Person struct {
    Name string
    Age  int
}

// 值接收者方法
func (p Person) Greet() string {
    return "Hello, " + p.Name
}

// 指针接收者方法
func (p *Person) Birthday() {
    p.Age++
}

// 创建实例
person := Person{Name: "Alice", Age: 25}
person.Birthday()
```

### 接口
```go
// 定义接口
type Speaker interface {
    Speak() string
}

// 实现接口（隐式）
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return d.Name + " says: Woof!"
}

// 使用接口
var speaker Speaker = Dog{Name: "Buddy"}
fmt.Println(speaker.Speak())
```

### 并发编程
```go
// 启动goroutine
go func() {
    fmt.Println("Hello from goroutine")
}()

// 使用channel
ch := make(chan string)
go func() {
    ch <- "Hello"
}()
message := <-ch

// 使用select
select {
case msg := <-ch1:
    fmt.Println("From ch1:", msg)
case msg := <-ch2:
    fmt.Println("From ch2:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
}
```

### 错误处理
```go
// 创建错误
err := errors.New("something went wrong")
err := fmt.Errorf("failed to process %s", filename)

// 处理错误
result, err := someFunction()
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

// 自定义错误类型
type MyError struct {
    Message string
}

func (e MyError) Error() string {
    return e.Message
}
```

### 包和导入
```go
// 标准导入
import "fmt"
import "strings"

// 分组导入
import (
    "fmt"
    "strings"
    "time"
)

// 别名导入
import str "strings"

// 点导入
import . "fmt"
```

## 🎯 学习路径

### 基础阶段（1-2小时）
1. **基础语法**：变量、常量、基本类型
2. **控制流**：if、switch、for循环
3. **函数**：定义、调用、参数、返回值

### 进阶阶段（1小时）
4. **结构体**：定义、方法、组合
5. **接口**：定义、实现、多态

### 高级阶段（1-2小时）
6. **并发**：goroutines、channels、select
7. **错误处理**：错误类型、最佳实践
8. **包管理**：模块化、导入

### 实战阶段（1小时）
9. **综合练习**：完整项目示例

## 💡 重要概念

### Go语言特色
- **简洁性**：语法简单，关键字少
- **并发性**：内置goroutines和channels
- **类型安全**：静态类型，编译时检查
- **垃圾回收**：自动内存管理
- **快速编译**：编译速度极快

### 最佳实践
- 使用`gofmt`格式化代码
- 错误要显式处理
- 接口要小而精
- 组合优于继承
- 通过通信来共享内存

## 🔧 常用命令
```bash
go run main.go        # 运行程序
go build             # 编译程序
go test              # 运行测试
go fmt ./...         # 格式化代码
go vet ./...         # 静态分析
go mod init myapp    # 初始化模块
go mod tidy          # 整理依赖
```

祝你学习愉快！🎉
