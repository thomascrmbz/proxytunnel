package handler

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"unsafe"

	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"
	"thomascrmbz.com/proxytunnel"
	"thomascrmbz.com/proxytunnel/agent"
)

func ShellHandler(a *agent.Agent, s ssh.Session) {
	sshExe(a, s, sshOptions{
		CopyStdin:  true,
		CopyStdout: true,
	})
}

type sshOptions struct {
	CopyStdin  bool
	CopyStdout bool
	Cmd        []string
	OnRun      func()
	OnExit     func()
	OnDone     func()
}

func sshExe(a *agent.Agent, s ssh.Session, options sshOptions, args ...string) {
	if len(options.Cmd) == 0 {
		options.Cmd = []string{"-tt", "-oLogLevel=QUIET", a.SSH, "-p", strconv.Itoa(a.Port)}
	}
	cmd := exec.Command("ssh", append(options.Cmd, args...)...)
	ptyReq, winCh, _ := s.Pty()
	cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
	f, err := pty.Start(cmd)
	if err != nil {
		s.Exit(int(proxytunnel.PTY_FAILED))
	}
	go func() {
		for win := range winCh {
			setWinsize(f, win.Width, win.Height)
		}
	}()

	if options.CopyStdin {
		go func() {
			io.Copy(f, s)
		}()
	}
	if options.CopyStdout {
		io.Copy(s, f)
	}

	if options.OnRun != nil {
		options.OnRun()
	}

	if err := cmd.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if options.OnExit != nil {
				options.OnExit()
			}
			s.Exit(exitError.ExitCode())
		}
	}

	if options.OnDone != nil {
		options.OnDone()
	}

}

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ), uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}
