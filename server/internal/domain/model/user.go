package model

type UserID string

func ParseUserID(s string) (UserID, error) {
	return UserID(s), nil
}

func (u UserID) String() string {
	return string(u)
}

type User struct {
	ID      UserID `dynamo:"UserId,hash"`
	Name    string `dynamo:"Name"`
	Picture string `dynamo:"Picture"`
}
