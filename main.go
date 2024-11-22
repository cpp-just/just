package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if (flag.NArg() == 1) {
		printUsage();
		os.Exit(1)
	}

	command := flag.Args()[1]

	switch (command) {
	case "install": {
		install(flag.Args()[1:])
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