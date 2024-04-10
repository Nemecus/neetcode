//go:build darwin || linux
// +build darwin linux

package clear

import (
	"os"
	"os/exec"
)

func CallClear() {
	var cmd *exec.Cmd

	cmd = exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}
