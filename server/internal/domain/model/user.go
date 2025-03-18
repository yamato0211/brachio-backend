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
	CardIDsWithCount map[string]int `dynamo:"cardWithCount"`

	// ユーザーが持っているアイテム
	ItemIDsWithCount map[string]int `dynamo:"itemWithCount"`
}

type MasterCardWithCount struct {
	MasterCard *MasterCard `dynamo:"-"`
	Count      int         `dynamo:"Count"`
}

type MasterItemWithCount struct {
	MasterItem *MasterItem `dynamo:"-"`
	Count      int         `dynamo:"Count"`
}
