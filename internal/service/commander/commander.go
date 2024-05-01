package commander

import (
	"github.com/FilinItPark/tg-bot-notifyer/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func New(bot *tgbotapi.BotAPI, service *service.Service) *Commander {
	return &Commander{
		bot:     bot,
		service: service,
	}
}

func (commander *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		log.Printf("%+v\n", update.Message)

		switch update.Message.Command() {
		case "start":
			commander.SendStartMessage(update.Message)
		case "help":
			commander.SendHelpMessage(update.Message)
		default:
			commander.HandleDefault(update.Message)
		}
	}
}
