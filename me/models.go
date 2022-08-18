package me

type MechanicalEnergy struct {
	isFalling bool
	potential float64
	speed     float64
}

func (self MechanicalEnergy) IsFalling() bool {
	return self.isFalling
}

func (self MechanicalEnergy) Potential() float64 {
	return self.potential
}

func (self MechanicalEnergy) Speed() float64 {
	return self.speed
}

func (self MechanicalEnergy) Total() float64 {
	return self.potential + self.speed
}

func (self *MechanicalEnergy) SetFalling(isFalling bool) {
	self.isFalling = isFalling
}

func (self *MechanicalEnergy) SetPotential(potential float64) {
	prev := *self

	// (U2 + V2) - (U1 + V1) = 0
	// => V2 = (U1 + V1) - U2
	self.potential = potential
	self.speed = prev.Total() - self.Potential()
}

func (self *MechanicalEnergy) SetSpeed(speed float64) {
	prev := *self

	// (U2 + V2) - (U1 + V1) = 0
	// => U2 = (U1 + V1) - V2
	self.speed = speed
	self.potential = prev.Total() - self.Speed()
}
