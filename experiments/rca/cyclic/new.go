package cyclic

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/lucasb-eyer/go-colorful"
)

func mustHexColor(val string) colorful.Color {
	s, err := colorful.Hex(val)
	if err != nil {
		panic(err)
	}

	return s
}

// NewRockPaperSissor ...
func NewRockPaperSissor(w, h int, threshold int, stocasticity int, randomSeed int64, images chan image.Image) (*PaperRockSissor, error) {
	rand.Seed(randomSeed)

	c1 := mustHexColor("#7376ac")
	c2 := mustHexColor("#212656")

	colorMap := map[byte]color.Color{
		0: c1,
		1: c1.BlendHcl(c2, 0.5),
		2: c2,
	}

	countsMap := map[byte]int{}

	for c := range colorMap {
		countsMap[c] = 0
	}

	board := make([][]byte, h)
	for i := range board {
		board[i] = make([]byte, w)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = byte(rand.Intn(len(colorMap)))
		}
	}

	return &PaperRockSissor{
		board:        board,
		countsMap:    countsMap,
		threshold:    threshold,
		frame:        image.NewRGBA(image.Rect(0, 0, w, h)),
		images:       images,
		colormap:     colorMap,
		stocasticity: stocasticity,
	}, nil
}
