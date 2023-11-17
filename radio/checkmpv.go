package radio

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckMPV() {
	check := exec.Command("mpv", "-v")
	_, err := check.Output()
	if err != nil {
		fmt.Errorf("MPV is not installed")
		os.Exit(1)
	}
}
