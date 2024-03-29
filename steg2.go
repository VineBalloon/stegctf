package stegctf

import (
	"image"
	"image/color"
)

// Steg2Encode generates the steg2 challenge image
func Steg2Encode(src image.Image, in string) image.Image {
	flag := []byte(in)
	img := src.(*image.RGBA)
	var ctr = 0
	var done = false
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if !done && y%2 == 1 && x%2 == 1 {
				enc := flag[ctr]
				dig1 := enc % 10
				enc /= 10
				dig2 := enc % 10
				enc /= 10
				dig3 := enc % 10

				_, _, _, a := img.At(x, y).RGBA()
				img.SetRGBA(x, y, color.RGBA{dig1, dig2, dig3, uint8(a)})
				ctr++
				if ctr == len(flag) {
					done = true
				}
			}
		}
	}
	return img
}

// Steg2Decode decodes the image assuming it was encoded using Steg2
func Steg2Decode(src, img image.Image) string {
	res := []byte{}
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if y%2 == 1 && x%2 == 1 && !(img.At(x, y) == src.At(x, y)) {
				pix1, pix2, pix3, _ := img.At(x, y).RGBA()
				res = append(res, byte(pix1+pix2*10+pix3*100))
			}
		}
	}

	return string(res)
}
