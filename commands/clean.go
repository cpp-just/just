package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Clean(args []string) error {
	if len(args) != 0 {
		fmt.Printf("Warning: ignoring args after \"init\".")
	}

	remove := []string{"build", "bin", ".just/CMakeLists.txt"}

	matches, err := filepath.Glob("*.cmake")
	if err != nil {
		return fmt.Errorf("failed to glob *.cmake files (%w)", err)
	}

	remove = append(remove, matches...)

	for _, path := range remove {
		err := os.RemoveAll(path)
		if err != nil {
			return fmt.Errorf("failed to remove %s (%w)", path, err)
		}
	}

	return nil
}
