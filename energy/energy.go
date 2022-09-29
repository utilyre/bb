package energy

import "math"

type Energy interface {
	// Returns falling status.
	IsFalling() bool

	// Returns potential energy.
	Potential() float64

	// Returns kinetic energy.
	Kinetic() float64

	// Returns mechanical energy.
	Mechanical() float64

	// Sets potential energy.
	// Has a side effect on falling status and kinetic energy.
	SetPotential(potential float64)

	// Exerts force on object.
	// Has a side effect on falling status and kinetic energy.
	ExertForce(mass, force, time float64)
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
	if e.Potential() == 0 {
		e.kinetic = 0.8 * e.Kinetic()
	}
}

func (e *energy) ExertForce(mass, force, time float64) {
	v0 := math.Sqrt(2 * e.Kinetic() / mass) // V = √(2K/m)
	dv := math.Abs(force) * time / mass     // ΔV = fΔt/m

	coefficient := 1.0
	if (e.IsFalling() && force > 0) || (!e.IsFalling() && force < 0) {
		coefficient = -1.0
	}

	v := v0 + coefficient*dv
	if v < 0 {
		e.isFalling = !e.IsFalling()
	}

	e.kinetic = mass * math.Pow(v, 2) / 2
}
