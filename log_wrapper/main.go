package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/diogovalentte/golang/log_wrapper/database"
	"github.com/diogovalentte/golang/log_wrapper/script"

	"github.com/enriquebris/goconcurrentqueue"
)

var wg = sync.WaitGroup{}

func main() {
	// Flags
	actualDirName, _ := os.Getwd()
	dirName := flag.String("database-folder", actualDirName, "Folder to create SQLite database file")
	databaseFileName := flag.String("database-name", "temp_log.db", "SQLite database file name")
	pythonScriptPath := flag.String("python-script", "", "Absolute path to python script")

	flag.Parse()

	// Database set up
	databaseFilePath := filepath.Join(*dirName, *databaseFileName)
	lw, err := database.CreateDB(databaseFilePath)
	if err != nil {
		log.Fatalf("Fatal Error while creating Database: %v", err)
	}

	// Channels for stdout and stderr
	stdoutChan := make(chan string)
	stderrChan := make(chan string)
	// Queue of stdout logs to process and register in database
	stdoutQueue := goconcurrentqueue.NewFIFO()

	// Execute Python Script
	script.ExecScript(*pythonScriptPath, stdoutChan, stderrChan)

	// Execute goroutine that process stdout logs from queue
	go func(stdoutQueue *goconcurrentqueue.FIFO, lw *database.LogWrapper) {
		// Get stdout from Queue, process and register logs in database
		for {
			// This loop will be terminated with when main goroutine finish execution
			stdoutLog, _ := stdoutQueue.DequeueOrWaitForNextElement()
			outValues := lw.ProcessLog(stdoutLog.(string))
			if err := lw.RegisterLog(outValues); err != nil {
				log.Fatalf("Fatal Error while registering log: %v", err)
			}
			wg.Done()
		}
	}(stdoutQueue, lw)

	// Continually add stdout logs from python script to queue
	for stdoutLog := range stdoutChan {
		stdoutQueue.Enqueue(stdoutLog)
		wg.Add(1)
	}

	// Wait all queue stdout logs be processed before continue execution
	wg.Wait()

	// Process and register in database buffered error messages if exits
	errMessage := <-stderrChan
	if errMessage != "" {
		errValues := lw.ProcessError(errMessage)
		if err := lw.RegisterErr(errValues); err != nil {
			log.Fatalf("Falta Error while registering error: %v", err)
		}
	}
}
