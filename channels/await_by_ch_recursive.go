package channels

import (
	"fmt"
	"time"
)

func AwaitGoroutineByChannelRecursive() {
	doneCh := make(chan bool)
	defer close(doneCh)

	fmt.Println("Starting goroutine...")
	go func() {
		fmt.Println("Processing task...")
		go anyProcessingAsyncTask(doneCh)
	}()

	fmt.Println("Let's wait channel...")
	<-doneCh

	fmt.Println("Done!")
}

func anyProcessingAsyncTask(doneCh chan<- bool) {
	fmt.Println("---> Starting recursive task...")
	time.Sleep(2 * time.Second)
	fmt.Println("---> Recursive task done!")

	fmt.Println("Sending result to channel")
	doneCh <- true
}
