package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LucasCzerny/just/commands"
)

func main() {
	flag.Parse()

	if (flag.NArg() == 0) {
		printUsage();
		os.Exit(1)
	}

	command := flag.Args()[0]

	switch (command) {
	case "install": {
		commands.Install()
	}
	case "update": {
		// update()
	}
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("just [install, update] [arguments...]")
}