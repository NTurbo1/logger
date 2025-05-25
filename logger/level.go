package logger

type level int

const (
	TRACE level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func (l level) String() string {
	return [...]string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[l]
}
