package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/VineBalloon/stegctf"
)

var (
	decode string
	flag3  = "FLAG{change_comes_from_within}"
)

func main() {
	var err error
	var decode = "steg3.png"
	var origin = "yin_yang.png"
	var deco *os.File
	var orig *os.File
	var dec image.Image
	var ori image.Image

	orig, err = os.Open(origin)
	if err != nil {
		fmt.Println("Error: " + origin + " not found!")
		return
	}

	deco, err = os.Open(decode)
	if err != nil {
		fmt.Println("No steg file detected, generating one...")

		ori, err = png.Decode(orig)
		if err != nil {
			return
		}

		dec = stegctf.Steg3Encode(ori, flag3)
		deco, err := os.Create(decode)
		if err != nil {
			return
		}
		defer deco.Close()
		png.Encode(deco, dec)
		fmt.Println("Steg file created!")
		return
	}
	defer deco.Close()

	fmt.Println("Attempting to decode using steg3 encoding scheme...")
	dec, _ = png.Decode(deco)
	ori, _ = png.Decode(orig)
	fmt.Println("Flag is: " + stegctf.Steg3Decode(ori, dec))
}
