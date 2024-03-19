package test

import (
	"fmt"
	"sync"
	"time"
)

func StopGoroutineFromInto() {
	die := make(chan bool)
	go anyWorker2(die)
	die <- true

	time.Sleep(2 * time.Second)

	<-die
}

func anyWorker2(die chan bool) {
	select {
	case <-die:
		fmt.Println("---> Stop executing task...")
		die <- true
		return
	default:
		fmt.Println("---> Executing task...")
		return
	}
}

func StopGoroutinesInOneMoment() {
	die := make(chan bool)

	for i := 0; i < 100; i++ {
		go anyWorker(die)
	}

	close(die)
}

func anyWorker(die <-chan bool) {
	select {
	case <-die:
		fmt.Println("---> Stopping before exec task...")
		return
	default:
		someProcessingTask()
	}
}

func someProcessingTask() {
	fmt.Println("---> Starting recursive task...")
	time.Sleep(2 * time.Second)
	fmt.Println("---> Recursive task done!")
}

func ExecGoroutinesInOneMoment() {
	wg := &sync.WaitGroup{}
	startCh := make(chan bool)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			anyProcessingTask(i, startCh)
		}()
	}

	close(startCh)
	wg.Wait()
}

func anyProcessingTask(id int, start <-chan bool) {
	<-start
	fmt.Println("---> Starting recursive task: ", id)
	time.Sleep(2 * time.Second)
	fmt.Println("---> Recursive task done: ", id)
}
