package gol

import (
	"image"

	"github.com/minskylab/rca"
)

type ImageRenderer struct {
	images chan image.Image
}

func (ir *ImageRenderer) Render(n uint64, s rca.Space) {
	ca := s.(*CA)

	flat := []byte{}
	for _, r := range ca.board {
		flat = append(flat, r...)
	}

	w := len(ca.board)
	h := len(ca.board[0])

	img := image.NewGray(image.Rect(0, 0, w, h))

	for i := range flat {
		img.Pix[i] = flat[i] * 255
	}

	ir.images <- img
}

func NewImageRenderer(images chan image.Image) *ImageRenderer {
	return &ImageRenderer{images: images}
}
