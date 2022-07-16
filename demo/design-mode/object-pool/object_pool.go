package objectpool

import (
	"fmt"
	"sync"
)

type Task struct {
	id int
	f  func() error
}

func (t *Task) Exec() error {
	return t.f()
}

func NewTask(id int, f func() error) *Task {
	return &Task{
		id: id,
		f:  f,
	}
}

type WorkPool struct {
	EntryChannel chan *Task
	ReadyChannel chan *Task
	maxNum       int
	wg           *sync.WaitGroup
}

func NewWorkPool(cap int, wg *sync.WaitGroup) *WorkPool {
	return &WorkPool{
		EntryChannel: make(chan *Task),
		ReadyChannel: make(chan *Task),
		maxNum:       cap,
		wg:           wg,
	}
}

func (p *WorkPool) worker(workId int) {
	for task := range p.ReadyChannel{
		err := task.Exec()
		if err!=nil{
			fmt.Println("task")
		}
	}
}
