package energy

type Energy interface {
	IsFalling() bool
	Potential() float64
	Speed() float64
	Mechanical() float64
	SetPotential(potential float64)
}

func NewEnergy(potential, speed float64) Energy {
	return &energy{
		isFalling: true,
		potential: potential,
		speed:     speed,
	}
}
