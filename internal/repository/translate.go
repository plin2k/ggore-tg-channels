package repository

import (
	"database/sql"

	"github.com/plin2k/ggore-tg-channels/internal/models"
)

type TranslateRepository struct {
	db *sql.DB
}

func NewTranslate(db *sql.DB) *TranslateRepository {
	return &TranslateRepository{db: db}
}

func (repo *TranslateRepository) Create(translate models.Translate) error {
	stmt, err := repo.db.Prepare("INSERT INTO " + translateTable + "(ru, en) values(?,?)")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(translate.RU, translate.EN)
	if err != nil {
		return err
	}

	return err
}

func (repo *TranslateRepository) GetTranslateFromRU(str string) (string, error) {
	row := repo.db.QueryRow("SELECT en FROM "+translateTable+" where ru = ? LIMIT 1", str)

	var result string

	if err := row.Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			return str, nil
		}

		return str, err
	}

	return result, nil
}
