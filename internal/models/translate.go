package models

type Translate struct {
	ID int    `db:"id"`
	RU string `db:"ru"`
	EN string `db:"en"`
}
