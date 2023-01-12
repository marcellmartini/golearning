package channel

import (
	"fmt"
	"sync"

	"basic/tools"
)

// example1 unblock c chan with func literal
func example1() {
	c := make(chan int)

	// one goroutine block c chan
	go func() {
		c <- 41
	}()

	// main goroutine unblock c chan
	fmt.Println(<-c)
	fmt.Printf("------\n")
}

// example1 unblock c chan buffer
func example11() {
	c := make(chan int, 1)

	c <- 41
	fmt.Println(<-c)
	fmt.Printf("------\n")
}

func example2() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func example3() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
	fmt.Printf("------\n")
}

func receive(r <-chan int) {
	for v := range r {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func example4() {

	q := make(chan int)
	c := gen2(q)

	receive2(c, q)

	fmt.Println("about to exit")
	fmt.Printf("------\n")
}

func receive2(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 12; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func example5() {
	c := make(chan int)

	go func() {
		c <- 41
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
	fmt.Printf("------\n")
}

func example6() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Printf("------\n")
}

func example7() {
	var wg sync.WaitGroup
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				for j := 0; j < 10; j++ {
					c <- (i * 10) + j
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Printf("------\n")
}

// Examples calls to examples of channel
func Examples() {
	tools.HeaderOutPut("channelExample")

	example1()
	example11()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
}
