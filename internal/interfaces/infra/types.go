package infra

import "time"

type InputLink struct {
	Link      string
	FakeLink  string
	EraseTime time.Time
}
