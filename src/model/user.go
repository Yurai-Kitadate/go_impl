package model

type User struct {
	UserID   int64
	UserName string
}
type Transaction struct {
	TransactionID int64
	LenderID      int64
	BorrowerID    int64
	Yen           int64
	Description   string
	IsDone        bool
	IsAccepted    bool
}
