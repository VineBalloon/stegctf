package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/VineBalloon/stegctf"
)

var (
	decode string
	flag2  = "FLAG{baK_1n_blAk}"
)

func init() {
	flag.StringVar(&decode, "decode", "steg2.png", "use flag to specify the image to decode")
}

func main() {
	var img image.Image
	var err error
	var file *os.File

	file, err = os.Open(decode)
	if err != nil {
		fmt.Println("No steg file detected, generating one...")

		img = stegctf.Steg2Encode(image.NewRGBA(image.Rect(0, 0, 888, 888)), flag2)
		file, err := os.Create("steg2.png")
		if err != nil {
			return
		}
		defer file.Close()
		png.Encode(file, img)
		fmt.Println("Steg file created!")
		return
	}
	defer file.Close()

	fmt.Println("Attempting to decode using steg2 encoding scheme...")
	img, _, err = image.Decode(file)
	fmt.Println("Flag is: " + stegctf.Steg2Decode(img))
}
