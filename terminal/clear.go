package terminal

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

// Clear
// the console
func Clear() error {
	runtimeOs := runtime.GOOS

	switch runtimeOs {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		return errors.New("console cannot be cleared, platform is unsupported")
	}

	return nil
}
