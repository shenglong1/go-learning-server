package model

type UserModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
