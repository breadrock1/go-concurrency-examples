package waitgroup

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

	fmt.Println("-> Waiting all tasks done")
	wg.Wait()
	fmt.Println("-> Done!")
}

func anyProcessingTaskWithoutCh(id int) {
	fmt.Println("---> Starting recursive task: ", id)
	time.Sleep(2 * time.Second)
	fmt.Println("---> Recursive task done!")
}
