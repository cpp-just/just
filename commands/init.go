package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"

	"just/common"
)

func Init(args []string) error {
	_, err := os.Stat(".just")
	if err == nil {
		return errors.New("just was already initialized")
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to check if the .just folder exists (%w)", err)
	}

	common.PrintLogo()

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get the current working directory (%w)", err)
	}
	fmt.Printf("Initializing just in %s. Press enter to continue.\n", cwd)

	if len(args) != 0 {
		fmt.Printf("Warning: ignoring args after \"init\".")
	}

	fmt.Scanln()

	metaData, err := askMetaData()
	if err != nil {
		return fmt.Errorf("failed to ask for the meta data (%w)", err)
	}

	fmt.Println()
	fmt.Println("Downloading the starter template...")

	err = downloadStarterTemplate()
	if err != nil {
		return fmt.Errorf("failed to download the starter template (%w)", err)
	}

	err = common.WriteMetaData(metaData)
	if err != nil {
		return fmt.Errorf("failed to write the meta data (%w)", err)
	}

	return nil
}

func askMetaData() (common.MetaData, error) {
	metaData := common.MetaData{}
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Whats the name of your project?")
	fmt.Print("> ")

	metaData.ProjectName, _ = scanner.ReadString('\n')
	metaData.ProjectName = strings.TrimSpace(strings.ReplaceAll(metaData.Path, "\n", ""))

	fmt.Println()

	fmt.Println("Which compiler do you want to use? [gcc/clang/... or a full path]")
	fmt.Print("> ")

	// TODO: this doesnt check in PATH lmao

	metaData.Path, _ = scanner.ReadString('\n')
	metaData.Path = strings.TrimSpace(strings.ReplaceAll(metaData.Path, "\n", ""))
	_, err := os.Stat(metaData.Path)
	if err == nil {
		return metaData, errors.New("invalid compiler path")
	} else if !errors.Is(err, os.ErrNotExist) {
		return metaData, fmt.Errorf("failed to check if the compiler path exists (%w)", err)
	}

	return metaData, nil
}

func downloadStarterTemplate() error {
	const githubStarterURL string = "https://github.com/cpp-just/just-folder-template"

	_, err := git.PlainClone(".just", false, &git.CloneOptions{
		URL:      githubStarterURL,
		Progress: os.Stdout,
	})
	if err != nil {
		return fmt.Errorf("failed to clone the starter template (%w)", err)
	}

	return nil
}
