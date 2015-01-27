package lib

type Character struct {
	Name     string
	Hp       int
	Weapon   *Weapon
	Position *Position
}

func NewCharacter(name string, hp int, weapon *Weapon) *Character {
	return &Character{Name: name, Hp: hp, Weapon: weapon, Position: NewPosition(0, 0)}
}

func (c *Character) setPosition(x int, y int) {
	if c.Position != nil {
		c.Position.X = x
		c.Position.Y = y
	} else {
		c.Position = NewPosition(x, y)
	}
}
