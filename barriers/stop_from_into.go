package barriers

import (
	"fmt"
	"time"
)

func StopGoroutineFromInto() {
	die := make(chan bool)
	fmt.Println("-> Created stop channel!")
	go anyWorkerFromInto(die)
	fmt.Println("-> Sending 'true' signal to channel...")
	die <- true

	time.Sleep(2 * time.Second)

	<-die
	fmt.Println("-> Stop channel has been closed!")
}

func anyWorkerFromInto(die chan bool) {
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
