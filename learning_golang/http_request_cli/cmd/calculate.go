/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Make a calculation",
	Long:  `Make a HTTP request to a local server calculate some math operations and return the result`,

	Run: func(cmd *cobra.Command, args []string) {
		result := doCalculation(args)
		fmt.Printf("%v", result)
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calculateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// JSON response struct
type JSONCalculateResponse struct {
	Num1     int
	Num2     int
	Operator string
	Result   int
	Error    string
}

func doCalculation(args []string) string {
	// Command args
	if len(args) != 3 {
		fmt.Println("Should provide 3 aguments")
		os.Exit(1)
	}
	num1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error loading num1: %v", err)
	}
	num2, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error loading num2: %v", err)
	}
	operator := args[2]

	// HTTP request
	hostAddr := "127.0.0.1"
	port := "3333"
	requestURL := fmt.Sprintf("http://%v:%v/calculate?num1=%v&num2=%v", hostAddr, port, num1, num2)

	body := []byte(operator)
	bodyReader := bytes.NewReader(body)

	req, _ := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error while doing HTTP request: %v", err)
		os.Exit(1)
	}
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	// Process HTTP response
	var JSONResult JSONCalculateResponse
	err = json.Unmarshal(responseBody, &JSONResult)
	if err != nil {
		fmt.Printf("Error while reading HTTP response JSON: %v", err)
		os.Exit(1)
	}

	// Build result
	var result string

	if JSONResult.Error == "" {
		result = fmt.Sprintf("%v %v %v = %v\n", JSONResult.Num1, JSONResult.Operator, JSONResult.Num2, JSONResult.Result)
	} else {
		fmt.Printf("Error in server: %v", JSONResult.Error)
		os.Exit(1)
	}

	return result
}
