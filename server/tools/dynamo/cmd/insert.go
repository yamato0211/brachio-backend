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
			IsLocal: cfg.Common.IsLocal,
			Region:  cfg.Dynamo.Region,
		})
		if err != nil {
			log.Fatal(err)
		}

		var users []model.MasterCard = []model.MasterCard{
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "Kizuku",
				CardType:     model.CardTypeMonster,
				Rarity:       8,
				HP:           110,
				Type:         model.MonsterTypeAlchohol,
				Weakness:     model.MonsterTypeMuscle,
				RetreatCost:  1,
				Ability: &model.Ability{
					Name: "技術的には可能です",
					Text: "自分の番に1回使える。自分のエネルギーゾーンからランダムにエネルギーを5個出し、このラムモンにつける。",
				},
				Skills: []model.Skill{
					{
						Name: "博打",
						Text: "コインを1回投げ表なら相手のベンチラムモン全員にも200ダメージ、裏ならこのラムモンについているエネルギーをすべてトラッシュする",
						Cost: []model.MonsterType{
							model.MonsterTypeAlchohol,
							model.MonsterTypeMoney,
							model.MonsterTypeKnowledge,
							model.MonsterTypeMuscle,
							model.MonsterTypePopularity,
						},
						Damage:       200,
						DamageOption: "x",
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "くりち",
				CardType:     model.CardTypeMonster,
				Rarity:       8,
				HP:           140,
				Type:         model.MonsterTypeKnowledge,
				Weakness:     model.MonsterTypePopularity,
				IsEx:         false,
				RetreatCost:  2,
				Skills: []model.Skill{
					{
						Name:         "Clean Architecture",
						Text:         "ウラが出るまでコインを投げ、オモテの数x30ダメージ",
						Damage:       30,
						DamageOption: "x",
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
						},
					},
					{
						Name:   "コードレビュー",
						Text:   "相手の手札からランダムに2枚トラッシュ",
						Damage: 70,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "ヤマト",
				Description: `
					筋肉.
					生息地はジム.
				`,
				Rarity:      8,
				CardType:    model.CardTypeMonster,
				HP:          200,
				Type:        model.MonsterTypeMuscle,
				Weakness:    model.MonsterTypeAlchohol,
				IsEx:        false,
				RetreatCost: 3,
				Skills: []model.Skill{
					{
						Name:   "ベンチプレス",
						Text:   "自分にも20ダメージを与える",
						Damage: 999,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "土佐犬",
				Rarity:       1,
				CardType:     model.CardTypeMonster,
				Description: `
					金と女に目がない.
				`,
				Type:        model.MonsterTypeMoney,
				Weakness:    model.MonsterTypeMoney,
				IsEx:        true,
				RetreatCost: 0,
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "土佐犬",
				Rarity:       1,
				CardType:     model.CardTypeMonster,
				IsEx:         false,
				Type:         model.MonsterTypeMoney,
				Weakness:     model.MonsterTypeMoney,
				RetreatCost:  0,
				HP:           50,
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "ドリー",
				Rarity:       7,
				CardType:     model.CardTypeMonster,
				Type:         model.MonsterTypePopularity,
				Weakness:     model.MonsterTypeMoney,
				RetreatCost:  1,
				Ability: &model.Ability{
					Name: "さめはだ",
					Text: "このラムモンが、バトル場で相手のラムモンからワザのダメージを受けたとき、ワザを使ったラムモンに40ダメージ",
				},
				Skills: []model.Skill{
					{
						Name:   "",
						Text:   "",
						Damage: 160,
						Cost: []model.MonsterType{
							model.MonsterTypeAlchohol,
							model.MonsterTypePopularity,
							model.MonsterTypePopularity,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "駆け出しエンジニア",
				Description:  "駆け出しのエンジニア.\n何者にでもなれる.",
				Rarity:       1,
				CardType:     model.CardTypeMonster,
				HP:           40,
				Type:         model.MonsterTypePopularity,
				Weakness:     model.MonsterTypeMoney,
				RetreatCost:  1,
				Skills: []model.Skill{
					{
						Name:   "#駆け出しエンジニア",
						Damage: 10,
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "セキュリティエンジニア",
				Description:  "滅多に現れない\n",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           90,
				Type:         model.MonsterTypeAlchohol,
				Weakness:     model.MonsterTypePopularity,
				RetreatCost:  1,
				Skills:       []model.Skill{},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "フロントエンドエンジニア",
				Description:  "Safariが嫌い\n",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           90,
				Type:         model.MonsterTypeMoney,
				Weakness:     model.MonsterTypeKnowledge,
				RetreatCost:  2,
				Skills: []model.Skill{
					{
						Name:   "lint設定",
						Text:   "自分のラムモン全員のHPを10回復",
						Damage: 30,
						Cost: []model.MonsterType{
							model.MonsterTypeMoney,
						},
					},
					{
						Name:   "デザイナーへの悪口",
						Text:   "",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "SRE",
				Description:  "チームメンバーの開発体験向上\nのことだけを考えている。",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           120,
				Type:         model.MonsterTypeKnowledge,
				Weakness:     model.MonsterTypeMuscle,
				RetreatCost:  2,
				Skills: []model.Skill{
					{
						Name:   "環境構築",
						Text:   "コインを3回投げ、オモテの数ぶんの[知識]エネルギーを自分のエネルギーゾーンから出し、ベンチの[知識]ラムモンに好きなようにつける。",
						Damage: 0,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
						},
					},
					{
						Name:   "デプロイ",
						Damage: 60,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
							model.MonsterTypeKnowledge,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "フルスタックエンジニア",
				Description:  "全ての領域を知る天才。",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           130,
				Type:         model.MonsterTypeNull,
				RetreatCost:  2,
				Ability: &model.Ability{
					Name: "広く浅く",
					Text: "このラムモンがたねラムモンから受けるワザのダメージを-20、2進化ラムモンから受けるワザのダメージを+10する。",
				},
				Skills: []model.Skill{
					{
						Name:   "",
						Text:   "自分の山札からラムモンをランダムに1枚、手札に加える。",
						Damage: 20,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "ニート",
				Description:  "全てを諦めたひと。",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           50,
				Type:         model.MonsterTypeNull,
				Weakness:     model.MonsterTypePopularity,
				Skills: []model.Skill{
					{
						Name:   "クソツイ",
						Damage: 5,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "パチカス",
				CardType:     model.CardTypeMonster,
				Rarity:       4,
				HP:           40,
				Type:         model.MonsterTypeAlchohol,
				Weakness:     model.MonsterTypeMuscle,
				RetreatCost:  1,

				Skills: []model.Skill{
					{
						Name: "パチンコ",
						Text: "コインを1回投げ裏ならこの技は失敗する",
						Cost: []model.MonsterType{
							model.MonsterTypeMoney,
						},
						Damage:       30,
						DamageOption: "x",
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "ヤニカス",
				CardType:     model.CardTypeMonster,
				Rarity:       4,
				HP:           40,
				Type:         model.MonsterTypeAlchohol,
				Weakness:     model.MonsterTypeMuscle,
				RetreatCost:  1,
				Ability: &model.Ability{
					Name: "タバコ休憩",
					Text: "自分の番に1回使える。[金]エネルギーを1つトラッシュする代わりにこのラムモンのHPを20回復する",
				},
				Skills: []model.Skill{
					{
						Name: "副流煙",
						Text: "相手のベンチラムモン1匹にも10ダメージ",
						Cost: []model.MonsterType{
							model.MonsterTypeAlchohol,
							model.MonsterTypeMoney,
						},
						Damage:       50,
						DamageOption: "x",
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "酒カス",
				CardType:     model.CardTypeMonster,
				Rarity:       4,
				HP:           40,
				Type:         model.MonsterTypeAlchohol,
				Weakness:     model.MonsterTypeMuscle,
				RetreatCost:  1,
				Ability: &model.Ability{
					Name: "酒カスエンジニア",
					Text: "このラムモンに[酒]エネルギーがついている、かつ、HPが10より多い場合、HPが0以下になるダメージを受けてもHP10で耐える。",
				},
				Skills: []model.Skill{
					{
						Name: "博打",
						Text: "コインを1回投げ表なら相手のラムモン全体に50ダメージ、裏ならこのラムモンについているエネルギーをすべてトラッシュする",
						Cost: []model.MonsterType{
							model.MonsterTypeAlchohol,
							model.MonsterTypeMoney,
						},
						Damage:       50,
						DamageOption: "x",
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "勤勉学生",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           50,
				Type:         model.MonsterTypeKnowledge,
				Weakness:     model.MonsterTypeMuscle,
				Skills: []model.Skill{
					{
						Name:   "勉強する",
						Damage: 30,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "JTC中堅社員",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           80,
				Type:         model.MonsterTypeKnowledge,
				Weakness:     model.MonsterTypeMuscle,
				Skills: []model.Skill{
					{
						Name:   "まじめに働く",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				Name:         "ベンチャー社長",
				Description:  "これまで真面目に働いてきたが、後悔したくないという思いから大きく挑戦した。",
				Rarity:       4,
				CardType:     model.CardTypeMonster,
				HP:           130,
				Type:         model.MonsterTypeKnowledge,
				Weakness:     model.MonsterTypeMuscle,
				Skills: []model.Skill{
					{
						Name:   "大盤振る舞い",
						Text:   "相手のベンチラムモン全員にも10ダメージ。",
						Damage: 120,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
						},
					},
				},
			},
			{
				MasterCardID: model.NewMasterCardID(),
				CardType:     model.CardTypeGoods,
				Name:         "オライリー本",
				Text:         "この番、自分の[知識]ラムモンが使うワザの、相手",
			},
			{
				MasterCardID: model.NewMasterCardID(),
				CardType:     model.CardTypeGoods,
				Name:         "エナジードリンク",
				Text:         "",
			},
			{
				MasterCardID: model.NewMasterCardID(),
				CardType:     model.CardTypeGoods,
				Name:         "プロテイン",
				Text:         "",
			},
			{
				MasterCardID: model.NewMasterCardID(),
				CardType:     model.CardTypeGoods,
				Name:         "プロテイン",
				Text:         "",
			},
		}

		tbl := dc.Table("Cards")

		for _, user := range users {
			if err := tbl.Put(user).Run(cmd.Context()); err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("inserted")
	},
}

func init() {
	// rootCmd.Flags().StringP("dynamo-endpoint", "dn", "http://localhost:8000", "dynamo db endpoint")
	rootCmd.AddCommand(insertCmd)
}
