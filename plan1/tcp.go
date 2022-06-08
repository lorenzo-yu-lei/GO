package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {
	//使用goroutine 并发扫描tcp

	//start := time.Now()
	//var wg sync.WaitGroup
	//for i := 21; i < 120; i++ {
	//	wg.Add(1)
	//	go func(j int) {
	//		defer wg.Done()
	//		address := fmt.Sprintf("192.168.1.1:%d", j)
	//		conn, err := net.Dial("tcp", address)
	//		if err != nil {
	//			fmt.Printf("%s 关闭了\n", address)
	//			return
	//		}
	//		conn.Close()
	//		fmt.Printf("%s 打开了！！\n", address)
	//	}(i)
	//}
	//wg.Wait()
	//elapsed := time.Since(start) / 90
	//fmt.Printf("\n\n%d seconds", elapsed)

	//	grount 线程池main启动
	play()
}

//使用goroutine池
func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("192.168.1.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- p
			continue
		}
		conn.Close()
		results <- p
	}
}

func play() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var closedports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	
	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)

		} else {
			closedports = append(closedports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)
	sort.Ints(closedports)

	for _, port := range closedports {
		fmt.Printf("%d closed \n", port)
	}
	for _, port := range openports {
		fmt.Printf("%d open \n", port)
	}
}
