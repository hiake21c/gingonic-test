package job

import (
	"fmt"
)

// Job Specific Functions
type Scheduler struct {
	// filtered
}

// Scheduler.Run() will get triggered automatically.
func (e Scheduler) Run() {
	fmt.Println("Every 5 sec excute Scheduler")
}
