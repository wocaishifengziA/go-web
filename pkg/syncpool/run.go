package syncpool

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var numCalcsCreated int32

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createBuffer() interface{} {
	// 这里要注意下，非常重要的一点。这里必须使用原子加，不然有并发问题；
	atomic.AddInt32(&numCalcsCreated, 1)
	student := Student{Name: "li", Age: 10}
	return &student
}

func Run() {
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	// 多 goroutine 并发测试
	numWorkers := 3
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		// go func() {
		// 	defer wg.Done()
		// 	// 申请一个 buffer 实例
		// 	buffer := bufferPool.Get()
		// 	fmt.Println("get buff: ", buffer)
		// 	a := buffer.(*Student)
		// 	a.Age = 11
		// 	fmt.Println("a buff: ", a)
		// 	// 释放一个 buffer 实例
		// 	defer bufferPool.Put(a)
		// }()
		// 申请一个 buffer 实例
		buffer := bufferPool.Get()
		fmt.Println("get buff: ", buffer)
		a := buffer.(*Student)
		a.Age = 11
		fmt.Println("a buff: ", a)
		// 释放一个 buffer 实例
		bufferPool.Put(buffer)
		wg.Done()
	}
	wg.Wait()
	fmt.Printf("%d buffer objects were created.\n", numCalcsCreated)
}
