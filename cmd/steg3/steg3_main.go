package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/VineBalloon/stegctf"
)

var (
	decode string
	flag3  = "FLAG{0_1}"
)

func main() {
	var img image.Image
	var err error
	var file *os.File
	var orig *os.File
	var decode = "steg3.png"
	var origin = "yin_yang.png"

	orig, err = os.Open(origin)
	if err != nil {
		log.Fatalln("Error: " + origin + " not found!")
		return
	}

	file, err = os.Open(decode)
	if err != nil {
		fmt.Println("No steg file detected, generating one...")

		img = stegctf.Steg3Encode(orig, flag3)
		file, err := os.Create(decode)
		if err != nil {
			return
		}
		defer file.Close()
		png.Encode(file, img)
		fmt.Println("Steg file created!")
		return
	}
	defer file.Close()

	fmt.Println("Attempting to decode using steg3 encoding scheme...")
	img, _, _ = image.Decode(file)
	ori, _, _ = image.Decode(orig)
	fmt.Println("Flag is: " + stegctf.Steg2Decode(ori, img))
}
