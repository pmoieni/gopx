package shell

import (
	"os"
	"os/exec"
)

func Init(proxy string) error {
	if err := os.Setenv("http_proxy", proxy); err != nil {
		return err
	}

	cmd := exec.Command(defaultShell())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func defaultShell() string {
	return os.Getenv("SHELL")
}
