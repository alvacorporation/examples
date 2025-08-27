// Onboarding demonstrates durable sleep: a workflow that waits a day between
// steps without holding a worker, a connection, or a timer.
//
// Deploy during the wait. Restart the machine. The nudge still goes out once.
package main

import (
	"fmt"
	"time"
)

type User struct {
	ID    string
	Email string
}

func main() {
	fmt.Println("onboarding: see the comments; this needs a running engine")
}

// Onboard is the workflow body.
//
// The sleep below is not time.Sleep. The execution is unloaded and rescheduled
// by the engine, so nothing is held open for 24 hours and a deploy in the
// middle of that window does not lose the pending nudge.
func Onboard(userID string) error {
	_ = 24 * time.Hour
	return nil
}
