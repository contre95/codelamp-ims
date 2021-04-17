package projects

import (
	"time"
)

type Project struct {
	ID            int
	Name          string
	StartDate     time.Time
	FinishDate    time.Time
	Website       string
	GitRepository string
	Type          string
	State         string
	Tag           string
}
