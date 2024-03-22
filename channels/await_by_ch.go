package channels

import (
	"fmt"
	"time"
)

func AwaitGoroutineByChannel() {
	doneCh := make(chan bool)
	defer close(doneCh)

	fmt.Println("Starting goroutine...")
	go func() {
		fmt.Println("---> Starting recursive task...")
		time.Sleep(2 * time.Second)
		fmt.Println("---> Recursive task done!")

		fmt.Println("Sending result to channel")
		doneCh <- true
	}()

	fmt.Println("Let's wait channel...")
	<-doneCh

	fmt.Println("Done!")
}
