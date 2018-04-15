package chans

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	t.Log(x, y, x+y)

	fmt.Println(x, y, x+y)
	fmt.Println("success")
}

func TestRangec(t *testing.T) {
	rangec()
}

func TestFibonacci(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

func TestTimeout(t *testing.T) {
	timeoutc()
}

func TestClosec(t *testing.T) {
	closec()
}
