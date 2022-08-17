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

func (self *MechanicalEnergy) SetPotential(potential float64) {
	self.potential = potential
	self.velocity = self.Total() - self.Potential()
}

func (self *MechanicalEnergy) SetVelocity(velocity float64) {
	self.velocity = velocity
	self.potential = self.Total() - self.Velocity()
}
