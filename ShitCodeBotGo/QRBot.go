package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/skip2/go-qrcode"
	"image/color"
	"io/ioutil"
	"log"
)

// /функция генерации QR кодов
func createQR(content string) {
	err := qrcode.WriteColorFile(content, qrcode.Medium, 256, color.Black, color.White, "secondfile.png")
	if err != nil {
		fmt.Printf("Sorry couldn't create qrcode:,%v", err)
	}
}
func WorkBot(tokens string) {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI(tokens)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		// универсальный ответ на любое сообщение
		reply := " "
		if update.Message == nil {
			continue
		}
		/// обработка сообщений
		if update.Message.Text != " " && update.Message.Text != "/start" {
			createQR(update.Message.Text)
			/// отправка png QR code
			photoBytes, err := ioutil.ReadFile("secondfile.png")
			if err != nil {
				panic(err)
			}
			photoFileBytes := tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: photoBytes,
			}
			chatID := update.Message.Chat.ID
			bot.Send(tgbotapi.NewPhotoUpload(int64(chatID), photoFileBytes))

		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот. Который умеет генерировать QR коды\n Отправь мне инфу и я кину тебе QR)"
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)

	}

}
func main() {
	keyBot := ""
	WorkBot(keyBot)
}
