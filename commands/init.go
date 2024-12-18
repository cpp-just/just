package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"just/common"
)

func Init(args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get the cwd.")
	}

	// Check if it was already initialized
	justFolderPath := fmt.Sprintf("%s/.just", cwd)
	_, err = os.Stat(justFolderPath)
	if err == nil {
		return errors.New("just was already initialized")
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to check if the .just folder exists (%w)", err)
	}

	common.PrintLogo()

	metaData := common.MetaData{}
	scanner := bufio.NewReader(os.Stdin)

	fmt.Printf("Initializing just in %s. Press enter to continue.\n", cwd)

	if len(args) != 0 {
		fmt.Printf("Warning: ignoring args after \"init\".")
	}

	fmt.Scanln()

	fmt.Println("Whats the name of your project?")
	fmt.Print("> ")

	metaData.ProjectName, _ = scanner.ReadString('\n')
	metaData.ProjectName = strings.TrimSpace(metaData.ProjectName)

	fmt.Println()

	fmt.Println("Which compiler do you want to use? [gcc/clang/... or a full path]")
	fmt.Print("> ")

	// TODO: this doesnt check in PATH lmao

	metaData.Path, _ = scanner.ReadString('\n')
	metaData.Path = strings.TrimSpace(metaData.Path)
	_, err = os.Stat(metaData.Path)
	if err == nil {
		return errors.New("invalid compiler path")
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to check if the compiler path exists (%w)", err)
	}

	err = os.Mkdir(".just", os.ModePerm)
	if err != nil {
		return errors.New("failed to create the .just folder")
	}
	err = common.WriteMetaData(metaData)
	if err != nil {
		return fmt.Errorf("failed to write the metadata (%w)", err)
	}

	fmt.Println()

	// TODO: cwd?

	_, err = os.Create(".just/buildoptions")
	if err != nil {
		return errors.New("failed to create the buildoptions file")
	}

	_, err = os.Create(".just/linkeroptions")
	if err != nil {
		return errors.New("failed to create the linkeroptions file")
	}

	_, err = os.Create(".just/defines")
	if err != nil {
		return errors.New("failed to create the defines file")
	}

	gitignoreContent := ".just/\nbuild/"
	err = os.WriteFile(".gitignore", []byte(gitignoreContent), os.ModePerm)
	if err != nil {
		return errors.New("failed to create the .gitignore file")
	}

	return nil
}
