package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	log := log.New(os.Stderr, "", 0)

	if len(os.Args) < 2 {
		log.Fatal("rep: number of repeats unspecified")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("rep: %q is not a number", os.Args[1])
	}

	if len(os.Args) < 3 {
		log.Fatal("rep: no command specified")
	}

	if _, err := exec.LookPath(os.Args[2]); err != nil {
		log.Fatal("rep: command not found")
	}

	for i := 0; i < n; i++ {
		cmd := exec.Command(os.Args[2], os.Args[3:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal("rep: unknown problem running command")
		}
	}
}
