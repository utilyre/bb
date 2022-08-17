package me

type MechanicalEnergy struct {
	potential float64
	velocity  float64
}

func (self MechanicalEnergy) Potential() float64 {
	return self.potential
}

func (self MechanicalEnergy) Velocity() float64 {
	return self.velocity
}

func (self MechanicalEnergy) Total() float64 {
	return self.potential + self.velocity
}

func (self *MechanicalEnergy) SetPotential(prev MechanicalEnergy, potential float64) {
	// (U2 + V2) - (U1 + V1) = 0
	// => V2 = (U1 + V1) - U2
	self.potential = potential
	self.velocity = prev.Total() - self.Potential()
}

func (self *MechanicalEnergy) SetVelocity(prev MechanicalEnergy, velocity float64) {
	// (U2 + V2) - (U1 + V1) = 0
	// => U2 = (U1 + V1) - V2
	self.velocity = velocity
	self.potential = prev.Total() - self.Velocity()
}
