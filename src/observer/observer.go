package main

import (
	"fmt"
)

// Subject 接口，它相当于是发布者的定义
type Subject interface {
	Subscribe(observer Observer)
	Notify(msg string)
}

// Observer 观察者接口
type Observer interface {
	Update(msg string)
}

// Subject 实现
type SubjectImpl struct {
	observers []Observer
}

// Subscribe 添加观察者（订阅者）
func (sub *SubjectImpl) Subscribe(observer Observer) {
	sub.observers = append(sub.observers, observer)
}


// Notify 发布通知
func (sub *SubjectImpl) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}

// Observer1 Observer1
type Observer1 struct{}

// Update 实现观察者接口
func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}

// Observer2 Observer2
type Observer2 struct{}

// Update 实现观察者接口
func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s\n", msg)
}

func main(){
	sub := &SubjectImpl{}
	sub.Subscribe(&Observer1{})
	sub.Subscribe(&Observer2{})
	sub.Notify("Hello")
}
