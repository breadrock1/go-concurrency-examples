package pipelines

import (
	"fmt"
	"time"
)

func RunSimplePipelines() {
	majorInCh := make(chan int)

	tmpInCh := make(chan int)
	tmpOutCh := majorInCh
	for i := 0; i < 10000; i++ {
		tmpInCh = tmpOutCh
		tmpOutCh = make(chan int)
		go anyProcessingTask(tmpInCh, tmpOutCh)
	}

	fmt.Println("-> Launching pipelines...")
	majorInCh <- 0
	fmt.Println("-> Returned result: ", <-tmpOutCh)
	fmt.Println("Done!")
}

func anyProcessingTask(in <-chan int, out chan<- int) {
	inData := <-in
	fmt.Println("---> Processing task: ", inData)
	time.Sleep(1 * time.Millisecond)
	out <- inData + 1
}
