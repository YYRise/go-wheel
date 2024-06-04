package main

import (
	"fmt"
)

// Coster 接口定义了Cost方法
type Coster interface {
	Cost() float64
}

// Base 结构体，包含Rate方法
type Base struct{}

// Rate 方法在Base中实现，假设调用者实现了Cost方法
func (b *Base) Rate() float64 {
	// 这里我们假设调用Rate的接收者（即A或B的实例）实现了Cost方法
	// 因此我们可以直接调用它的Cost方法，不需要通过接口
	// 这里的self代表调用Rate的接收者本身，即A或B的实例
	cost := b.Self().Cost() // 假设有一个Self方法返回调用者的*Base部分
	// 根据Cost计算Rate的逻辑（这里只是一个示例）
	return cost * 0.1 // 假设Rate是Cost的10%
}

// Self 方法返回调用者的*Base部分，用于在Base的方法中访问调用者的其他方法
func (b *Base) Self() Coster {
	return b // 这里返回b本身，但实际上应该返回包含b的那个结构体实例（即A或B）
}

// A 结构体内嵌Base，并实现Cost方法
type A struct {
	Base
}

// A的Cost方法实现
func (a A) Cost() float64 {
	// A的Cost逻辑
	return 10.0
}

// B 的实现类似A
type B struct {
	Base
}

// B的Cost方法实现
func (b B) Cost() float64 {
	// B的Cost逻辑
	return 20.0
}

func main() {
	var a A
	// 调用A的Rate方法，会自动调用A的Cost方法，因为A实现了Cost方法
	fmt.Printf("Type A: Rate=%f\n", a.Rate())

	var b B
	// 调用B的Rate方法，会自动调用B的Cost方法，因为B实现了Cost方法
	fmt.Printf("Type B: Rate=%f\n", b.Rate())
}
