package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LucasCzerny/just/commands"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get the cwd.")
	}

	fmt.Printf("cwd is %s. Press enter to continue.\n", cwd)
	fmt.Scanln()

	flag.Parse()

	if (flag.NArg() == 0) {
		printUsage();
		os.Exit(1)
	}

	command := flag.Args()[0]

	switch (command) {
	case "install":
		commands.Install(flag.Args()[1:])

	case "update":
		// update()

	default:
		printUsage()
		os.Exit(1)
	}

	fmt.Println("Thanks for using just!")
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("just [install, update] [arguments...]")
}