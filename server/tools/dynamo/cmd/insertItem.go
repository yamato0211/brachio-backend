/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/pkg/dynamo"
)

// insertItemCmd represents the insertItem command
var insertItemCmd = &cobra.Command{
	Use:   "insertItem",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
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

		items := []*model.MasterItem{
			{
				ItemID:      model.MasterItemID("pack-power"),
				Name:        "パックパワー",
				Description: "10個で1パック引けるよ",
				ImageURL:    "https://pokepoke.kurichi.dev/images/kizuku-piece.avif",
			},
		}

		tbl := dc.Table("MasterItems")
		for _, item := range items {
			if err := tbl.Put(item).Run(cmd.Context()); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("inserted")
	},
}

func init() {
	rootCmd.AddCommand(insertItemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertItemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertItemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
