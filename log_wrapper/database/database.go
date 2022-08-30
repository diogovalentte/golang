package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

func CreateDB(databaseFilePath *string, verbose *bool) (*LogWrapper, error) {
	// Create sqlite database file, and "logs" and "errors" tables if not exists
	logWrapper := &LogWrapper{
		DatabaseFilePath: *databaseFilePath,
		Verbose:          *verbose,
	}

	db, err := sql.Open("sqlite", *databaseFilePath)
	if err != nil {
		return logWrapper, err
	}

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS logs (
		log_tstamp timestamp,
		log_message text
	);
	`); err != nil {
		return logWrapper, err
	}
	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS errors (
		error_tstamp timestamp,
		error_message text
	);
	`); err != nil {
		return logWrapper, err
	}

	if err = db.Close(); err != nil {
		return logWrapper, err
	}

	return logWrapper, err
}

type LogWrapper struct {
	// Struct for encapsulating useful methods

	// DatabaseFilePath is the sqlite database file path
	DatabaseFilePath string

	// Verbose if true, show SQL queries for inserting Logs and Errors in SQLite database
	Verbose bool
}

type LogValues struct {
	TimeStamp string
	Message   string
}

type ErrorValues struct {
	TimeStamp string
	Message   string
}

func (LogWrapper) ProcessLog(logMessage string) *LogValues {
	tstamp := strings.Split(strings.Split(logMessage, "[")[1], "]")[0]
	message := strings.Replace(strings.Split(logMessage, "] ")[1], "'", "''", -1) // Replace log message single quotes to double quotes in order to execute SQL insert

	logValues := LogValues{
		TimeStamp: tstamp,
		Message:   message,
	}

	return &logValues
}

func (LogWrapper) ProcessError(errMessage string) *ErrorValues {
	tstamp := time.Now().Format("2006-01-02 15:04:05")
	message := strings.Replace(errMessage, "'", "''", -1) // Replace python err single quotes to double quotes in order to execute SQL insert

	errorValues := ErrorValues{
		TimeStamp: tstamp,
		Message:   message,
	}

	return &errorValues
}

func register(db *sql.DB, table string, tstamp, message *string, verbose *bool) error {
	query := fmt.Sprintf("insert into %v values ('%v', '%v')", table, *tstamp, *message)

	if *verbose {
		fmt.Println(query)
	}
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return err
}

func (lw LogWrapper) RegisterLog(logValues *LogValues) error {
	db, err := sql.Open("sqlite", lw.DatabaseFilePath)
	if err != nil {
		return err
	}
	defer db.Close()

	err = register(db, "logs", &logValues.TimeStamp, &logValues.Message, &lw.Verbose)
	if err != nil {
		return err
	}

	return err
}

func (lw LogWrapper) RegisterErr(errorValues *ErrorValues) error {
	db, err := sql.Open("sqlite", lw.DatabaseFilePath)
	if err != nil {
		return err
	}
	defer db.Close()

	err = register(db, "errors", &errorValues.TimeStamp, &errorValues.Message, &lw.Verbose)
	if err != nil {
		return err
	}

	return err
}
