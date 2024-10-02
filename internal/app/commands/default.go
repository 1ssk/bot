package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote:"+inputMessage.Text)

	c.bot.Send(msg)
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue) // don't stop program! @ssk
		}
	}()

	if update.Message == nil {
		return
	}
	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
