package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// 类型 T 只有接受者是 T 的方法；而类型 *T 拥有接受者是 T 和 *T 的方法。语法上 T 能直接调 *T 的方法仅仅是 Go 的语法糖。
func (s *Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}

func (s *Student) growup(){
	s.Age++
}

func main() {
	var s = Student{
		Name: "name",
		Age:  17,
	}
	s.growup()
	fmt.Println(s)  // {name 18}
	fmt.Println(&s) // [Name: name], [Age: 18]
}
