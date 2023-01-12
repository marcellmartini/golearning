package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

	"basic/tools"
)

type person struct {
	First string
	Last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("Meu nome é:", p.First, p.Last, "e tenho", p.age, "anos")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}

func exercise1() {
	tools.HeaderOutPut("Goroutines")

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("Arch\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("goRoutine\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("goRoutine\t", runtime.NumGoroutine())
	wg.Wait()
}

func exercise2() {
	tools.HeaderOutPut("exercise2")

	p := person{
		First: "Marcell",
		Last:  "Martini",
		age:   42,
	}

	saySomething(&p)
}

// exercise2 with race conditions
func exercise3() {
	tools.HeaderOutPut("exercise3")

	var wg3 sync.WaitGroup

	incrementer := 0
	gs := 100
	wg3.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg3.Done()
		}()
	}
	wg3.Wait()
	fmt.Println("end value:", incrementer)
}

// exercise4 without race conditions with mutex
func exercise4() {
	tools.HeaderOutPut("exercise4")

	var wg4 sync.WaitGroup
	var mux sync.Mutex

	incrementer := 0
	gs := 100
	wg4.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mux.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			mux.Unlock()
			wg4.Done()
		}()
	}
	wg4.Wait()
	fmt.Println("end value:", incrementer)
}

// exercise4 without race conditions with atomic
func exercise5() {
	tools.HeaderOutPut("exercise5")

	var wg5 sync.WaitGroup
	var incrementer int64

	gs := 100
	wg5.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg5.Done()
		}()
	}
	wg5.Wait()
	fmt.Println("end value:", incrementer)
}

func exercise6() {
	tools.HeaderOutPut("exercise6")

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}

// Examples of concurrency
func Examples() {
	tools.HeaderOutPut("concurrencyExample")

	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}
