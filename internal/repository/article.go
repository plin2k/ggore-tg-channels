package repository

import (
	"database/sql"

	"github.com/plin2k/ggore-tg-channels/internal/models"
)

type ArticlesRepository struct {
	db *sql.DB
}

func NewArticle(db *sql.DB) *ArticlesRepository {
	return &ArticlesRepository{db: db}
}

func (repo *ArticlesRepository) Create(article models.Article) error {
	stmt, err := repo.db.Prepare("INSERT INTO " + articleTable + "(signature, channel_id, html_message) values(?,?,?)")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(article.Signature, article.ChannelID, article.HTMLMessage)
	if err != nil {
		return err
	}

	return err
}

func (repo *ArticlesRepository) Exists(article models.Article) (bool, error) {
	row := repo.db.QueryRow("SELECT id FROM "+articleTable+" where signature = ? LIMIT 1", article.Signature)

	var model models.Article

	if err := row.Scan(&model.ID); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
