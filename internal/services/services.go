package services

import (
	"github.com/plin2k/ggore-tg-channels/internal/models"
	"github.com/plin2k/ggore-tg-channels/internal/services/parser"
	"github.com/plin2k/ggore-tg-channels/internal/services/telegram"
)

type Parser interface {
	Parse(source models.Source) (models.Contents, error)
}

type Telegram interface {
	Send(channelID int64, message string) error
}

type Service struct {
	Parser
	Telegram
}

func NewService(cfg *models.Config) (*Service, error) {
	tg, err := telegram.New(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}

	return &Service{
		Parser:   parser.New(),
		Telegram: tg,
	}, nil
}
