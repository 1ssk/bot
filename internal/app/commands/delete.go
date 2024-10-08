package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) {
	// Создаем сообщение для удаления предыдущего сообщения
	deleteMsg := tgbotapi.NewDeleteMessage(inputMessage.Chat.ID, inputMessage.MessageID-1)

	// Отправляем запрос на удаление
	_, err := c.bot.Send(deleteMsg)
	if err != nil {
		// Обрабатываем ошибку, если удаление не удалось
		fmt.Println("Ошибка удаления сообщения:", err)
	}

	// Отправляем сообщение, что предыдущее сообщение было удалено
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Предыдущее сообщение было удалено",
	)

	c.bot.Send(msg)
}
