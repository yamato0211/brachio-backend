/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/pkg/dynamo"
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
		cfg, err := config.GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		dc, err := dynamo.New(cmd.Context(), &dynamo.Config{
			IsLocal:  cfg.Common.IsLocal,
			Region:   cfg.Dynamo.Region,
			Endpoint: cfg.Dynamo.Endpoint,
		})
		if err != nil {
			log.Fatal(err)
		}

		// var users []model.MasterCard = []model.MasterCard{
		// 	{
		// 		MasterCardID: model.NewMasterCardID(),
		// 		Name:         "Alice",
		// 	},
		// 	{
		// 		MasterCardID: model.NewMasterCardID(),
		// 		Name:         "Bob",
		// 	},
		// 	{
		// 		MasterCardID: model.NewMasterCardID(),
		// 		Name:         "Charlie",
		// 	},
		// 	{
		// 		MasterCardID: model.NewMasterCardID(),
		// 		Name:         "Diana",
		// 	},
		// 	{
		// 		MasterCardID: model.NewMasterCardID(),
		// 		Name:         "Eve",
		// 	},
		// }

		// type User struct {
		// 	UserID string `dynamo:"UserId,hash"`
		// 	Name   string `dynamo:"Name"`
		// }
		// var users []User

		tbs, err := dc.ListTables().All(cmd.Context())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tbs)
	},
}

func init() {
	// rootCmd.Flags().StringP("dynamo-endpoint", "dn", "http://localhost:8000", "dynamo db endpoint")
	rootCmd.AddCommand(insertCmd)
}
