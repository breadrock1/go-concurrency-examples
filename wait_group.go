package test

import (
	"fmt"
	"sync"
	"time"
)

func AwaitGoroutineByWaitGroup() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			anyProcessingTaskWithoutCh(i)
		}()
	}

	wg.Wait()
}

func anyProcessingTaskWithoutCh(id int) {
	fmt.Println("---> Starting recursive task: ", id)
	time.Sleep(2 * time.Second)
	fmt.Println("---> Recursive task done!")
}

func AwaitGoroutineByWaitGroup2() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go anyProcessingTaskWithWaitGroup(i, wg)
	}

	wg.Wait()
}

func anyProcessingTaskWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("---> Starting recursive task: ", id)
	time.Sleep(2 * time.Second)
	fmt.Println("---> Recursive task done: ", id)
}
