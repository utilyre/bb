package me

func New(potential, speed float64) MechanicalEnergy {
	return MechanicalEnergy{
		isFalling: true,
		potential: potential,
		speed:     speed,
	}
}
