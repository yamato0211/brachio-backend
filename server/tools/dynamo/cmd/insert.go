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
				MasterCardID: model.NewMasterCardID("kizuku"),
				Name:         "Kizuku",
				CardType:     model.CardTypeMonster,
				Rarity:       8,
				HP:           130,
				ImageURL:     "https://pokepoke.kurichi.dev/images/kizuku-piece.avif",
				Type:         model.MonsterTypeAlchohol,
				Weakness:     model.MonsterTypeMuscle,
				RetreatCost:  1,
				Ability: &model.Ability{
					Name: "技術的には可能です",
					Text: "自分の番に1回使える。自分のエネルギーゾーンからランダムにエネルギーを5個出し、このポケモンにつける。",
				},
				Skills: []*model.Skill{
					{
						Name: "博打",
						Text: "コインを1回投げ表なら相手のベンチポケモン全員にも200ダメージ、裏ならこのポケモンについているエネルギーをすべてトラッシュする",
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
				IsEx:    true,
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("kurichi"),
				Name:         "くりち",
				CardType:     model.CardTypeMonster,
				Rarity:       8,
				HP:           160,
				ImageURL:     "https://pokepoke.kurichi.dev/images/kurichi.avif",
				Type:         model.MonsterTypeKnowledge,
				Weakness:     model.MonsterTypePopularity,
				IsEx:         true,
				RetreatCost:  2,
				Skills: []*model.Skill{
					{
						Name:   "Clean Architecture",
						Text:   "このポケモンから2エネルギートラッシュする。相手のバトル場のポケモンから1枚選び、そのカードを山札に戻す。",
						Damage: 0,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
						},
					},
					{
						Name:   "コードレビュー",
						Text:   "相手の手札からランダムに2枚トラッシュ。トラッシュできない場合は山札をトラッシュ",
						Damage: 0,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("yamato"),
				Name:         "ヤマト",
				Description:  "筋肉.\n生息地はジム.",
				Rarity:       8,
				ImageURL:     "https://pokepoke.kurichi.dev/images/yamato.avif",
				CardType:     model.CardTypeMonster,
				HP:           200,
				Type:         model.MonsterTypeMuscle,
				Weakness:     model.MonsterTypeAlchohol,
				IsEx:         true,
				RetreatCost:  3,
				Skills: []*model.Skill{
					{
						Name:   "ベンチプレス",
						Text:   "相手の場のポケモン全てに999ダメージ",
						Damage: 999,
						Cost: []model.MonsterType{
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("tosaken-ex"),
				Name:         "土佐犬",
				Rarity:       1,
				CardType:     model.CardTypeMonster,
				Description:  "金と女に目がない.",
				Type:         model.MonsterTypeMoney,
				Weakness:     model.MonsterTypeMoney,
				IsEx:         true,
				RetreatCost:  0,
				HP:           80,
				ImageURL:     "https://pokepoke.kurichi.dev/images/tosaken-ex.avif",
				Ability: &model.Ability{
					Name: "金がすべて",
					Text: "このポケモンがワザを使うとき、このポケモンについている[金]エネルギーの数x10ダメージ追加。",
				},
				Skills: []*model.Skill{
					{
						Name:   "バイクで突っ込む",
						Text:   "相手のベンチポケモン全員にも50ダメージ。このポケモンについているエネルギーをすべてトラッシュ。このポケモンにも70ダメージ",
						Damage: 200,
						Cost: []model.MonsterType{
							model.MonsterTypeMoney,
							model.MonsterTypeMoney,
							model.MonsterTypeMoney,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("tosaken"),
				Name:         "土佐犬",
				Rarity:       1,
				CardType:     model.CardTypeMonster,
				IsEx:         false,
				Type:         model.MonsterTypeMoney,
				Weakness:     model.MonsterTypeMoney,
				RetreatCost:  0,
				HP:           30,
				ImageURL:     "https://pokepoke.kurichi.dev/images/tosaken.avif",
				Skills: []*model.Skill{
					{
						Name:   "かみつく",
						Damage: 10,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("dolly"),
				Name:         "ドリー",
				Rarity:       7,
				HP:           160,
				IsEx:         true,
				ImageURL:     "https://pokepoke.kurichi.dev/images/dlly.gif",
				CardType:     model.CardTypeMonster,
				Type:         model.MonsterTypePopularity,
				Weakness:     model.MonsterTypeMoney,
				RetreatCost:  3,
				// Ability: &model.Ability{
				// 	Name: "さめはだ",
				// 	Text: "このポケモンが、バトル場で相手のポケモンからワザのダメージを受けたとき、ワザを使ったポケモンに40ダメージ",
				// },
				Skills: []*model.Skill{
					{
						Name:   "ふいうち",
						Text:   "",
						Damage: 40,
						Cost: []model.MonsterType{
							model.MonsterTypePopularity,
						},
					},
					{
						Name:   "ひっかく",
						Text:   "コインを2回投げ2回ともウラなら、このポケモンにも100ダメージ",
						Damage: 200,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypePopularity,
							model.MonsterTypePopularity,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("newbie-engineer"),
				Name:         "駆け出しエンジニア",
				Description:  "駆け出しのエンジニア.\n何者にでもなれる.",
				Rarity:       1,
				CardType:     model.CardTypeMonster,
				HP:           40,
				ImageURL:     "https://pokepoke.kurichi.dev/images/kakedasi.avif",

				Type:     model.MonsterTypePopularity,
				Weakness: model.MonsterTypeMoney,
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("security-engineer"),
					model.NewMasterCardID("frontend-engineer"),
					model.NewMasterCardID("designer"),
					model.NewMasterCardID("backend-engineer"),
					model.NewMasterCardID("backend-engineer"),
					model.NewMasterCardID("sre"),
					model.NewMasterCardID("full-stack-engineer"),
				},
				RetreatCost: 1,
				Skills: []*model.Skill{
					{
						Name:   "#駆け出しエンジニア",
						Damage: 10,
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("security-engineer"),
				Name:         "セキュリティエンジニア",
				Description:  "滅多に現れない\n",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           90,
				ImageURL:     "https://pokepoke.kurichi.dev/images/secure.avif",

				Type:        model.MonsterTypeAlchohol,
				Weakness:    model.MonsterTypePopularity,
				RetreatCost: 0,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("newbie-engineer"),
				},
				Skills: []*model.Skill{
					{
						Name:   "リバースエンジニアリング",
						Text:   "自分のトラッシュのカードを1枚選び、手札に加える",
						Damage: 0,
						Cost: []model.MonsterType{
							model.MonsterTypeAlchohol,
							model.MonsterTypeAlchohol,
						},
					},
					{
						Name:   "脆弱性診断",
						Damage: 70,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeAlchohol,
							model.MonsterTypeAlchohol,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("frontend-engineer"),
				Name:         "フロントエンドエンジニア",
				Description:  "Safariが嫌い\n",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           90,
				ImageURL:     "https://pokepoke.kurichi.dev/images/frontend.avif",
				Type:         model.MonsterTypeMoney,
				Weakness:     model.MonsterTypeKnowledge,
				RetreatCost:  2,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("newbie-engineer"),
				},
				Skills: []*model.Skill{
					{
						Name:   "lint設定",
						Text:   "自分のポケモン全員のHPを10回復",
						Damage: 30,
						Cost: []model.MonsterType{
							model.MonsterTypeMoney,
						},
					},
					{
						Name:         "デザイナーへの悪口",
						Text:         "相手のバトルポケモンが「デザイナー」のとき、50ダメージ追加",
						Damage:       50,
						DamageOption: "+50",
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeMoney,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("designer"),
				Name:         "デザイナー",
				Description:  "CSSはフロントエンドの仕事\n",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           90,
				Type:         model.MonsterTypeMuscle,
				Weakness:     model.MonsterTypeAlchohol,
				RetreatCost:  2,
				ImageURL:     "https://pokepoke.kurichi.dev/images/designer.jpeg",
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("newbie-engineer"),
				},
				Skills: []*model.Skill{
					{
						Name:   "あっそこのUI変えていいですか？",
						Text:   "相手のベンチポケモン全員に10ダメージ",
						Damage: 10,
						Cost: []model.MonsterType{
							model.MonsterTypeMuscle,
						},
					},
					{
						Name:         "フロントエンドエンジニアへの悪口",
						Text:         "相手のバトルポケモンが「フロントエンドエンジニア」のとき、50ダメージ追加",
						Damage:       50,
						DamageOption: "+50",
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeMuscle,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("backend-engineer"),
				Name:         "バックエンドエンジニア",
				Description:  "動的型付け言語が嫌い\n",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           90,
				ImageURL:     "https://pokepoke.kurichi.dev/images/backend.avif",

				Type:     model.MonsterTypePopularity,
				Weakness: model.MonsterTypeMoney,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("newbie-engineer"),
				},
				RetreatCost: 2,
				Skills: []*model.Skill{
					{
						Name:         "DB設計",
						Text:         "このポケモンがダメージを受けているなら、60ダメージ追加",
						Damage:       40,
						DamageOption: "+60",
						Cost: []model.MonsterType{
							model.MonsterTypePopularity,
							model.MonsterTypePopularity,
						},
					},
					{
						Name:   "デザイナーへの悪口",
						Text:   "",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypePopularity,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("sre"),
				Name:         "SRE",
				Description:  "チームメンバーの開発体験向上\nのことだけを考えている。",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           120,
				ImageURL:     "https://pokepoke.kurichi.dev/images/sre.avif",

				Type:     model.MonsterTypeKnowledge,
				Weakness: model.MonsterTypeMuscle,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("newbie-engineer"),
				},
				RetreatCost: 2,
				Skills: []*model.Skill{
					{
						Name:   "環境構築",
						Text:   "コインを3回投げ、オモテの数ぶんの[知識]エネルギーを自分のエネルギーゾーンから出し、ベンチの[知識]ポケモンに好きなようにつける。",
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
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("full-stack-engineer"),
				Name:         "フルスタックエンジニア",
				Description:  "全ての領域を知る天才。",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           130,
				ImageURL:     "https://pokepoke.kurichi.dev/images/fullstuck.avif",

				Type:        model.MonsterTypeNull,
				RetreatCost: 2,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("newbie-engineer"),
				},
				Ability: &model.Ability{
					Name: "広く浅く",
					Text: "このポケモンがたねポケモンから受けるワザのダメージを-20、2進化ポケモンから受けるワザのダメージを+10する。",
				},
				Skills: []*model.Skill{
					{
						Name:   "全知全能",
						Text:   "自分の山札からポケモンをランダムに1枚、手札に加える。",
						Damage: 20,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("neet"),
				Name:         "ニート",
				Description:  "全てを諦めたひと。",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				ImageURL:     "https://pokepoke.kurichi.dev/images/neet.avif",

				HP:          50,
				Type:        model.MonsterTypeNull,
				Weakness:    model.MonsterTypePopularity,
				RetreatCost: 3,
				Skills: []*model.Skill{
					{
						Name:   "クソツイ",
						Damage: 5,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("pachikasu"),
				Name:         "パチカス",
				CardType:     model.CardTypeMonster,
				Rarity:       2,
				HP:           30,
				ImageURL:     "https://pokepoke.kurichi.dev/images/pachinko.avif",

				Type:        model.MonsterTypeAlchohol,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 0,
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("yanikasu"),
				},
				Skills: []*model.Skill{
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
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("yanikasu"),
				Name:         "ヤニカス",
				CardType:     model.CardTypeMonster,
				Rarity:       3,
				HP:           60,
				ImageURL:     "https://pokepoke.kurichi.dev/images/タバコ.avif",

				Type:        model.MonsterTypeAlchohol,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 0,
				Ability: &model.Ability{
					Name: "タバコ休憩",
					Text: "自分の番に1回使える。[金]エネルギーを1つトラッシュする代わりにこのポケモンのHPを20回復する",
				},
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("pachikasu"),
				},
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("sakekasu"),
				},
				Skills: []*model.Skill{
					{
						Name: "副流煙",
						Text: "相手のベンチポケモン1匹にも10ダメージ",
						Cost: []model.MonsterType{
							model.MonsterTypeAlchohol,
							model.MonsterTypeMoney,
						},
						Damage:       40,
						DamageOption: "x",
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("sakekasu"),
				Name:         "酒カスエンジニア",
				CardType:     model.CardTypeMonster,
				Rarity:       4,
				HP:           110,
				ImageURL:     "https://pokepoke.kurichi.dev/images/sakekasu.avif",

				Type:        model.MonsterTypeAlchohol,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 1,
				Ability: &model.Ability{
					Name: "酒の力",
					Text: "自分のエネルギーゾーンからこのポケモンに[酒]エネルギーをつけるたび、このポケモンのHPを10回復する",
				},
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("yanikasu"),
				},
				Skills: []*model.Skill{
					{
						Name: "飲酒駆動開発",
						Text: "このポケモンについている[金]エネルギーの数x20ダメージ追加",
						Cost: []model.MonsterType{
							model.MonsterTypeAlchohol,
							model.MonsterTypeAlchohol,
							model.MonsterTypeAlchohol,
						},
						Damage:       100,
						DamageOption: "+",
					},
				},
				SubType: model.MonsterSubTypeStage2,
			},
			{
				MasterCardID: model.NewMasterCardID("student"),
				Name:         "勤勉学生",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           50,
				ImageURL:     "https://pokepoke.kurichi.dev/images/student.avif",

				Type:        model.MonsterTypeKnowledge,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 1,
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("jtc"),
				},
				Skills: []*model.Skill{
					{
						Name:   "勉強する",
						Damage: 30,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("jtc"),
				Name:         "JTC中堅社員",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           80,
				ImageURL:     "https://pokepoke.kurichi.dev/images/syakaijin.avif",

				Type:        model.MonsterTypeKnowledge,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 1,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("student"),
				},
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("venture-president"),
				},
				Skills: []*model.Skill{
					{
						Name:   "まじめに働く",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("venture-president"),
				Name:         "ベンチャー社長",
				Description:  "これまで真面目に働いてきたが、後悔したくないという思いから大きく挑戦した。",
				Rarity:       4,
				CardType:     model.CardTypeMonster,
				HP:           130,
				ImageURL:     "https://pokepoke.kurichi.dev/images/syatyo.avif",

				Type:        model.MonsterTypeKnowledge,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 1,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("jtc"),
				},
				Skills: []*model.Skill{
					{
						Name:   "大盤振る舞い",
						Text:   "相手のベンチポケモン全員にも20ダメージ。",
						Damage: 120,
						Cost: []model.MonsterType{
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
							model.MonsterTypeKnowledge,
						},
					},
				},
				SubType: model.MonsterSubTypeStage2,
			},

			{
				MasterCardID: model.NewMasterCardID("garigari"),
				Name:         "ガリガリエンジニア",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           50,
				ImageURL:     "https://pokepoke.kurichi.dev/images/garigari.avif",

				Type:        model.MonsterTypeMuscle,
				Weakness:    model.MonsterTypeAlchohol,
				RetreatCost: 1,
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("muchimuchi"),
				},
				Skills: []*model.Skill{
					{
						Name:   "筋トレ",
						Damage: 10,
						Cost: []model.MonsterType{
							model.MonsterTypeMuscle,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("muchimuchi"),
				Name:         "ムチムチエンジニア",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           90,
				ImageURL:     "https://pokepoke.kurichi.dev/images/mutimurti.avif",

				Type:        model.MonsterTypeMuscle,
				Weakness:    model.MonsterTypeAlchohol,
				RetreatCost: 2,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("garigari"),
				},
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("mukimuki"),
				},
				Skills: []*model.Skill{
					{
						Name:   "筋トレ",
						Damage: 40,
						Cost: []model.MonsterType{
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("mukimuki"),
				Name:         "ムキムキエンジニア",
				Rarity:       4,
				CardType:     model.CardTypeMonster,
				HP:           170,
				ImageURL:     "https://pokepoke.kurichi.dev/images/mukimuki.avif",

				Type:        model.MonsterTypeMuscle,
				Weakness:    model.MonsterTypeAlchohol,
				RetreatCost: 3,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("muchimuchi"),
				},
				Ability: &model.Ability{
					Name: "パワーーー！！！",
					Text: "このポケモンがいるかぎり、自分の[筋肉]ポケモンが使うワザのダメージを+30する",
				},
				Skills: []*model.Skill{
					{
						Name:   "筋トレ",
						Damage: 100,
						Cost: []model.MonsterType{
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
						},
					},
				},
				SubType: model.MonsterSubTypeStage2,
			},
			{
				MasterCardID: model.NewMasterCardID("multi-business"),
				Name:         "マルチ商法",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           50,
				ImageURL:     "https://pokepoke.kurichi.dev/images/multi_syouhou_kanyuu.avif",

				Type:        model.MonsterTypeMoney,
				Weakness:    model.MonsterTypePopularity,
				RetreatCost: 0,
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("info-product"),
				},
				Skills: []*model.Skill{
					{
						Name:   "ともだちをさがす",
						Text:   "自分の山札から[金]ポケモンをランダムに1枚、手札に加える",
						Damage: 0,
						Cost: []model.MonsterType{
							model.MonsterTypeMoney,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("info-product"),
				Name:         "情報商材屋",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           80,
				ImageURL:     "https://pokepoke.kurichi.dev/images/johosyozai.avif",

				Type:        model.MonsterTypeMoney,
				Weakness:    model.MonsterTypePopularity,
				RetreatCost: 0,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("multi-business"),
				},
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("high-tax-payer"),
				},
				Ability: &model.Ability{
					Name: "荒稼ぎ",
					Text: "自分の番に1回使える。自分のエネルギーゾーンから[金]エネルギーを1個出し、このポケモンにつける。",
				},
				Skills: []*model.Skill{
					{
						Name:   "有料note販売",
						Text:   "自分のエネルギーゾーンから[金]エネルギーを1個出し、このポケモンにつける。",
						Damage: 20,
						Cost: []model.MonsterType{
							model.MonsterTypeMoney,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("high-tax-payer"),
				Name:         "高額納税者",
				Rarity:       4,
				CardType:     model.CardTypeMonster,
				HP:           160,
				ImageURL:     "https://pokepoke.kurichi.dev/images/kougakunouzeisya.avif",

				Type:        model.MonsterTypeMoney,
				Weakness:    model.MonsterTypePopularity,
				RetreatCost: 3,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("info-product"),
				},
				Skills: []*model.Skill{
					{
						Name:   "納税",
						Text:   "このポケモンから[金]エネルギーを3個トラッシュし、このポケモンのHPを100回復",
						Damage: 0,
						Cost: []model.MonsterType{
							model.MonsterTypeMoney,
							model.MonsterTypeMoney,
							model.MonsterTypeMoney,
						},
					},
					{
						Name:   "札束ビンタ",
						Text:   "このポケモンから[金]エネルギーを2個トラッシュ",
						Damage: 100,
						Cost: []model.MonsterType{
							model.MonsterTypeMoney,
							model.MonsterTypeMoney,
						},
					},
				},
				SubType: model.MonsterSubTypeStage2,
			},
			{
				MasterCardID: model.NewMasterCardID("startup-cto"),
				Name:         "スタートアップCTO",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           60,
				ImageURL:     "https://pokepoke.kurichi.dev/images/start-up-cto.avif",

				Type:        model.MonsterTypePopularity,
				Weakness:    model.MonsterTypeMoney,
				RetreatCost: 1,
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("venture-cto"),
				},
				Skills: []*model.Skill{
					{
						Name:   "がむしゃらに働く",
						Text:   "このポケモンにも50ダメージ",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypePopularity,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("venture-cto"),
				Name:         "ベンチャーCTO",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           120,
				ImageURL:     "https://pokepoke.kurichi.dev/images/bencha-cto.avif",

				Type:        model.MonsterTypePopularity,
				Weakness:    model.MonsterTypeMoney,
				RetreatCost: 2,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("startup-cto"),
				},
				EvolvesTo: []model.MasterCardID{
					model.NewMasterCardID("big-company-cto"),
				},
				Ability: &model.Ability{
					Name: "まだ諦めない",
					Text: "このポケモンのワザの効果により、このポケモンがダメージを受けHPが0以下にならなかった場合、このポケモンのHPを50回復する。",
				},
				Skills: []*model.Skill{
					{
						Name:   "血反吐を吐く",
						Text:   "このポケモンにも100ダメージ",
						Damage: 100,
						Cost: []model.MonsterType{
							model.MonsterTypePopularity,
							model.MonsterTypePopularity,
						},
					},
				},
				SubType: model.MonsterSubTypeStage1,
			},
			{
				MasterCardID: model.NewMasterCardID("big-company-cto"),
				Name:         "大企業CTO",
				Rarity:       4,
				CardType:     model.CardTypeMonster,
				HP:           180,
				ImageURL:     "https://pokepoke.kurichi.dev/images/big-cto.avif",

				Type:        model.MonsterTypePopularity,
				Weakness:    model.MonsterTypeMoney,
				RetreatCost: 3,
				EvolvesFrom: []model.MasterCardID{
					model.NewMasterCardID("venture-cto"),
				},
				Ability: &model.Ability{
					Name: "組織の奴隷",
					Text: "自分の番に1回使える。このポケモンに50ダメージを与える代わりに、自分のエネルギーゾーンから[人気]エネルギーを2個出し、このポケモンにつける。",
				},
				Skills: []*model.Skill{
					{
						Name:         "一斉攻撃",
						Text:         "このポケモンについている[人気]エネルギーの数x20ダメージ追加",
						Damage:       100,
						DamageOption: "x20",
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypePopularity,
							model.MonsterTypePopularity,
						},
					},
				},
				SubType: model.MonsterSubTypeStage2,
			},
			{
				MasterCardID: model.NewMasterCardID("freelance-engineer"),
				Name:         "フリーランスエンジニア",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           100,
				ImageURL:     "https://pokepoke.kurichi.dev/images/free-eng.avif",

				Type:        model.MonsterTypeNull,
				Weakness:    model.MonsterTypeMoney,
				RetreatCost: 2,
				Skills: []*model.Skill{
					{
						Name:   "業務委託",
						Text:   "",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("liberal-arts-engineer"),
				Name:         "文系エンジニア",
				Rarity:       3,
				CardType:     model.CardTypeMonster,
				HP:           80,
				ImageURL:     "https://pokepoke.kurichi.dev/images/bunkei-engeer.avif",

				Type:        model.MonsterTypeNull,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 1,
				Skills: []*model.Skill{
					{
						Name:   "コーディング",
						Text:   "",
						Damage: 80,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
							model.MonsterTypeNull,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("ses-engineer"),
				Name:         "SESエンジニア",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           60,
				ImageURL:     "https://pokepoke.kurichi.dev/images/ses-eng.avif",

				Type:        model.MonsterTypeNull,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 3,
				Skills: []*model.Skill{
					{
						Name:   "秘密の業務",
						Text:   "自分の山札を1枚引く。",
						Damage: 30,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("zinzinia"),
				Name:         "ジンジニア",
				Rarity:       2,
				CardType:     model.CardTypeMonster,
				HP:           120,
				ImageURL:     "https://pokepoke.kurichi.dev/images/jinjinia.avif",

				Type:        model.MonsterTypeNull,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 2,
				Skills: []*model.Skill{
					{
						Name:   "人事面接",
						Text:   "自分の山札からポケモンをランダムに1枚、手札に加える。",
						Damage: 80,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
							model.MonsterTypeNull,
							model.MonsterTypeNull,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("gopher"),
				Name:         "Gopher",
				Rarity:       5,
				CardType:     model.CardTypeMonster,
				HP:           100,
				ImageURL:     "https://pokepoke.kurichi.dev/images/gopher.avif",

				Type:        model.MonsterTypeAlchohol,
				Weakness:    model.MonsterTypeKnowledge,
				RetreatCost: 2,
				Ability: &model.Ability{
					Name: "Goroutineの追撃",
					Text: "このポケモンが、相手のバトルポケモンにワザを使ったとき、ウラが出るまでコインを投げ、オモテの数x10ダメージ追加",
				},
				Skills: []*model.Skill{
					{
						Name:   "ビンタ",
						Text:   "",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypeAlchohol,
							model.MonsterTypeAlchohol,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("rustacean"),
				Name:         "Rustacean",
				Rarity:       5,
				CardType:     model.CardTypeMonster,
				HP:           100,
				ImageURL:     "https://pokepoke.kurichi.dev/images/Ferris.avif",
				Type:         model.MonsterTypeMuscle,
				Weakness:     model.MonsterTypeAlchohol,
				RetreatCost:  2,
				Ability: &model.Ability{
					Name: "所有権の濫用",
					Text: "このポケモンがバトル場にいる限り、相手は手札からグッズを使えない。",
				},
				Skills: []*model.Skill{
					{
						Name:   "はさむ",
						Text:   "",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypeMuscle,
							model.MonsterTypeMuscle,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("typescripter"),
				Name:         "TypeScripter",
				Rarity:       5,
				CardType:     model.CardTypeMonster,
				HP:           100,
				ImageURL:     "https://pokepoke.kurichi.dev/images/typescript.avif",

				Type:        model.MonsterTypeMoney,
				Weakness:    model.MonsterTypeMuscle,
				RetreatCost: 2,
				Ability: &model.Ability{
					Name: "as unknown as",
					Text: "このポケモンがいる限り、相手のポケモンのワザのタイプをNULLにし、追加効果を無効化する。",
				},
				Skills: []*model.Skill{
					{
						Name:   "as any",
						Text:   "",
						Damage: 40,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("swifter"),
				Name:         "Swifter",
				Rarity:       5,
				CardType:     model.CardTypeMonster,
				HP:           100,
				ImageURL:     "https://pokepoke.kurichi.dev/images/swift.webp",

				Type:        model.MonsterTypePopularity,
				Weakness:    model.MonsterTypeMoney,
				RetreatCost: 2,
				Ability: &model.Ability{
					Name: "Apple Developer Program",
					Text: "このポケモンがいる限り、相手は手札からポケモンを出せない。また、このポケモンは自分の番の終了時に、40ダメージを受ける。",
				},
				Skills: []*model.Skill{
					{
						Name:   "つばめがえし",
						Text:   "",
						Damage: 50,
						Cost: []model.MonsterType{
							model.MonsterTypePopularity,
							model.MonsterTypePopularity,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("pythonista"),
				Name:         "Pythonista",
				Rarity:       5,
				CardType:     model.CardTypeMonster,
				HP:           100,
				ImageURL:     "https://pokepoke.kurichi.dev/images/python.avif",
				Type:         model.MonsterTypeKnowledge,
				Weakness:     model.MonsterTypePopularity,
				RetreatCost:  2,
				Ability: &model.Ability{
					Name: "破壊的変更",
					Text: "このポケモンがいる限り、相手がポケモンを進化させた時にそのポケモンに40ダメージを与える。",
				},
				Skills: []*model.Skill{
					{
						Name:   "まきつく",
						Text:   "",
						Damage: 40,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeKnowledge,
						},
					},
				},
				SubType: model.MonsterSubTypeBasic,
			},
			{
				MasterCardID: model.NewMasterCardID("oreilly-book"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/oreilly.avif",
				Name:         "オライリー本",
				Text:         "この番、自分の[知識]ポケモンが使うワザの、相手のバトルポケモンへのダメージを+40する",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("protein"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/protain.avif",
				Name:         "プロテイン",
				Text:         "自分のエネルギーゾーンからバトル場のポケモンに[筋肉]エネルギーを2つける",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("credit-card"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/crebitcard.avif",
				Name:         "クレカ",
				Text:         "自分の[金]ポケモン1匹に[金]エネルギーを5つつける。",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("hackz-parker"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/hackz-parcker.avif",
				Name:         "ハックツパーカー",
				Text:         "この番と次の相手の番、自分の[人気]ポケモン1匹は、ワザの追加効果や特性によるダメージを受けない。",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("sake-bottle"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/issyoubin.avif",
				Name:         "一升瓶",
				Text:         "自分のバトルポケモンに自分のエネルギーゾーンから[酒]エネルギーを1つつける。",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("energy-drink"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/energy-drink.avif",
				Name:         "エナジードリンク",
				Text:         "次の相手のターンに受けるダメージを全て無効にし、その分のダメージを次の自分の番の終わりに受ける。",
				Rarity:       5,
			},
			{
				MasterCardID: model.NewMasterCardID("starbucks"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/starbacks.avif",
				Name:         "スタバ",
				Text:         "自分のポケモン1匹のHPを20回復",
				Rarity:       2,
			},
			{

				MasterCardID: model.NewMasterCardID("gopher-doll"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/gophers.avif",
				Name:         "Gopherくん人形",
				Text:         "この番、自分のバトルポケモンのにげるためのエネルギーを、1個少なくする。",
				Rarity:       2,
			},
			{
				MasterCardID: model.NewMasterCardID("hot-reload"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/hotreload.avif",
				Name:         "ホットリロード",
				Text:         "自分の手札をすべて山札に戻し、山札から同じ枚数のカードを引く。",
				Rarity:       2,
			},
			{
				MasterCardID: model.NewMasterCardID("recruitment-agency"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/e-jent.avif",
				Name:         "転職エージェント",
				Text:         "自分の山札からたねポケモン以外のポケモンをランダムに1枚、手札に加える。",
				Rarity:       2,
			},
			{
				MasterCardID: model.NewMasterCardID("programming-school"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/programing-school.avif",
				Name:         "プログラミングスクール",
				Text:         "自分の山札からたねポケモンをランダムに1枚、手札に加える。",
				Rarity:       2,
			},
			{
				MasterCardID: model.NewMasterCardID("lan-cable"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/lan.avif",
				Name:         "LANケーブル",
				Text:         "自分のベンチポケモン1匹を選び、そのポケモンについているエネルギーを1つバトルポケモンに付け替える。",
				Rarity:       2,
			},
			{
				MasterCardID: model.NewMasterCardID("hhkb"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/hhkb.avif",
				Name:         "HHKB",
				Text:         "自分の山札から「駆け出しエンジニア」の進化先のポケモンをランダムに1枚、手札に加える。",
				Rarity:       2,
			},
			{
				MasterCardID: model.NewMasterCardID("macbook"),
				CardType:     model.CardTypeGoods,
				ImageURL:     "https://pokepoke.kurichi.dev/images/macbook.avif",
				Name:         "MacBook",
				Text:         "「駆け出しエンジニア」とその進化先のポケモンが使うワザの、相手のバトルポケモンへのダメージを+20する",
				Rarity:       2,
			},
			{
				MasterCardID: model.NewMasterCardID("refactoring"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/refactor.avif",
				Name:         "リファクタリング",
				Text:         "相手のベンチポケモン1体を選び、バトル場のポケモンと入れ替える。",
				Rarity:       5,
			},
			{
				MasterCardID: model.NewMasterCardID("chat-gpt"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/ChatGPT.webp",
				Name:         "ChatGPT",
				Text:         "この番、自分のバトルポケモンのにげるためのエネルギーを、2個少なくする。",
				Rarity:       3,
			},
			{
				MasterCardID: model.NewMasterCardID("spaghetti-code"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/spagettexi.avif",
				Name:         "スパゲッティコード",
				Text:         "相手のバトルポケモンのランダムなエネルギー1個を、ランダムなエネルギーに変える。",
				Rarity:       4,
			},
			{
				MasterCardID: model.NewMasterCardID("flaming-project"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/enjou.avif",
				Name:         "炎上プロジェクト",
				Text:         "お互いのバトルポケモンについているエネルギーを1つずつトラッシュする。",
				Rarity:       3,
			},
			{
				MasterCardID: model.NewMasterCardID("security-soft"),
				CardType:     model.CardTypeSupporter,
				Name:         "セキュリティソフト",
				ImageURL:     "https://pokepoke.kurichi.dev/images/secure-soft2.avif",
				Text:         "相手の手札を全て山札に戻す。相手は相手自身の勝つために必要なポイントの数分山札を引く。",
				Rarity:       3,
			},
			{
				MasterCardID: model.NewMasterCardID("strict-mode"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/react.avif",
				Name:         "React.StrictMode",
				Text:         "自分の山札を2枚引く。",
				Rarity:       3,
			},
			{
				MasterCardID: model.NewMasterCardID("firewall"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/firewall.avif",
				Name:         "ファイヤーウォール",
				Text:         "次の相手の番、自分のポケモン全員が、相手のポケモンから受けるダメージを-20する。",
				Rarity:       3,
			},
			{
				MasterCardID: model.NewMasterCardID("bartender"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/bar.avif",
				Name:         "バーテンダー",
				Text:         "5回コインを投げ、オモテの数ぶんの好きなエネルギーを自分のエネルギーゾーンから[酒]エネルギーを出し、バトル場のポケモンにつける。ウラの数1つにつき、そのポケモンに20ダメージ与える。きぜつする場合はHPを10にする。",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("izakaya-taisyo"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/taisyo.avif",
				Name:         "居酒屋大将",
				Text:         "コイントスを裏が出るまで行い、表の数分自分についていないエネルギーの種類の中からランダムで1種類を選び、バトル場のポケモンにつける。",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("personal-trainer"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/personal-trainer.avif",
				Name:         "パーソナルトレーナー",
				Text:         "自分のバトルポケモンに[筋肉]エネルギーを2つつける。",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("yamikin-gyosya"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/yamikin.avif",
				Name:         "ヤミ金業者",
				Text:         "自分の[金]ポケモンを1匹選ぶ。そのポケモンに[金]エネルギーを10個つけ、HPを10にする。",
				Rarity:       6,
			},
			{
				MasterCardID: model.NewMasterCardID("librarian"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/tosyokan.avif",
				Name:         "図書館司書",
				Text:         "自分の山札から本のグッズカードを2枚まで手札に加える。",
				Rarity:       3,
			},
			{
				MasterCardID: model.NewMasterCardID("ci-cd-pipeline"),
				CardType:     model.CardTypeSupporter,
				ImageURL:     "https://pokepoke.kurichi.dev/images/CICD.webp",
				Name:         "CI/CDパイプライン",
				Text:         "自分の手札全てを山札に戻し、同じ枚数のカードを引く。同時に、バトル場にいる自分の[知識]ポケモン1体は、次のターンまで受けるダメージが-20される。",
				Rarity:       3,
			},
			{
				MasterCardID: model.NewMasterCardID("release-party"),
				CardType:     model.CardTypeSupporter,
				Name:         "リリースパーティ",
				ImageURL:     "https://pokepoke.kurichi.dev/images/riripa.avif",
				Text:         "自分のバトルポケモン全体のHPを20回復し、次の自分のターンまで相手のワザによるダメージを20減少させる。",
				Rarity:       3,
			},
			{
				MasterCardID: model.NewMasterCardID("wall"),
				CardType:     model.CardTypeMonster,
				Name:         "壁",
				Text:         "",
				HP:           100,
				ImageURL:     "https://pokepoke.kurichi.dev/images/wall.avif",
				Rarity:       1,
				Type:         model.MonsterTypeNull,
				SubType:      model.MonsterSubTypeBasic,
				RetreatCost:  3,
				Skills: []*model.Skill{
					{
						Name:   "壁殴り",
						Text:   "",
						Damage: 40,
						Cost: []model.MonsterType{
							model.MonsterTypeNull,
							model.MonsterTypeNull,
							model.MonsterTypeNull,
						},
					},
				},
			},
		}

		tbl := dc.Table("MasterCards")

		for _, user := range users {
			if err := tbl.Put(user).Run(cmd.Context()); err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("inserted")

		// byts, err := json.Marshal(users)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// // ファイルに書き込み
		// if err := os.WriteFile("master_cards.json", byts, 0666); err != nil {
		// 	log.Fatal(err)
		// }
	},
}

func init() {
	// rootCmd.Flags().StringP("dynamo-endpoint", "dn", "http://localhost:8000", "dynamo db endpoint")
	rootCmd.AddCommand(insertCmd)
}
