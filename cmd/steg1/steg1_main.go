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
)

func init() {
	flag.StringVar(&decode, "decode", "steg1.png", "use flag to specify the image to decode")
}

func main() {
	var img image.Image
	var err error
	var file *os.File

	file, err = os.Open(decode)
	if err != nil {
		fmt.Println("No steg file detected, generating one...")
		img = stegctf.Steg1Encode()
		file, err := os.Create("steg1.png")
		if err != nil {
			return
		}
		defer file.Close()
		png.Encode(file, img)
		fmt.Println("Steg file created!")
		return
	}
	defer file.Close()

	fmt.Println("Attempting to decode using steg1 encoding scheme...")
	img, _, err = image.Decode(file)
	fmt.Println("Flag is: " + stegctf.Steg1Decode(img))
}
