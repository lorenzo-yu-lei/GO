package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("————————————goroutine和并发(concurrent)————————分段—————")
	go sleepyGopher(0)          //分支线路
	time.Sleep(4 * time.Second) //主线路

	fmt.Println("————————————学习goroutine—————————分段—————")

	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)

	fmt.Println("————————————学习channel—————————分段—————")
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go c_sleepyGopher(i, c)
	}
	timeout := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c:
			fmt.Println("gopher", gopherId, "has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}

	}
}
func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)
}

func c_sleepyGopher(id int, c chan int) {
	fmt.Println("...", id, "snore...")
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)

	c <- id
}
