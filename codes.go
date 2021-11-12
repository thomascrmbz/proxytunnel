package proxytunnel

type ExitCode int

var (
	OK                ExitCode = 0
	COMMAND_NOT_FOUND ExitCode = 101
	PTY_FAILED        ExitCode = 102
	AGENT_NOT_FOUND   ExitCode = 103
	NOT_ALLOWED       ExitCode = 104
)
