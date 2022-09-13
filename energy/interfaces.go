package energy

type Energy interface {
	IsFalling() bool
	Potential() float64
	Kinetic() float64
	Mechanical() float64
	SetPotential(velocity float64)
}

func NewEnergy(potential, kinetic float64) Energy {
	return &energy{
		isFalling: true,
		potential: potential,
		kinetic:   kinetic,
	}
}
