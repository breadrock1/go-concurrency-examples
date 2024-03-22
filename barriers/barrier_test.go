package barriers

import "testing"

func TestBarrierExamples(t *testing.T) {
	t.Run("Test: RunGoroutinesInOneMoment", func(t *testing.T) {
		RunGoroutinesInOneMoment()
	})

	t.Run("Test: StopGoroutinesInOneMoment", func(t *testing.T) {
		StopGoroutinesInOneMoment()
	})

	t.Run("Test: StopGoroutineFromInto", func(t *testing.T) {
		StopGoroutineFromInto()
	})
}
