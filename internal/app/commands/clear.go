package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Clear(inputMessage *tgbotapi.Message) {
	offset := 0
	for {
		// Получаем сообщения с ограничением 100, начиная с offset
		updates, err := c.bot.GetUpdatesChan(tgbotapi.UpdateConfig{
			Timeout: 1,      //  Ожидание  обновлений  в  течение  60  секунд
			Limit:   10,     //  Максимальное  количество  обновлений  за  один  запрос
			Offset:  offset, //  Смещение  для  получения  следующих  сообщений
		})

		if err != nil {
			fmt.Println("Ошибка получения истории сообщений:", err)
			return // Выходим из функции, если произошла ошибка
		}

		// Удаляем каждое сообщение в цикле
		for update := range updates {
			// Удаляем только сообщения в текущем чате
			if update.Message.Chat.ID == inputMessage.Chat.ID {
				deleteMsg := tgbotapi.NewDeleteMessage(inputMessage.Chat.ID, update.Message.MessageID)
				_, err := c.bot.Send(deleteMsg)
				if err != nil {
					fmt.Println("Ошибка удаления сообщения:", err)
					// Продолжаем цикл, даже если удаление не удалось
				}
			}
		}

		// Проверяем, есть ли еще сообщения в истории
		if len(updates) < 10 { // Если получено менее 100 сообщений, значит, история закончилась
			break
		}

		// Обновляем offset для получения следующих сообщений
		offset += 10
	}

	// Отправляем сообщение, что все сообщения были удалены
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Все сообщения были удалены",
	)

	c.bot.Send(msg)
}
