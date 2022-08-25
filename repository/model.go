package repository

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"user_name"`
}
type Transaction struct {
	TransactionID int64  `db:"user_id"`
	LenderID      int64  `db:"lender_id"`
	BorrowerID    int64  `db:"borrower_id"`
	Yen           int64  `db:"yen"`
	Description   string `db:"description"`
	IsDone        bool   `db:"is_done"`
	IsAccepted    bool   `db:"is_accepted"`
}
