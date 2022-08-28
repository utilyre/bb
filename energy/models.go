package energy

type energy struct {
	isFalling bool
	potential float64
	speed     float64
}

func (e *energy) IsFalling() bool {
	return e.isFalling
}

func (e *energy) Potential() float64 {
	return e.potential
}

func (e *energy) Speed() float64 {
	return e.speed
}

func (e *energy) Mechanical() float64 {
	return e.potential + e.speed
}

func (e2 *energy) SetPotential(potential float64) {
	e1 := *e2

	e2.potential = potential
	e2.speed = e1.Mechanical() - e2.Potential() // V₂ = E₁ - U₂

	if e2.Potential() <= 0 {
		e2.isFalling = false
	}
	if e2.Speed() <= 0 {
		e2.isFalling = true
	}

	if e1.IsFalling() && !e2.IsFalling() {
		e2.speed -= e2.speed * 0.4
	}
}
