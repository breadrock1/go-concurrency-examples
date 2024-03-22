package channels

import "testing"

func TestChannelsExamples(t *testing.T) {
	t.Run("Test: AwaitGoroutineByChannel", func(t *testing.T) {
		AwaitGoroutineByChannel()
	})

	t.Run("Test: AwaitGoroutineByChannelRecursive", func(t *testing.T) {
		AwaitGoroutineByChannelRecursive()
	})
}
