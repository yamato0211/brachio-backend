/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/pkg/dynamo"
)

// presentCmd represents the present command
var presentCmd = &cobra.Command{
	Use:   "present",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		dc, err := dynamo.New(cmd.Context(), &dynamo.Config{
			IsLocal: cfg.Common.IsLocal,
			Region:  cfg.Dynamo.Region,
		})
		if err != nil {
			log.Fatal(err)
		}

		// presents := []*model.Present{
		// 	{
		// 		PresentID:       model.NewPresentID(),
		// 		Time:            int(time.Now().Unix()),
		// 		ReceivedUserIDs: []model.UserID{},
		// 		ItemID:          model.MasterItemID("pack-power"),
		// 		ItemCount:       100,
		// 		Message:         "運営からのプレゼントです！",
		// 	},
		// }

		// tbl := dc.Table("Presents")
		// for _, p := range presents {
		// 	if err := tbl.Put(p).Run(cmd.Context()); err != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		// log.Println("inserted")

		tbl := dc.Table("Users")
		var user model.User
		if err := tbl.Get("UserId", "97843a68-e051-7048-3cff-fcd6162d57d4").One(cmd.Context(), &user); err != nil {
			log.Fatal(err)
		}
		user.ItemIDsWithCount["pack-power"] += 10000
		if err := tbl.Put(&user).Run(cmd.Context()); err != nil {
			log.Fatal(err)
		}
		log.Println("updated")
	},
}

func init() {
	rootCmd.AddCommand(presentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// presentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// presentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
