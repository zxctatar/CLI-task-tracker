package utils

import "time"

func FormattedCurrentTime() string {
	return time.Now().Format("02.01.2006 15:04")
}