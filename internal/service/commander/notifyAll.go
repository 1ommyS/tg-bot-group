package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) NotifyAll(message *tgbotapi.Message) {
	chatId := message.Chat.ID

	chatConfig := tgbotapi.ChatConfig{ChatID: chatId}

	administratorsConfig := tgbotapi.ChatAdministratorsConfig{ChatConfig: chatConfig}

	chatAdmins, err := c.bot.GetChatAdministrators(administratorsConfig)

	if err != nil {
		log.Fatalln("Can't get chat administrators", err)
		return
	}

	messageText := "Уведомление от создателя:\n"

	for _, admin := range chatAdmins {
		messageText += "@" + admin.User.UserName + " "
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	c.bot.Send(msg)
}
