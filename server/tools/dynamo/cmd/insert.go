/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "this command is used to insert data into dynamo db",
	Long: `this command is used to insert data into dynamo db
	For example:
		dynamo-cli insert
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("insert called")
	},
}

func init() {
	// rootCmd.Flags().StringP("dynamo-endpoint", "dn", "http://localhost:8000", "dynamo db endpoint")
	rootCmd.AddCommand(insertCmd)
}
