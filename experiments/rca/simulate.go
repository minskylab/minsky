package rca

// Tick execute one tick dynamical system evolution.
func (ds *DynamicalSystem) Tick() {
	ds.state = ds.rule.Evolve(ds.state)
	ds.ticks++
}

// RunSimulation run simulation 'n' ticks.
func (ds *DynamicalSystem) RunSimulation(n uint64, cn chan uint64) {
	go func(cn chan uint64) {
		for i := uint64(0); i < n; i++ {
			ds.Tick()
			cn <- ds.ticks
		}
		close(cn)
	}(cn)
}

// RunInfiniteSimulation runs a infinite (but closable) simulation.
func (ds *DynamicalSystem) RunInfiniteSimulation(cn chan uint64, finish chan struct{}) {
	go func(f chan struct{}) {
		<-finish
		ds.sopped = true
		close(cn)
		close(finish)
	}(finish)

	go func(cn chan uint64, stop *bool) {
		cn <- ds.ticks
		for !*stop {
			ds.Tick()
			cn <- ds.ticks
		}
	}(cn, &ds.stopped)
}

// Observe execute a function on every tick from ticker channel.
func (ds *DynamicalSystem) Observe(cn chan uint64, evol func(n uint64, s Space)) {
	for n := range cn {
		evol(n, ds.state)
	}
}
