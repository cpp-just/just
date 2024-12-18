package commands

import (
	"errors"
	"fmt"
	"os/exec"
)

func Update(args []string) error {
	submoduleCommand := "git submodule update --init --recursive"
	cmd := exec.Command(submoduleCommand)

	err := cmd.Run()
	if err != nil {
		errorMessage := fmt.Sprintf("failed to update the packages. Reason: %s", err.Error())
		return errors.New(errorMessage)
	}

	return errors.New("update is not implemented yet")
}