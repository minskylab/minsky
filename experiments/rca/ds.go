package rca

// Space have the task to describe the lattice space of a any dynamical system.
type Space interface {
	State(i ...int64) uint64
	Neighbours(i ...int64) []uint64
}

// Evolvable saves how the space evolve over time.
type Evolvable interface {
	Evolve(space Space) Space
}

// DynamicalSystem represents a generalized dynamical system.
type DynamicalSystem struct {
	state Space
	rule  Evolvable

	ticks   uint64
	running bool
}

// type Renderable interface {
// 	RenderFrame(frame )
// }

// BulkDynamicalSystem bulkanize a new DS.
func BulkDynamicalSystem(s Space, r Evolvable) *DynamicalSystem {
	return &DynamicalSystem{
		state:   s,
		rule:    r,
		ticks:   0,
		running: false,
	}
}

// State returns the state propierty of the DS structure.
func (ds *DynamicalSystem) State() Space {
	return ds.state
}

// type ComplexSystem interface {
// 	Run(*DynamicalSystem)
// }
