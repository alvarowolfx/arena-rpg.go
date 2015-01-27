package lib

import (
	"math"
)

type Arena struct {
	Width      int
	Height     int
	Combatents []*Character
}

func NewArena(width int, height int) *Arena {
	return &Arena{Width: width, Height: height}
}

func (arena *Arena) InCombate(c1 *Character, c2 *Character) bool {
	if arena.InArena(c1) && arena.InArena(c2) {
		distance := arena.distanceBetween(c1, c2)
		max_weapon_distance := math.Max(float64(c1.Weapon.Distance), float64(c2.Weapon.Distance))

		return distance < max_weapon_distance

	} else {
		return false
	}
}

func (arena *Arena) distanceBetween(c1 *Character, c2 *Character) float64 {
	p1, p2 := c1.Position, c2.Position
	a := math.Pow(float64(p1.X-p2.X), 2)
	b := math.Pow(float64(p1.Y-p2.Y), 2)
	return math.Sqrt(a + b)
}

func (arena *Arena) InArena(char *Character) bool {
	return arena.getCombatent(char) != nil
}

func (arena *Arena) getCombatent(char *Character) *Character {
	for _, charInArena := range arena.Combatents {
		if charInArena.Name == char.Name {
			return charInArena
		}
	}
	return nil
}

func (arena *Arena) PutCombatent(char *Character, x int, y int) bool {
	if arena.validPosition(x, y) {
		arena.Combatents = append(arena.Combatents, char)
		char.Position.X = x
		char.Position.Y = y
		return true
	} else {
		return false
	}
}

func (arena *Arena) validPosition(x int, y int) bool {
	return x <= arena.Width && y <= arena.Height && x >= 0 && y >= 0
}

func (arena *Arena) Move(char *Character, x int, y int) bool {
	if arena.validPosition(x, y) && arena.InArena(char) {
		char.Position.X = x
		char.Position.Y = y
		return true
	} else {
		return false
	}
}
