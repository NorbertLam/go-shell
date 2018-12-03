package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputCommand(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Printf(scanner.Err())
	}
}

func inputCommand(input string) {
	strings.Replace(input, "\n", "", -1)
	args := strings.Fields(input)
	cmd := exec.Command(args[0], args[1:]...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Printf(stderr.String())
	}

	fmt.Printf(stdout.String())
	fmt.Printf("> ")
}
