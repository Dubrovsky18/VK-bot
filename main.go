package main

import (
	"github.com/Dubrovsky18/VK-bot/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.TelBot) // введите ваш токен
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // игнорируем все, кроме сообщений
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyMarkup = getMainKeyboard() // устанавливаем клавиатуру

		switch update.Message.Text {
		case "/start":
			msg.Text = "Привет! Я бот настроения. Выбери свое настроение:"

		case "Злость":
			msg.Text = "Выбери цвет:"
			msg.ReplyMarkup = getAngryKeyboard()

		case "Радость":
			msg.Text = "Выбери цвет:"
			msg.ReplyMarkup = getHappyKeyboard()

		case "Спокойствие":
			msg.Text = "Выбери цвет:"
			msg.ReplyMarkup = getCalmKeyboard()

		case "Грусть":
			msg.Text = "Выбери цвет:"
			msg.ReplyMarkup = getSadKeyboard()

		default:
			msg.Text = "Я не знаю, что ответить на это :("
		}

		bot.Send(msg)
	}
}

// функции, возвращающие клавиатуры
func getMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.KeyboardButton{
					Text: "Злость",
				},
				tgbotapi.KeyboardButton{
					Text: "Радость",
				},
			},
			{
				tgbotapi.KeyboardButton{
					Text: "Спокойствие",
				},
				tgbotapi.KeyboardButton{
					Text: "Грусть",
				},
			},
		},
		ResizeKeyboard: true,
	}
}

func getAngryKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.KeyboardButton{
					Text: "Красный",
				},
				tgbotapi.KeyboardButton{
					Text: "Пурпурный",
				},
			},
		},
		ResizeKeyboard: true,
	}
}

func getHappyKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.KeyboardButton{
					Text: "Желтый",
				},
				tgbotapi.KeyboardButton{
					Text: "Оранжевый",
				},
			},
		},
		ResizeKeyboard: true,
	}
}

func getCalmKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.KeyboardButton{
					Text: "Светло-зеленый",
				},
				tgbotapi.KeyboardButton{
					Text: "Голубой",
				},
			},
		},
		ResizeKeyboard: true,
	}
}

func getSadKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.KeyboardButton{
					Text: "Синий",
				},
				tgbotapi.KeyboardButton{
					Text: "Фиолетовый",
				},
			},
		},
		ResizeKeyboard: true,
	}
}
