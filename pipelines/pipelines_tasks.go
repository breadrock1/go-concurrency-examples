package pipelines

import (
	"fmt"
	"time"
)

func RunPipelinesWithTasks(allTasks []func(int)) {
	majorInCh := make(chan int)

	tmpInCh := make(chan int)
	tmpOutCh := majorInCh
	for _, task := range allTasks {
		tmpInCh = tmpOutCh
		tmpOutCh = make(chan int)
		go TaskWrap(tmpInCh, tmpOutCh, task)
	}

	fmt.Println("-> Launching pipelines...")
	majorInCh <- 0
	fmt.Println("-> Returned result: ", <-tmpOutCh)
	fmt.Println("Done!")
}

func TaskWrap(in <-chan int, out chan<- int, task func(int)) {
	inData := <-in
	task(inData)
	out <- inData + 1
}

func Task(id int) {
	fmt.Println("---> Executing task: ", id)
	time.Sleep(1 * time.Millisecond)
}
