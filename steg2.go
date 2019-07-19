package stegctf

import (
	"image"
	"image/color"
)

// Steg2Encode generates the steg1 challenge image
func Steg2Encode() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 800, 600))
	var ctr = 0
	var done = false
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if !done && x%2 == 0 {
				img.SetRGBA(x, y, color.RGBA{flag2[ctr], 0, 0, 0xff})
				ctr++
				if ctr == len(flag2) {
					done = true
				}
				continue
			}
			img.SetRGBA(x, y, black)
		}
	}
	return img
}

// Steg2Decode decodes the image assuming it was encoded using Steg1
func Steg2Decode(img image.Image) string {
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
