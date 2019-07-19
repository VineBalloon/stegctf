package stegctf

import (
	"image"
	"image/color"
)

// Steg1Encode generates the steg1 challenge image
func Steg1Encode() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 800, 600))
	var ctr = 0
	var done = false
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if !done && x%2 == 0 {
				img.SetRGBA(x, y, color.RGBA{flag1[ctr], 0xff, 0xff, 0xff})
				ctr++
				if ctr == len(flag1) {
					done = true
				}
				continue
			}
			img.SetRGBA(x, y, white)
		}
	}
	return img
}

// Steg1Decode decodes the image assuming it was encoded using Steg1
func Steg1Decode(img image.Image) string {
	res := []byte{}
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if x%2 == 0 && !(img.At(x, y) == white) {
				pix, _, _, _ := img.At(x, y).RGBA()
				res = append(res, byte(pix))
			}
		}
	}

	return string(res)
}
