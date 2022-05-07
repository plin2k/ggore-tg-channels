package repository

import (
	"database/sql"

	"github.com/plin2k/ggore-tg-channels/internal/models"
)

type Article interface {
	Create(article models.Article) error
	Exists(article models.Article) (bool, error)
}

type Translate interface {
	Create(post models.Translate) error
	GetTranslateFromRU(str string) (string, error)
}

type Repository struct {
	Article
	Translate
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Article:   NewArticle(db),
		Translate: NewTranslate(db),
	}
}
