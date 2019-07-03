package domain

import "time"

type Payment struct {
	UserID int64
	ID     int
	Value  float32
	Date   time.Time
}
