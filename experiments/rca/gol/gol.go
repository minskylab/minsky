package gol

import (
	"math/rand"

	"github.com/minskylab/rca"
)

// GoLCA implements a Game Of Life CA for DynamicalSystem model.
type GoLCA struct {
	board [][]byte
}

// State implementes DS Space interface.
func (ca *GoLCA) State(i ...uint64) uint64 {
	return uint64(ca.board[i[0]][i[1]])
}

// Neighbours implementes DS Space interface.
func (ca *GoLCA) Neighbours(i ...uint64) []uint64 {
	return []uint64{0, 0, 0, 0, 0, 0, 0, 0}
}

// Evolve implements a Evolvable interface.
func (ca *GoLCA) Evolve(space rca.Space) rca.Space {
	w := len(ca.board)
	h := len(ca.board[0])
	ca.board[rand.Intn(w)][rand.Intn(h)] = byte(rand.Intn(2))

	return ca
}

// new creates a new GoL system.
func new(w, h int, randomSeed int64) *GoLCA {
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

	return &GoLCA{
		board: board,
	}

}

// New returns a new GoL DS.
func New(w, h int, randomSeed int64) *rca.DynamicalSystem {
	gol := new(w, h, randomSeed)
	return rca.BulkDynamicalSystem(gol, gol)
}
