package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products\n"+
			"/get -  get a entity\n"+
			"/delete - delete an existing entity\n"+
			"/clear - clear chat",
	)

	c.bot.Send(msg)
}
