package energy

import "math"

type Energy interface {
	// Returns potential energy.
	Potential() float64

	// Returns kinetic energy.
	Kinetic() float64

	// Calculates and returns velocity.
	Velocity() float64

	// Calculates and returns mechanical energy.
	Mechanical() float64

	// Sets potential energy.
	// Has a side effect on falling status and kinetic energy.
	SetPotential(potential float64)

	// Exerts force on object.
	// Has a side effect on falling status and kinetic energy.
	ExertForce(force, time float64)
}

type energy struct {
	isFalling bool
	mass      float64
	potential float64
	kinetic   float64
}

func NewEnergy(mass, potential, kinetic float64) Energy {
	return &energy{
		isFalling: true,
		mass:      mass,
		potential: potential,
		kinetic:   kinetic,
	}
}

func (e *energy) Potential() float64 {
	return e.potential
}

func (e *energy) Kinetic() float64 {
	return e.kinetic
}

func (e *energy) Velocity() float64 {
	coefficient := 1.0
	if e.isFalling {
		coefficient = -1.0
	}

	return coefficient *
		math.Sqrt(2*e.Kinetic()/e.mass) // V = √(2K/m)
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
	if e.Kinetic() < 0.0005 {
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
	// HACK: This is a workaround and is not according to physics
	if e.Potential() == 0 {
		e.kinetic = 0.6 * e.Kinetic()
	}
}

func (e *energy) ExertForce(force, time float64) {
	v0 := e.Velocity()
	dv := force * time / e.mass // ΔV = fΔt/m

	v := v0 + dv
	if v0*dv <= 0 && math.Abs(dv) > math.Abs(v0) {
		e.isFalling = !e.isFalling
	}

	e.kinetic = 0.5 * e.mass * math.Pow(v, 2) // K = 1/2mV²
}
