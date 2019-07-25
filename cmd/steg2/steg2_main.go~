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
	tmp    *image.RGBA
)

func init() {
	flag.StringVar(&decode, "decode", "steg2.png", "use flag to specify the image to decode")
	tmp = image.NewRGBA(image.Rect(0, 0, 888, 888))
	for y := tmp.Bounds().Min.Y; y < tmp.Bounds().Max.Y; y++ {
		for x := tmp.Bounds().Min.X; x < tmp.Bounds().Max.X; x++ {
			tmp.SetRGBA(x, y, stegctf.Black)
		}
	}
}

func main() {
	var img image.Image
	var err error
	var file *os.File

	file, err = os.Open(decode)
	if err != nil {
		fmt.Println("No steg file detected, generating one...")

		img = stegctf.Steg2Encode(tmp, flag2)
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
	img, _, _ = image.Decode(file)
	fmt.Println("Flag is: " + stegctf.Steg2Decode(tmp, img))
}
