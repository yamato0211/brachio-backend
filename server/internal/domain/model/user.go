package model

type UserID string

func ParseUserID(s string) (UserID, error) {
	return UserID(s), nil
}

func (u UserID) String() string {
	return string(u)
}

type User struct {
	ID       UserID `dynamo:"UserId,hash"`
	Name     string `dynamo:"Name"`
	ImageURL string `dynamo:"ImageUrl"`

	// ユーザーが持っているカード
	CardIDsWithCount []*MasterCardIDWithCount `dynamo:"cardWithCount"`
	CardsWithCount   []*MasterCardWithCount   `dynamo:"-"`

	// ユーザーが持っているアイテム
	ItemIDsWithCount []*MasterItemIDWithCount `dynamo:"itemWithCount"`
	ItemsWithCount   []*MasterItemWithCount   `dynamo:"-"`
}

type MasterCardIDWithCount struct {
	MasterCardID MasterCardID `dynamo:"MasterCardId"`
	Count        int          `dynamo:"Count"`
}

type MasterCardWithCount struct {
	MasterCard *MasterCard `dynamo:"-"`
	Count      int         `dynamo:"Count"`
}

type MasterItemIDWithCount struct {
	MasterItemID MasterItemID `dynamo:"MasterItemId"`
	Count        int          `dynamo:"Count"`
}

type MasterItemWithCount struct {
	MasterItem *MasterItem `dynamo:"-"`
	Count      int         `dynamo:"Count"`
}
