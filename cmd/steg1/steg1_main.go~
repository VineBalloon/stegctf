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
	flag1  = "FLAG{yIng_MaIN_btw}"
	tmp    *image.RGBA
)

func init() {
	flag.StringVar(&decode, "decode", "steg1.png", "use flag to specify the image to decode")
	tmp = image.NewRGBA(image.Rect(0, 0, 888, 888))
	for y := tmp.Bounds().Min.Y; y < tmp.Bounds().Max.Y; y++ {
		for x := tmp.Bounds().Min.X; x < tmp.Bounds().Max.X; x++ {
			tmp.SetRGBA(x, y, stegctf.White)
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

		img = stegctf.Steg1Encode(tmp, flag1)

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
	img, _, _ = image.Decode(file)
	fmt.Println("Flag is: " + stegctf.Steg1Decode(tmp, img))
}
