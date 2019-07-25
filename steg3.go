package stegctf

import (
	"fmt"
	"image"
	"image/color"
)

// Steg3Encode generates the steg3 challenge image
func Steg3Encode(src image.Image, in string) image.Image {
	// split message
	var half1 string
	var half2 string
	for i := 0; i < len(in); i++ {
		if i%2 == 0 {
			half1 += string(in[i])
			continue
		}
		half2 += string(in[i])
	}

	url := "https://en.wikipedia.org/wiki/File:Yin_yang.svg"
	flag := half1 + half2
	img := image.NewRGBA(image.Rect(0, 0, 888, 888))
	var ctr = 0
	var urlc = 0
	var done = false
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if urlc < len(url) && x%2 == 0 && y < 444 && x < 444 {
				_, g, b, a := src.At(x, y).RGBA()
				if a == 0 {
					continue
				}
				// set red channel
				img.SetRGBA(x, y, color.RGBA{uint8(url[urlc]), uint8(g), uint8(b), uint8(a)})
				fmt.Println("(", x, y, ")", url[urlc])
				urlc++
				continue
			}

			if !done && y%8 == 0 && x%8 == 0 && y > 444 && x > 444 {
				enc := flag[ctr]
				r, g, b, a := src.At(x, y).RGBA()
				// skip transparent bits
				if a == 0 {
					continue
				}
				dig1 := (uint8(r)/10)*10 + enc%10
				enc /= 10
				dig2 := (uint8(g)/10)*10 + enc%10
				enc /= 10
				dig3 := (uint8(b)/10)*10 + enc%10

				img.SetRGBA(x, y, color.RGBA{dig1, dig2, dig3, uint8(a)})
				ctr++
				if ctr == len(flag) {
					done = true
				}
				continue
			}
			r, g, b, a := src.At(x, y).RGBA()
			img.SetRGBA(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	return img
}

// Steg3Decode decodes the image assuming it was encoded using Steg3
func Steg3Decode(src, img image.Image) string {
	urb := []byte{}
	res := []byte{}
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if x%2 == 0 && y < 444 && x < 444 && img.At(x, y) != src.At(x, y) {
				r, _, _, _ := img.At(x, y).RGBA()
				urb = append(urb, byte(uint8(r)))
			}

			if y%8 == 0 && x%8 == 0 && y > 444 && x > 444 {
				r1, g1, b1, _ := img.At(x, y).RGBA()
				r2, g2, b2, _ := src.At(x, y).RGBA()
				by := byte(((uint8(r1) - uint8(r2)) % 10) + ((uint8(g1)-uint8(g2))%10)*10 + ((uint8(b1)-uint8(b2))%10)*100)
				if by == 0 {
					continue
				}
				res = append(res, by)
			}
		}
	}
	// broken lmao
	fmt.Println("url:", string(urb))

	flag := make([]byte, len(res))
	ctr := 0
	for i := 0; i < len(res)/2; i++ {
		flag[ctr] = res[i]
		ctr += 2
	}

	ctr = 1
	for i := len(res) / 2; i < len(res); i++ {
		flag[ctr] = res[i]
		ctr += 2
	}

	return string(flag)
}
