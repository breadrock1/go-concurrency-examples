package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func AwaitGoroutineByPassingWaitGroup() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go anyProcessingTaskWithWaitGroup(i, wg)
	}

	fmt.Println("-> Waiting all tasks done")
	wg.Wait()
	fmt.Println("-> Done!")
}

func anyProcessingTaskWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("---> Starting recursive task: ", id)
	time.Sleep(2 * time.Second)
	fmt.Println("---> Recursive task done: ", id)
}
