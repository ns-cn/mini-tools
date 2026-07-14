package status

type Status int

const (
	INIT Status = 1 << iota
	RUNNING
	PAUSED
	TERMINATED
)
