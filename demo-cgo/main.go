package main

// 头文件的位置，相对于源文件是当前目录，所以是 .，头文件在多个目录时写多个  #cgo CFLAGS: ...
// 从哪里加载动态库，位置与文件名，-ladd 加载 libadd.so 文件

//#cgo CFLAGS: -I./include
//#cgo LDFLAGS: -L./lib -ladd-x86_64 -Wl,-rpath,lib
//#include "add.h"
import "C"

// import "C" 要独占一行

//
//通过注释代码来告诉 Go 编译器从哪里引入头文件与加载动态库. 本例中 *.h 和 *.go 文件在同一个目录的情况下， #cgo CFLAGS: -I. 可不写。
//CFLAGS: -I 和 LDFLAGS: -L 都是相对于源文件 main.go 的位置

import "fmt"

func main() {
	fmt.Println("demo-cgo start main")
	val := C.Add(C.CString("go"), 2021)
	fmt.Println("run c: ", C.GoString(val))
}
