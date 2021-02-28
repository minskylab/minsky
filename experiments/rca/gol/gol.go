package gol

import (
	"math/rand"

	"github.com/minskylab/rca"
)

// CA implements a Game Of Life CA for DynamicalSystem model.
type CA struct {
	board    [][]byte
	cellSize int
}

// State implementes DS Space interface.
func (ca *CA) State(i ...uint64) uint64 {
	return uint64(ca.board[i[0]][i[1]])
}

// Neighbours implementes DS Space interface.
func (ca *CA) Neighbours(i ...uint64) []uint64 {
	return []uint64{0, 0, 0, 0, 0, 0, 0, 0}
}

// Evolve implements a Evolvable interface.
func (ca *CA) Evolve(space rca.Space) rca.Space {
	w := len(ca.board)
	h := len(ca.board[0])
	ca.board[rand.Intn(w)][rand.Intn(h)] = byte(rand.Intn(2))

	return ca
}

// new creates a new GoL system.
func new(w, h int, randomSeed int64) *CA {
	rand.Seed(randomSeed)

	board := make([][]byte, h)
	for i := range board {
		board[i] = make([]byte, w)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = byte(rand.Intn(2))
		}
	}

	return &CA{
		board:    board,
		cellSize: 4,
	}
}

// NewGoLDynamicalSystem returns a new GoL DS.
func NewGoLDynamicalSystem(w, h int, randomSeed int64) *rca.DynamicalSystem {
	gol := new(w, h, randomSeed)
	return rca.BulkDynamicalSystem(gol, gol)
}
