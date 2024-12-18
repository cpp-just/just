package main

import (
	"flag"
	"fmt"
	"os"

	"just/commands"
)

func main() {
	basePath := flag.String("base", "https://github.com/LucasCzerny/justpackage-", "package database url")
	flag.Parse()

	if flag.NArg() == 0 {
		printUsage()
		os.Exit(1)
	}

	command := flag.Args()[0]
	args := flag.Args()[1:]
	var commandErr error

	switch command {
	case "install":
		commandErr = commands.Install(*basePath, args)

	case "update":
		commandErr = commands.Update(args)

	case "init":
		commandErr = commands.Init(args)

	default:
		printUsage()
		os.Exit(1)
	}

	if commandErr != nil {
		printCommandError(command, commandErr)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("just [install, update] [arguments...]")
}

func printCommandError(command string, err error) {
	fmt.Printf("[Error @ %s] %v.\n", command, err)
}
