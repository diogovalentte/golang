/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/diogovalentte/golang/gRPC_server/pb"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Do some simple math operation",
	Long:  `Make a gRPC request to function Calculate() to do a math operation `,
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value
		calculate(&port, args)
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
func calculate(port *pflag.Value, args []string) {
	URL := fmt.Sprintf("localhost:%v", *port)
	conn, err := grpc.Dial(URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculateServiceClient(conn)

	intNum1, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Error loading num1: %v", err)
	}
	num1 := int64(intNum1)
	intNum2, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Error loading num2: %v", err)
	}
	num2 := int64(intNum2)
	operator := args[2]

	request := &pb.CalculateRequest{
		Num1:     num1,
		Num2:     num2,
		Operator: operator,
	}
	res, err := client.Calculate(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during execution: %v", err)
	}

	result := fmt.Sprintf("%v %v %v = %v", res.Num1, res.Operator, res.Num2, res.Result)
	fmt.Println(result)
}
