package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Name string
	Age  int
}

// 类型 T 只有接受者是 T 的方法；而类型 *T 拥有接受者是 T 和 *T 的方法。语法上 T 能直接调 *T 的方法仅仅是 Go 的语法糖。
func (s *Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}

/*
编译报错
.\main.go:18:6: method redeclared: Student.String
	method(*Student) func() string
	method(Student) func() string

func (s Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}

*/

func (s *Student) growup() {
	s.Age++
}

// 结构体响应不同的接收者（值类型和指针类型）
// interface响应不同的接收者（值类型和指针类型）
func StructReceiver() {
	var s = Student{
		Name: "name",
		Age:  17,
	}
	s.growup()
	fmt.Println(s)  // {name 18}
	fmt.Println(&s) // [Name: name], [Age: 18]
	/*
		只有 *Student 实现了String()方法
	*/
}

func main() {
	StructReceiver()
	InterfaceSlice()
}

// []interface{} 的data字长是[]T 的data字长的2倍
func InterfaceSlice() {
	intSl := []int{1, 2, 3, 4}
	interfaceSl := make([]interface{}, len(intSl))

	fmt.Println("sizeof(intSl) = ", unsafe.Sizeof(intSl))             // sizeof(intSl) =  24
	fmt.Println("sizeof(interfaceSl) = ", unsafe.Sizeof(interfaceSl)) // sizeof(interfaceSl) =  24
	intDataSize := 0
	interfaceDataSize := 0
	for _, v := range interfaceSl {
		interfaceDataSize += int(unsafe.Sizeof(v))
	}
	for _, v := range intSl {
		intDataSize += int(unsafe.Sizeof(v))
	}
	fmt.Println("sizeof(interfaceSl.data) = ", interfaceDataSize) // sizeof(interfaceSl.data) =  64
	fmt.Println("sizeof(intSl.data) = ", intDataSize)             // sizeof(intSl.data) =  32
}
