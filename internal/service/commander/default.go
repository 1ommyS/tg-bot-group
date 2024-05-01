package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) HandleDefault(inputMessage *tgbotapi.Message) {
	if inputMessage.Text[0] == '/' {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Неизвестная команда")

		c.bot.Send(msg)
	}
}
