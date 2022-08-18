package me

func New(potential, speed float64) MechanicalEnergy {
	return MechanicalEnergy{
		potential: potential,
		speed:     speed,
	}
}
