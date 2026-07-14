package signal

type Control int

const (
	QUIT Control = 1 << iota
	PAUSE
	CONTINUE
	RESUME
)
