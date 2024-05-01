package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) SendHelpMessage(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/start для запуска, /add для добавления дня рождения, /edit для редактирования, /delete для удаления")

	c.bot.Send(msg)

}
