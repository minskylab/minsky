package rca

// Tick execute one tick dynamical system evolution.
func (ds *DynamicalSystem) Tick() {
	ds.state = ds.rule.Evolve(ds.state)
	ds.ticks++
}

// RunSimulation run simulation 'n' ticks.
func (ds *DynamicalSystem) RunSimulation(n uint64, cn chan uint64) {
	ds.running = true
	go func(cn chan uint64, running *bool) {
		for i := uint64(0); i < n; i++ {
			ds.Tick()
			cn <- ds.ticks
		}
		close(cn)
		*running = false
	}(cn, &ds.running)
}

// RunInfiniteSimulation runs a infinite (but closable) simulation.
func (ds *DynamicalSystem) RunInfiniteSimulation(cn chan uint64, finish chan struct{}) {
	ds.running = true

	go func(f chan struct{}, running *bool) {
		<-finish
		*running = false
	}(finish, &ds.running)

	go func(cn chan uint64, running *bool) {
		cn <- ds.ticks
		for *running {
			ds.Tick()
			cn <- ds.ticks
		}
		close(cn)
		close(finish)
	}(cn, &ds.running)
}

// Observe execute a function on every tick from ticker channel.
func (ds *DynamicalSystem) Observe(cn chan uint64, evol RenderFunction) {
	for n := range cn {
		evol(n, ds.state)
	}
}
