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

// deckCmd represents the deck command
var deckCmd = &cobra.Command{
	Use:   "deck",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example`,
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

		decks := []*model.Deck{
			{
				DeckID:          model.DeckID("template-deck-1"),
				Name:            "テンプレートデッキ 脳筋・ギャンブル",
				UserID:          model.UserID("master"),
				ThumbnailCardID: model.NewMasterCardID("kizuku"),
				Color:           model.MonsterTypeAlchohol,
				Energies:        []model.MonsterType{model.MonsterTypeAlchohol, model.MonsterTypeMuscle},
				MasterCardIDs: []model.MasterCardID{
					model.NewMasterCardID("yamato"),
					model.NewMasterCardID("yamato"),
					model.NewMasterCardID("kizuku"),
					model.NewMasterCardID("kizuku"),
					model.NewMasterCardID("dolly"),
					model.NewMasterCardID("wall"),
					model.NewMasterCardID("wall"),
					model.NewMasterCardID("strict-mode"),
					model.NewMasterCardID("strict-mode"),
					model.NewMasterCardID("programming-school"),
					model.NewMasterCardID("programming-school"),
					model.NewMasterCardID("personal-trainer"),
					model.NewMasterCardID("personal-trainer"),
					model.NewMasterCardID("protein"),
					model.NewMasterCardID("protein"),
					model.NewMasterCardID("sake-bottle"),
					model.NewMasterCardID("sake-bottle"),
					model.NewMasterCardID("bartender"),
					model.NewMasterCardID("bartender"),
					model.NewMasterCardID("izakaya-taisyo"),
				},
			},
			{
				DeckID:          model.DeckID("template-deck-2"),
				Name:            "テンプレートデッキ 金・知識",
				UserID:          model.UserID("master"),
				ThumbnailCardID: model.NewMasterCardID("tosaken-ex"),
				Color:           model.MonsterTypeMoney,
				Energies:        []model.MonsterType{model.MonsterTypeMoney, model.MonsterTypeKnowledge},
				MasterCardIDs: []model.MasterCardID{
					model.NewMasterCardID("tosaken-ex"),
					model.NewMasterCardID("tosaken-ex"),
					model.NewMasterCardID("kurichi"),
					model.NewMasterCardID("kurichi"),
					model.NewMasterCardID("dolly"),
					model.NewMasterCardID("wall"),
					model.NewMasterCardID("wall"),
					model.NewMasterCardID("strict-mode"),
					model.NewMasterCardID("strict-mode"),
					model.NewMasterCardID("programming-school"),
					model.NewMasterCardID("programming-school"),
					model.NewMasterCardID("yamikin-gyosya"),
					model.NewMasterCardID("yamikin-gyosya"),
					model.NewMasterCardID("credit-card"),
					model.NewMasterCardID("credit-card"),
					model.NewMasterCardID("refactoring"),
					model.NewMasterCardID("refactoring"),
					model.NewMasterCardID("ci-cd-pipeline"),
					model.NewMasterCardID("ci-cd-pipeline"),
					model.NewMasterCardID("security-engineer"),
				},
			},
		}

		tbl := dc.Table("Decks")
		for _, d := range decks {
			if err := tbl.Put(d).Run(cmd.Context()); err != nil {
				log.Fatal(err)
			}
		}

		log.Println("success")
	},
}

func init() {
	rootCmd.AddCommand(deckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
