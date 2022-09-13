package energy

type Energy interface {
	// Returns falling status
	IsFalling() bool

	// Returns potential energy
	Potential() float64

	// Returns kinetic energy
	Kinetic() float64

	// Returns mechanical energy
	Mechanical() float64

	// Sets potential energy and has side effect on kinetic energy
	SetPotential(potential float64)
}

type energy struct {
	isFalling bool
	potential float64
	kinetic   float64
}

func NewEnergy(potential, kinetic float64) Energy {
	return &energy{
		isFalling: true,
		potential: potential,
		kinetic:   kinetic,
	}
}

func (e *energy) IsFalling() bool {
	return e.isFalling
}

func (e *energy) Potential() float64 {
	return e.potential
}

func (e *energy) Kinetic() float64 {
	return e.kinetic
}

func (e *energy) Mechanical() float64 {
	return e.Potential() + e.Kinetic()
}

func (e *energy) SetPotential(potential float64) {
	// Takes a copy of e
	e0 := *e

	e.potential = potential
	e.kinetic = e0.Mechanical() - e.Potential() // K₁ = E₀ - U₁

	// Ensures that potential and kinetic energies are unsigned values
	if e.Potential() < 0 {
		e.potential = 0
	}
	if e.Kinetic() < 0 {
		e.kinetic = 0
	}

	// Changes falling status based on the current potential and kinetic energies
	if e.Potential() == 0 {
		e.isFalling = false
	}
	if e.Kinetic() == 0 {
		e.isFalling = true
	}

	// Subtracts wasted energy from kinetic energy whenever object hits the ground
	// BEWARE: This is a workaround and is not according to physics
	if e0.IsFalling() && !e.IsFalling() {
		e.kinetic = 0.8 * e.Kinetic()
	}
}
