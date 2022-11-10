package main

import (
	"fmt"
	"sync"
)

// Pool
type Pool struct {
	Tasks []*Task

	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

func (p *Pool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}
	p.wg.Add(len(p.Tasks))

	// Input tasks into task channel
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	close(p.tasksChan)
	p.wg.Wait()
}

// Task model
type Task struct {
	Err error
	f   func() error
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}
func (t *Task) Run(wg *sync.WaitGroup) {
	// Worker do task
	t.Err = t.f()
	wg.Done()
}

//MAIN

func main() {
	var tasks []*Task
	for i := 1; i <= 10; i++ {
		str := fmt.Sprintf("Task %d", i)
		tasks = append(tasks, NewTask(func() error {
			fmt.Println(str)
			return nil
		}))
	}

	pool := NewPool(tasks, 3)
	pool.Run()

	var numErrors int
	for _, task := range pool.Tasks {
		if task.Err != nil {
			fmt.Println(task.Err)
		}
		if numErrors >= 10 {
			fmt.Printf("Too many errors")
			break
		}
	}

}
