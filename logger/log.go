package logger

import "time"

type log struct {
	timestamp time.Time
	level level
	message string 
}

func (l log) String() string {
	return "[" + l.timestamp.Format(time.RFC3339) + "] " + l.level.String() + ": " + l.message
}
