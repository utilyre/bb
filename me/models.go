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

// (U2 + V2) - (U1 + V1) = 0
// => V2 = (U1 + V1) - U2
func (self *MechanicalEnergy) SetPotential(potential float64) {
	prev := *self

	self.potential = potential
	self.speed = prev.Total() - self.Potential()

	self.checkFalling()
}

// (U2 + V2) - (U1 + V1) = 0
// => U2 = (U1 + V1) - V2
func (self *MechanicalEnergy) SetSpeed(speed float64) {
	prev := *self

	self.speed = speed
	self.potential = prev.Total() - self.Speed()

	self.checkFalling()
}

func (self *MechanicalEnergy) checkFalling() {
	if self.Potential() <= 0 {
		self.isFalling = false
	}

	if self.Speed() <= 0 {
		self.isFalling = true
	}
}
