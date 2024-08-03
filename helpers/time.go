package helpers

import (
	"time"
)

func GetTimestamp() int {
	return int(time.Now().Unix())
}
