package barriers

import (
	"fmt"
	"time"
)

func StopGoroutinesInOneMoment() {
	die := make(chan bool)
	fmt.Println("-> Created stop channel!")

	for i := 0; i < 100; i++ {
		go anyWorkerInOneMoment(die)
	}

	close(die)
	fmt.Println("-> Start channel has been closed")
}

func anyWorkerInOneMoment(die <-chan bool) {
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
