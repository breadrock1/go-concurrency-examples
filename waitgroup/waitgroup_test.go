package waitgroup

import "testing"

func TestWaitGroupExamples(t *testing.T) {

	t.Run("Test: AwaitGoroutineByChannel", func(t *testing.T) {
		AwaitGoroutineByWaitGroup()
	})

	t.Run("Test: AwaitGoroutineByPassingWaitGroup", func(t *testing.T) {
		AwaitGoroutineByPassingWaitGroup()
	})

}
