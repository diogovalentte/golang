/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// checkStatusCmd represents the checkStatus command
var checkStatusCmd = &cobra.Command{
	Use:   "check-status",
	Short: "Check server status",
	Long:  `Make a HTTP request to the server and get the request status code`,
	Run: func(cmd *cobra.Command, args []string) {
		checkStatus()
	},
}

func init() {
	rootCmd.AddCommand(checkStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkStatus() {
	hostAddr := "127.0.0.1"
	port := "3333"
	requestURL := fmt.Sprintf("http://%v:%v/check_status", hostAddr, port)

	req, _ := http.NewRequest(http.MethodPost, requestURL, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error while doing HTTP request: %v", err)
	}
	statusCode := res.StatusCode

	if statusCode == 200 {
		fmt.Printf("Server Status is OK\n")
	} else {
		fmt.Printf("Server Status Code: %v\n", statusCode)
	}
}
