// Go语言并发编程
package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 简单的goroutine示例
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello %s! (%d)\n", name, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

// 2. 使用channel通信的goroutine
func producer(ch chan<- int, name string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s 生产了 %d\n", name, i)
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, name string) {
	for value := range ch {
		fmt.Printf("%s 消费了 %d\n", name, value)
		time.Sleep(150 * time.Millisecond)
	}
}

// 3. 工作池模式
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d 开始处理 job %d\n", id, job)
		time.Sleep(500 * time.Millisecond) // 模拟工作
		results <- job * 2 // 返回结果
		fmt.Printf("Worker %d 完成了 job %d\n", id, job)
	}
}

// 4. 使用sync.WaitGroup
func taskWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 任务完成时调用
	fmt.Printf("任务 %d 开始\n", id)
	time.Sleep(time.Duration(id*100) * time.Millisecond)
	fmt.Printf("任务 %d 完成\n", id)
}

// 5. 使用sync.Mutex保护共享资源
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func incrementCounter(counter *Counter, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		counter.Increment()
	}
}

// 6. select语句示例
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("斐波那契生成器退出")
			return
		case <-time.After(2 * time.Second):
			fmt.Println("超时，斐波那契生成器退出")
			return
		}
	}
}

func main() {
	fmt.Println("=== Go语言并发编程 ===")
	
	// 1. 基本goroutine
	fmt.Println("\n1. 基本goroutine:")
	go sayHello("Alice")
	go sayHello("Bob")
	time.Sleep(500 * time.Millisecond) // 等待goroutines完成
	
	// 2. channel基础
	fmt.Println("\n2. Channel基础:")
	
	// 无缓冲channel
	ch := make(chan string)
	go func() {
		ch <- "Hello from goroutine!"
	}()
	message := <-ch
	fmt.Println("收到消息:", message)
	
	// 有缓冲channel
	bufferedCh := make(chan int, 3)
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	
	fmt.Printf("从缓冲channel读取: %d\n", <-bufferedCh)
	fmt.Printf("从缓冲channel读取: %d\n", <-bufferedCh)
	fmt.Printf("从缓冲channel读取: %d\n", <-bufferedCh)
	
	// 3. 生产者-消费者模式
	fmt.Println("\n3. 生产者-消费者模式:")
	productCh := make(chan int)
	
	go producer(productCh, "生产者A")
	go consumer(productCh, "消费者1")
	
	time.Sleep(2 * time.Second)
	
	// 4. 工作池模式
	fmt.Println("\n4. 工作池模式:")
	const numJobs = 5
	const numWorkers = 3
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// 启动工作者
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	// 发送工作
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// 收集结果
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("收到结果: %d\n", result)
	}
	
	// 5. 使用WaitGroup
	fmt.Println("\n5. 使用WaitGroup:")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加等待的goroutine数量
		go taskWithWaitGroup(i, &wg)
	}
	
	wg.Wait() // 等待所有goroutine完成
	fmt.Println("所有任务完成")
	
	// 6. 使用Mutex保护共享资源
	fmt.Println("\n6. 使用Mutex保护共享资源:")
	counter := &Counter{}
	var wg2 sync.WaitGroup
	
	// 启动多个goroutine同时增加计数器
	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go incrementCounter(counter, 100, &wg2)
	}
	
	wg2.Wait()
	fmt.Printf("最终计数器值: %d\n", counter.Value())
	
	// 7. select语句
	fmt.Println("\n7. Select语句:")
	
	// 非阻塞的channel操作
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自ch1的消息"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自ch2的消息"
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("收到:", msg1)
		case msg2 := <-ch2:
			fmt.Println("收到:", msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("超时")
		}
	}
	
	// 8. 斐波那契示例
	fmt.Println("\n8. 斐波那契与select:")
	fibCh := make(chan int)
	quit := make(chan int)
	
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("斐波那契数: %d\n", <-fibCh)
		}
		quit <- 0
	}()
	
	fibonacci(fibCh, quit)
	
	// 9. Channel方向（单向channel）
	fmt.Println("\n9. Channel方向:")
	
	// 只写channel
	sendOnly := make(chan<- int)
	go func() {
		sendOnly <- 42
		close(sendOnly)
	}()
	
	// 只读channel（需要从双向channel转换）
	bothWays := make(chan int, 1)
	bothWays <- 100
	
	readOnly := (<-chan int)(bothWays)
	fmt.Printf("从只读channel读取: %d\n", <-readOnly)
	
	// 10. 优雅关闭
	fmt.Println("\n10. 优雅关闭:")
	done := make(chan bool)
	
	go func() {
		fmt.Println("工作中...")
		time.Sleep(1 * time.Second)
		fmt.Println("工作完成")
		done <- true
	}()
	
	<-done
	fmt.Println("程序结束")
}
