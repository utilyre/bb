package me

func New(potential, velocity float64) MechanicalEnergy {
	return MechanicalEnergy{
		potential: potential,
		velocity:  velocity,
	}
}
