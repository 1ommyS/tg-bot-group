package main

import (
	"github.com/FilinItPark/tg-bot-notifyer/config"
	service "github.com/FilinItPark/tg-bot-notifyer/internal/service"
	"github.com/FilinItPark/tg-bot-notifyer/internal/service/commander"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func validateTokensAndGetActive() string {
	validateTokensAndChatId()

	var activeToken string

	if config.DEV_TOKEN != "" {
		activeToken = config.DEV_TOKEN
	} else {
		activeToken = config.PROD_TOKEN
	}
	return activeToken
}

func validateTokensAndChatId() {
	/*if config.PROD_CHAT_ID == "" && config.DEV_CHAT_ID == "" {
		log.Fatalf("You need to set PROD_CHAT_ID or DEV_CHAT_ID")

		panic("Ошибка запуска")
	}
	*/
	if config.DEV_TOKEN == "" && config.PROD_TOKEN == "" {
		log.Fatalf("You need to set DEV_TOKEN or PROD_TOKEN")

		panic("Ошибка запуска")
	}
}

func main() {
	activeToken := validateTokensAndGetActive()

	bot, err := tgbotapi.NewBotAPI(activeToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	serviceInstance := service.New()
	commanderInstance := commander.New(
		bot,
		serviceInstance,
	)

	cronManager := cron.New(cron.WithLocation(time.Local))

	_, err = cronManager.AddFunc("* * * * *", func() {
		birthdays := serviceInstance.CheckBirthdays()
		for _, birthday := range birthdays {
			msg := tgbotapi.NewMessage(config.DEV_CHAT_ID, birthday)

			bot.Send(msg)
		}
	})

	if err != nil {
		log.Panic(err)
	}

	cronManager.Start()

	for update := range updates {
		commanderInstance.HandleUpdate(update)
	}
}
