package main

import (
	// "errors"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//指针题目1
	num := 10
	increaseNum(&num)
	fmt.Println("指针的值增加10后为：", num)
	//指针题目2
	slices := []int{1, 2, 3, 4, 5}
	multiplyTwoNum(&slices)
	fmt.Println("切片中的每个元素乘以2为：", slices)
	//Goroutine题目1
	//编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
	go printOdd()
	go printEven()
	time.Sleep(1 * time.Second) // 等待一秒以确保所有输出都完成
	//Goroutine题目2
	//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
	tasks := []Task{
		exampleTask(1 * time.Second),
		exampleTask(2 * time.Second),
		exampleTask(3 * time.Second),
	}
	results, errors := goTasks(tasks)
	for i, result := range results {
		fmt.Printf("Task %d result: %v, error: %v\n", i+1, result, errors[i])
	}
	//面向对象题目1
	rect := Rectangle{Width: 3, Height: 4}
	circ := Circle{Radius: 5}
	shapes := []Shape{rect, circ}
	for _, s := range shapes {
		fmt.Printf("Area面积: %v, Perimeter周长: %v\n", s.Area(), s.Perimeter())
	}
	//面向对象题目2
	employee := Employee{
		Person: Person{
			Name: "张三",
			Age:  20,
		},
		EmployeeID: "a123",
	}
	employee.PrintInfo()
	//Channel题目1
	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
	ch1 := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	for num := range ch1 {
		fmt.Println(num)
	}
	//Channel题目2
	//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
	ch2 := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch2 <- i
			time.Sleep(time.Millisecond * 10)
		}
		close(ch2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch2 {
			fmt.Println("消费者协程从通道中接收这些整数并打印:", i)
		}
	}()
	wg.Wait() //等待所有协程完成
	//锁机制题目1
	countNumber()
	//锁机制题目2
	counterNumAtomic()
}

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
func increaseNum(ptr *int) {
	*ptr = *ptr + 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func multiplyTwoNum(x *[]int) {
	for i := range *x {
		(*x)[i] *= 2
	}
}

// 打印从1到10的奇数
func printOdd() {
	for i := 0; i < 5; i++ {
		fmt.Println("打印从1到10的奇数：", 2*i+1)
	}
}

// 打印从2到10的偶数
func printEven() {
	for i := 1; i < 6; i++ {
		fmt.Println("打印从2到10的偶数：", 2*i)
	}
}

// 定义一个任务函数
type Task func() (interface{}, error)

// 执行任务收集结果
func goTasks(tasks []Task) ([]interface{}, []error) {
	var wg sync.WaitGroup
	results := make([]interface{}, len(tasks))
	errors := make([]error, len(tasks))

	for i, task := range tasks {
		wg.Add(1)
		go func(i int, task Task) {
			defer wg.Done()
			result, err := task()
			results[i] = result
			errors[i] = err
		}(i, task)
	}
	wg.Wait() //等待所有goroutine完成

	return results, errors
}

// 示例任务函数
func exampleTask(delay time.Duration) Task {
	return func() (interface{}, error) {
		time.Sleep(delay) //模拟耗时操作
		return delay, nil
	}
}

// 定义Shape函数
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area矩形面积计算方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter矩形周长计算方法
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// Circle结构体
type Circle struct {
	Radius float64
}

// Area圆面积计算方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter圆周长计算方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 定义一个Person结构体
type Person struct {
	Name string
	Age  int
}

// 定义Employee结构体
type Employee struct {
	Person
	EmployeeID string
}

// 输出员工信息
func (e Employee) PrintInfo() {
	fmt.Println("员工姓名：", e.Name)
	fmt.Println("员工年龄：", e.Age)
	fmt.Println("员工工号：", e.EmployeeID)
}

// 定义一个结构体计数器和互斥锁
type Counter struct {
	number int
	mu     sync.Mutex
}

// 增加计数器的方法
func (c *Counter) Increment() {
	c.mu.Lock()         //在修改之前加锁
	defer c.mu.Unlock() //返回参数时释放锁
	c.number++
}

// 获取计数器的当前值
func (c *Counter) GetNumber() int {
	c.mu.Lock()         //获取之前加锁
	defer c.mu.Unlock() //返回时释放锁
	return c.number
}

// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func countNumber() {
	//初始化计数器
	counter := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("最后输出计数器的值为：", counter.GetNumber())
}

var countNum int64

func counterNumAtomic() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&countNum, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("使用原子操作，最后输出计数器的值为：", atomic.LoadInt64(&countNum))
}
