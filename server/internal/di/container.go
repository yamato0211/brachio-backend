package di

import "github.com/yamato0211/brachio-backend/internal/usecase"

type Container struct {
	DrawGachaInputPort      usecase.DrawGachaInputPort
	GetMasterCardsInputPort usecase.GetMasterCardsInputPort
	MatchingInputPort       usecase.MatchingInputPort
	GetMyDecksInputPort     usecase.GetMyDecksInputPort
	UseGoodsInputPort       usecase.UseGoodsInputPort
	UseSupporterInputPort   usecase.UseSupporterInputPort
}
