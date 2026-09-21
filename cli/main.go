package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "go-shop"}
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	rootCmd.AddCommand(cmdInsertProduct)
	rootCmd.AddCommand(cmdListProducts)
	rootCmd.AddCommand(cmdGrpcInsertProduct)
	rootCmd.AddCommand(cmdGrpcListProducts)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}
