package test

import (
	"testing"
)

func TestParseCaughtFiles(t *testing.T) {

	t.Run("Test: AwaitGoroutineByChannel", func(t *testing.T) {
		AwaitGoroutineByChannel()
	})

	t.Run("Test: AwaitGoroutineByChannelRecursive", func(t *testing.T) {
		AwaitGoroutineByChannelRecursive()
	})

	t.Run("Test: AwaitGoroutineByChannel", func(t *testing.T) {
		AwaitGoroutineByWaitGroup()
	})

	t.Run("Test: AwaitGoroutineByWaitGroup2", func(t *testing.T) {
		AwaitGoroutineByWaitGroup2()
	})

	t.Run("Test: ExecGoroutinesInOneMoment", func(t *testing.T) {
		ExecGoroutinesInOneMoment()
	})

	t.Run("Test: StopGoroutinesInOneMoment", func(t *testing.T) {
		StopGoroutinesInOneMoment()
	})

	t.Run("Test: StopGoroutineFromInto", func(t *testing.T) {
		StopGoroutineFromInto()
	})
}
