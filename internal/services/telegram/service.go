package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/pkg/errors"
)

type service struct {
	Bot *tgbotapi.BotAPI
}

func New(token string) (*service, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.Wrap(err, "telegram.NewBot")
	}

	return &service{Bot: bot}, nil
}

func (s *service) Send(channelID int64, message string) error {
	msg := tgbotapi.NewMessage(channelID, message)
	msg.ParseMode = tgbotapi.ModeHTML

	_, err := s.Bot.Send(msg)
	if err != nil {
		return errors.Errorf("telegram.SendMessage : %w", err)
	}

	return nil
}

//https://api.telegram.org/bot<TOKEN>/getUpdates
