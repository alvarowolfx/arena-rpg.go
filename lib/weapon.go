package lib

type Weapon struct {
	Name     string
	Distance int
	Damage   int
}

func NewWeapon(name string, distance int, damage int) *Weapon {
	return &Weapon{Name: name, Distance: distance, Damage: damage}
}
