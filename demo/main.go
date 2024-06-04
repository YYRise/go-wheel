package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		for {
			i := <-ch1
			fmt.Println(i)
		}
	}()
	go func() {
		i := 0
		for {
			i = i + 1
			ch1 <- i
			time.Sleep(time.Second)
		}
	}()
	select {}
}
