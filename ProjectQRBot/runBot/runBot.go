package runBot

import (
	qr "awesomeProjectQRBot/createQR"
	logi "awesomeProjectQRBot/logifunc"
	"github.com/Syfaro/telegram-bot-api"
	"io/ioutil"
	"log"
)

// /функции работы бота
func WorkBot(tokens string) {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI(tokens)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// cоздание файла с логами
	logi.CreateLogFile()

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
		/// обработка сообщений ( запрещаем отправлять всякое в бот )
		if update.Message.Photo != nil || update.Message.Sticker != nil || update.Message.Voice != nil || update.Message.Video != nil || update.Message.Audio != nil || update.Message.VideoNote != nil {
			reply = " не умею с таким работать"
		}
		if update.Message.VideoNote == nil && update.Message.Audio == nil && update.Message.Video == nil && update.Message.Voice == nil && update.Message.Sticker == nil && update.Message.Photo == nil && update.Message.Text != " " && update.Message.Text != "/start" {
			//createQR(update.Message.Text)
			qr.CreateQR(update.Message.Text)
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
		logi.LoggingData(string(update.Message.From.UserName + " " + update.Message.Text))
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
