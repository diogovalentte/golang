package script

import (
	"bufio"
	"log"
	"os/exec"
)

func ExecScript(fileName string, stdoutChan, stderrChan chan<- string) {
	// Execute a Python script forcing stdout and stderr streams to be unbuffered into the function channels arguments

	// Prepare command execution and stream of stdout and stderr
	cmd := exec.Command("python3", "-u", fileName) // Python script executed with "-u" arg that "force the stdout and stderr streams to be unbuffered"
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Command stout and stderr to channels
	stdoutScanner := bufio.NewScanner(stdoutPipe)
	stderrScanner := bufio.NewScanner(stderrPipe)

	go func(stdoutChan chan<- string) {
		// Continually read script stdout and send to stdout channel
		defer close(stdoutChan)
		for stdoutScanner.Scan() {
			stdoutChan <- stdoutScanner.Text()
		}
	}(stdoutChan)

	go func(stderrChan chan<- string) {
		// Continually read script stderr and buffer all messages to send in stderr channel if a error in python script occour
		defer close(stderrChan)

		var buffer string
		for stderrScanner.Scan() {
			buffer = buffer + stderrScanner.Text() + "\n"
		}
		if buffer != "" {
			buffer = buffer[:(len(buffer) - 1)] // Remove last "\n"
			stderrChan <- buffer
		}
	}(stderrChan)

	// Start command
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error while trying to execute python file: %v", err)
	}
}
