package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	log := log.New(os.Stderr, "", 0)

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("rep: not enough arguments provided")
	}

	var bounded bool
	n, err := strconv.Atoi(args[0])
	if err == nil {
		bounded = true
		args = args[1:]
	}

	if len(args) < 1 {
		log.Fatal("rep: no command specified")
	}

	if _, err := exec.LookPath(args[0]); err != nil {
		log.Fatal("rep: command not found")
	}

	if !bounded {
		log.Println("rep: repeating forever")
		time.Sleep(3 * time.Second)
	}

	for i := 0; !bounded || i < n; i++ {
		cmd := exec.Command(args[0], args[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal("rep: unknown problem running command")
		}
	}
}
