package model

type MasterItem struct {
	ItemID   string `dynamo:"ItemId,hash"`
	Name     string `dynamo:"Name"`
	ImageURL string `dynamo:"ImageUrl"`
}
