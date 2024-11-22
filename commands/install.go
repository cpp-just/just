package commands

import (
	"fmt"
	// "os/exec"
)

func Install(packages ...string) {
	const basePath = "https://github.com/LucasCzerny/justpackage-"

	for _, pkg := range packages {
		packagePath := basePath + pkg

		command := fmt.Sprintf("git submodule add %s packages/%s", packagePath, pkg)
		fmt.Printf("Executing \"%s\"\n", command)
		// exec.Command(command)
	}
}