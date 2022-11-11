package main

import (
	pdata "awesomeProjectQRBot/privatdata"
	bot "awesomeProjectQRBot/runBot"
)

func main() {
	bot.WorkBot(pdata.TokenKey)
}
