package domain

import (
	"time"
)

func GenerateID() int64 {
	return time.Now().UnixNano()
}
