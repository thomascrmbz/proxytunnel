package proxytunnel

type ExitCode int

var (
	OK              ExitCode = 0
	PTY_FAILED      ExitCode = 102
	AGENT_NOT_FOUND ExitCode = 103
)
