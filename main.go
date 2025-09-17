package main

import (
	"fmt"
	"log"
	"os/exec"
)

func detectApp(name string) {
	cmd := exec.Command("ps", "-a")
	namePattern := fmt.Sprintf("*%s*", name)
	grepCmd := exec.Command("grep", namePattern)

	// Pipe ps output to grep input
	psOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Could not get ps output pipe: %v", err)
	}
	grepCmd.Stdin = psOut

	// Start ps command
	if err := cmd.Start(); err != nil {
		log.Fatalf("Could not start ps command: %v", err)
	}

	// Start grep command
	grepOut, err := grepCmd.Output()
	if err != nil {
		log.Fatalf("Could not find %s running", name)
	}

	fmt.Printf("%s", grepOut)
}

func main() {
	detectApp("cursor")
}
