package createQR

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
)

// /функция генерации QR кодов
func CreateQR(content string) {
	err := qrcode.WriteColorFile(content, qrcode.Medium, 256, color.Black, color.White, "secondfile.png")
	if err != nil {
		fmt.Printf("Sorry couldn't create qrcode:,%v", err)
	}
}
