package commands

import (
	"errors"
	"fmt"
	"os/exec"
)

func Configure(args []string) error {
	if len(args) > 1 {
		return errors.New("configure only zero or one argument")
	}

	var target string
	if len(args) == 1 {
		target = args[0]
	} else {
		target = "cmake"
	}

	cmd := exec.Command("cd .just; ./premake5", target, "; cd ..")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run premake (%w)", err)
	}

	return nil
}
