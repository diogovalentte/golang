/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/diogovalentte/golang/gRPC_server/pb"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

// checkStatusCmd represents the check-status command
var checkStatusCmd = &cobra.Command{
	Use:   "check-status",
	Short: "Check status of gRPC server",
	Long:  `Make a gRPC request to function CheckStatus() to find if gRPC server is OK`,
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value
		checkStatus(&port)
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
func checkStatus(port *pflag.Value) {
	URL := fmt.Sprintf("localhost:%v", *port)
	conn, err := grpc.Dial(URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCheckStatusServiceClient(conn)

	request := &pb.CheckStatusRequest{}
	res, err := client.CheckStatus(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during execution: %v", err)
	}
	if res.Code == 200 {
		fmt.Println("Server is OK")
	}
}
