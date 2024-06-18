package shell

import (
	"os"
	"os/exec"
)

func Init(proxy string) error {
	os.Setenv("http_proxy", proxy)

	cmd := exec.Command(defaultShell())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func defaultShell() string {
	return os.Getenv("SHELL")
}
