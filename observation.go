package main

import (
	"time"
)

// Observation interface will be used by
type Observation struct {
	objectID int64
	time     time.Time
	score    int
}
