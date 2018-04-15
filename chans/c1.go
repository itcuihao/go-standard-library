package chans

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func rangec() {
	go func() {
		time.Sleep(1 * time.Second)
	}()

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		// close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("Finished")
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func timeoutc() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 0)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}

func closec() {
	go func() {
		time.Sleep(time.Second)
	}()

	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	// c <- 3 // send on closed channel
	i, ok := <-c
	fmt.Println(i, ok)
	fmt.Println(c)
	fmt.Println(<-c)
	i, ok = <-c
	fmt.Println(i, ok)
}
