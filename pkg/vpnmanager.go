package vpnmanager

import (
	"os/exec"
)

func StartVPN() error {
	cmd := exec.Command("systemctl", "start", "wg-manager")
	return cmd.Run()
}

func StopVPN() error {
	cmd := exec.Command("systemctl", "stop", "wg-manager")
	return cmd.Run()
}

func GetStatus() (string, error) {
	cmd := exec.Command("systemctl", "is-active", "wg-manager")
	out, err := cmd.Output()
	return string(out), err
}
