package stegctf

import (
	"image"
	"image/color"
)

// Steg3Encode generates the steg3 challenge image
func Steg3Encode(src image.Image, in string) image.Image {
	// encode half the message in both with half transparency
	//enc1 := Steg1Encode(src, in[0:len(in)/2])
	//enc2 := Steg2Encode(src, in[len(in)/2:len(in)])
	var ctr = 0
	var done = false

	img := src.(*image.RGBA)
	flag := []byte(in)
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if !done {
				enc := flag[ctr]
				dig1 := enc
				dig2 := enc
				dig3 := enc

				img.SetRGBA(x, y, color.RGBA{dig1, dig2, dig3, 0xff})
				ctr++
				if ctr == len(flag) {
					done = true
				}
				continue
			}
			img.SetRGBA(x, y, black)
		}
	}
	return img
}

// Steg3Decode decodes the image assuming it was encoded using Steg3
func Steg3Decode(src, img image.Image) string {
	res := []byte{}
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if y%2 == 1 && x == 420 && !(img.At(x, y) == black) {
				pix1, pix2, pix3, _ := img.At(x, y).RGBA()
				res = append(res, byte(pix1+pix2*10+pix3*100))
			}
		}
	}

	return string(res)
}
