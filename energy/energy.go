package energy

type Energy interface {
	IsFalling() bool
	Potential() float64
	Kinetic() float64
	Mechanical() float64
	SetPotential(velocity float64)
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
	e0 := *e

	e.potential = potential
	e.kinetic = e0.Mechanical() - e.Potential() // K₁ = E₀ - U₁

	if e.Potential() < 0 {
		e.potential = 0
	}
	if e.Kinetic() < 0 {
		e.kinetic = 0
	}

	if e.Potential() == 0 {
		e.isFalling = false
	}
	if e.Kinetic() == 0 {
		e.isFalling = true
	}

	if e0.IsFalling() && !e.IsFalling() {
		e.kinetic = 0.8 * e.Kinetic()
	}
}
