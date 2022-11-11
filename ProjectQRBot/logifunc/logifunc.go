package logifunc

import (
	"log"
	"os"
)

/// функции для логирования

func CreateLogFile() {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("Start of logging... \n"); err != nil {
		log.Println(err)
	}
}

func LoggingData(logi string) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "префикс: ", log.LstdFlags)
	logger.Println(logi)

}
