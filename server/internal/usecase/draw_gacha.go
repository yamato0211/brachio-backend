package usecase

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sort"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type rarityDistribution map[int]float64

var (
	// 1～3枚目はレアリティ1～4のみ出現
	dist1To3 rarityDistribution = rarityDistribution{
		1: 0.38,  // レアリティ1: 38%
		2: 0.28,  // レアリティ2: 28%
		3: 0.15,  // レアリティ3: 15%
		4: 0.09,  // レアリティ4: 9%
		5: 0.05,  // レアリティ5: 5%
		6: 0.03,  // レアリティ6: 3%
		7: 0.015, // レアリティ7: 1.5%
		8: 0.005, // レアリティ8: 0.5%
	}

	// 4枚目はレアリティ1～8が対象（レアリティ5～8は非常に低い確率）
	dist4 rarityDistribution = rarityDistribution{
		1: 0.28,  // レアリティ1: 28%
		2: 0.24,  // レアリティ2: 24%
		3: 0.20,  // レアリティ3: 20%
		4: 0.12,  // レアリティ4: 12%
		5: 0.08,  // レアリティ5: 8%
		6: 0.06,  // レアリティ6: 6%
		7: 0.015, // レアリティ7: 1.5%
		8: 0.005, // レアリティ8: 0.5%
	}

	// 5枚目は、4枚目に比べてレアリティ5～8の排出確率が若干アップ
	dist5 rarityDistribution = rarityDistribution{
		1: 0.24, // レアリティ1: 24%
		2: 0.20, // レアリティ2: 20%
		3: 0.2,  // レアリティ3: 20%
		4: 0.2,  // レアリティ4: 20%
		5: 0.07, // レアリティ5: 7%
		6: 0.05, // レアリティ6: 5%
		7: 0.03, // レアリティ7: 3%
		8: 0.01, // レアリティ8: 1%
	}
)

const (
	// 1パックのカード枚数
	packCardCount = 5
	// パックパワー
	packPower = "pack-power"
	// 必要なパックパワー
	requiredPackPower = 10
)

type DrawGachaInputPort interface {
	Execute(ctx context.Context, input *DrawGachaInput) ([]*model.MasterCard, error)
}

type DrawGachaInput struct {
	IsTen  bool
	UserID string
}

type DrawGachaInteractor struct {
	MasterCardRepository repository.MasterCardRepository
	UserRepository       repository.UserRepository
	MasterItemRepository repository.MasterItemRepository
}

func NewDrawGachaUsecase(mcr repository.MasterCardRepository, ur repository.UserRepository, mir repository.MasterItemRepository) DrawGachaInputPort {
	return &DrawGachaInteractor{
		MasterCardRepository: mcr,
		UserRepository:       ur,
		MasterItemRepository: mir,
	}
}

func (d *DrawGachaInteractor) draw(cards []*model.MasterCard) ([]*model.MasterCard, error) {
	drawn := make([]*model.MasterCard, 0, packCardCount)
	for i := range packCardCount {
		slot := i + 1
		var dist rarityDistribution
		switch slot {
		case 1, 2, 3:
			dist = dist1To3
		case 4:
			dist = dist4
		case 5:
			dist = dist5
		default:
			return nil, fmt.Errorf("invalid card slot: %d", slot)
		}

		rarity := selectRarity(dist)
		// マスターカード一覧から該当のレアリティのカードを抽出
		available := lo.Filter(cards, func(card *model.MasterCard, _ int) bool {
			return card.Rarity == rarity
		})
		var selected *model.MasterCard
		if len(available) == 0 {
			selected = &model.MasterCard{
				MasterCardID: "dummy",
				Name:         "ダミーカード",
				Rarity:       rarity,
			}
		} else {
			// available の中からランダムに1枚選択
			selected = lo.Sample(available)
		}
		drawn = append(drawn, selected)
	}
	return drawn, nil
}

func selectRarity(dist rarityDistribution) int {
	keys := lo.Keys(dist)
	sort.Ints(keys)
	r := rand.Float64()
	var cumulative float64
	for _, rarity := range keys {
		cumulative += dist[rarity]
		if r <= cumulative {
			return rarity
		}
	}
	// もしループで返されなかった場合は、最も高いレアリティを返す
	return keys[len(keys)-1]
}

func (d *DrawGachaInteractor) Execute(ctx context.Context, input *DrawGachaInput) ([]*model.MasterCard, error) {
	// パックパワー数を確認
	uid, err := model.ParseUserID(input.UserID)
	if err != nil {
		return nil, err
	}

	user, err := d.UserRepository.Find(ctx, uid)
	if err != nil {
		return nil, err
	}

	count := lo.Ternary(input.IsTen, 10, 1)

	fmt.Println(user.ItemIDsWithCount)
	userPackPower := user.ItemIDsWithCount[packPower]

	fmt.Println(userPackPower)

	// パックパワーが足りない場合
	if userPackPower < requiredPackPower*count {
		fmt.Println("パックパワーが足りません")
		return nil, model.ErrNoEnoughPackPower
	}

	masterCards, err := d.MasterCardRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	drawnCards := make([]*model.MasterCard, 0, count*packCardCount)
	for i := 0; i < count; i++ {
		drawn, err := d.draw(masterCards)
		if err != nil {
			return nil, err
		}
		drawnCards = append(drawnCards, drawn...)
	}

	return drawnCards, nil
}
