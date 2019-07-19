package ctf

import (
	"image"
	"image/png"
	"io"

	// add decoders
	_ "image/gif"
	_ "image/jpeg"
)

// ConvertToPNG converts from any recognized format to PNG.
func ConvertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}
