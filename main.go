package main

import (
	"github.com/stvp/rollbar"
)

func main() {
	rollbar.Token = "d9b441c1b032459eaa34ebbcfe7ed232"
	rollbar.Environment = "production" // defaults to "development"

	// Message reporting
	rollbar.Message("info", "Message body goes here")

	// Block until all queued messages are sent to Rollbar.
	// You can do this in a defer() if needed.
	rollbar.Wait()
}
