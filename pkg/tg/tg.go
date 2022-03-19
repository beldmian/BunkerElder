package tg

import (
	"github.com/beldmian/bunkerelder/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type TgBot struct {
	api *tgbotapi.BotAPI	
	l *zap.Logger
}

func ProvideTg(conf *config.Config, l *zap.Logger) *TgBot {
	l.Info("creating bot", zap.String("token", conf.TelegramAPIToken))
	bot, err := tgbotapi.NewBotAPI(conf.TelegramAPIToken)
	if err != nil {
		l.Fatal("Error occured while creating telegram bot", zap.Error(err))
	}
	tgbot := TgBot{
		api: bot,
		l: l,
	}
	return &tgbot
}

func (bot TgBot) Start() {
	u := tgbotapi.NewUpdate(0)	
	u.Timeout = 60

	updates := bot.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "recieved")
			if _, err := bot.api.Send(reply); err != nil {
				bot.l.Warn("Cannot send message", zap.Error(err))			
			}
		}
	}
}
