package energy

type energy struct {
	isFalling bool
	potential float64
	speed     float64
}

func (self *energy) IsFalling() bool {
	return self.isFalling
}

func (self *energy) Potential() float64 {
	return self.potential
}

func (self *energy) Speed() float64 {
	return self.speed
}

func (self *energy) Mechanical() float64 {
	return self.potential + self.speed
}

// (U2 + V2) - (U1 + V1) = 0
// => V2 = (U1 + V1) - U2
func (self *energy) SetPotential(potential float64) {
	prev := *self

	self.potential = potential
	self.speed = prev.Mechanical() - self.Potential()

	if self.Potential() <= 0 {
		self.isFalling = false
	}
	if self.Speed() <= 0 {
		self.isFalling = true
	}
}