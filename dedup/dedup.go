package dedup

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type CustomEvent struct {
	Data  string
	Error error
}

func RunDedupChannelData(nw chan *CustomEvent) {
	var (
		mu      sync.Mutex
		timers  = make(map[string]*time.Timer)
		waitFor = 100 * time.Millisecond

		taskFunc = func(id string) {
			fmt.Println("---> Processing task: ", id)
			time.Sleep(5 * time.Millisecond)
		}
	)

	for {
		select {
		case event := <-nw:
			if event.Error != nil {
				fmt.Println("-> Closing processing! Done!")
				return
			}

			mu.Lock()
			t, ok := timers[event.Data]
			mu.Unlock()

			if !ok {
				t = time.AfterFunc(math.MaxInt64, func() { taskFunc(event.Data) })
				t.Stop()

				mu.Lock()
				timers[event.Data] = t
				mu.Unlock()
			}

			t.Reset(waitFor)
		}
	}
}
