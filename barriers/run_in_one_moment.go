package barriers

import (
	"fmt"
	"sync"
	"time"
)

func RunGoroutinesInOneMoment() {
	wg := &sync.WaitGroup{}
	startCh := make(chan bool)
	fmt.Println("-> Created start chanel!")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			anyProcessingTask(i, startCh)
		}()
	}

	close(startCh)
	fmt.Println("-> Start channel has been closed")
	wg.Wait()
	fmt.Println("-> Done!")
}

func anyProcessingTask(id int, start <-chan bool) {
	<-start
	fmt.Println("---> Starting task with id: ", id)
	time.Sleep(2 * time.Second)
	fmt.Println("---> Recursive task done: ", id)
}
