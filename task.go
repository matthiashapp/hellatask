package main

import (
	"fmt"
	"time"
)

type Task struct {
	it int
}

const NW = 10

// hella task example no halting implemented for brevity

func main() {

	// make a bufferd channel of type Task with buffer of size of the worker that are available
	// so everyone can receive
	ch := make(chan Task, NW)
	for i := 0; i < NW; i++ {
		go worker(ch)
	}

	hellaTasks := getTasks()

	for _, task := range hellaTasks {
		ch <- task
	}
}

// receive task from channel and process one at a time
func worker(ch chan Task) {
	for {
		task := <-ch
		process(task)
	}
}

// simulate slow task
func process(t Task) {
	time.Sleep(time.Second * 5)
	fmt.Println(t.it)
	fmt.Println("done")
}

// return a number of tasks that need to be done
// totally arbitrary
func getTasks() []Task {
	t := new([]Task)
	for i := 0; i < 1001; i++ {
		it := Task{i}
		*t = append(*t, it)
	}
	return *t
}
