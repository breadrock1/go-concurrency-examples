package dedup

import (
	"errors"
	"testing"
	"time"
)

func TestDedupExamples(t *testing.T) {
	t.Run("Test: RunDedupChannelData", func(t *testing.T) {
		inputCh := make(chan *CustomEvent)
		go RunDedupChannelData(inputCh)

		for i := 0; i < 10; i++ {
			event := &CustomEvent{Data: "Some data", Error: nil}
			time.Sleep(5 * time.Millisecond)
			inputCh <- event
		}

		for i := 0; i < 10; i++ {
			event := &CustomEvent{Data: "Some data 2", Error: nil}
			time.Sleep(5 * time.Millisecond)
			inputCh <- event
		}

		lastEvent := &CustomEvent{Data: "", Error: errors.New("final data")}
		inputCh <- lastEvent
		time.Sleep(2 * time.Second)
	})
}
