//go:build windows
// +build windows

package clear

func CallClear() {
	var cmd *exec.Cmd

	cmd = exec.Command("cmd", "/c", "cls")

	cmd.Stdout = os.Stdout
	cmd.Run()
}
