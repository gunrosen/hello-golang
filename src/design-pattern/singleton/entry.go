package singleton

import (
	"fmt"
	"sync"
	"time"
)

type Singleton interface {
	AddOne() int
}

// name is normal case means private
type singleton struct {
	count int
}

var (
	instance *singleton
	once     sync.Once
)

func init() {
	//instance = &singleton{count: 100}
	fmt.Println("init func")
}

// GetInstance Cho phep cac package khac co the truy xuat vao lay singleton
func GetInstance() Singleton {
	// Chay nhieu goroutine co the gay loi -> tao nhieu doi tuong thay vi 1
	once.Do(func() {
		time.Sleep(time.Second)
		instance = &singleton{count: 100}
	})
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
