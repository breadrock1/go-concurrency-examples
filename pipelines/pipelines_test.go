package pipelines

import "testing"

func TestPipelinesExamples(t *testing.T) {
	t.Run("Test: RunSimplePipelines", func(t *testing.T) {
		RunSimplePipelines()
	})

	t.Run("Test: RunPipelinesWithTasks", func(t *testing.T) {
		var tasks []func(int)
		for i := 0; i < 10000; i++ {
			tasks = append(tasks, Task)
		}

		RunPipelinesWithTasks(tasks)
	})
}
