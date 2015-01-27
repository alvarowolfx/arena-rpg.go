package test

import (
	. "dojo.go/lib"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGame(t *testing.T) {
	var espada, lanca, flecha *Weapon
	var arena *Arena

	Convey("Arena de combate", t, func() {

		espada = NewWeapon("Espada", 1, 20)
		lanca = NewWeapon("Lanca", 3, 15)
		arena = NewArena(20, 20)

		Convey("Criação de diferentes armas", func() {

			Convey("Criar espada com dano 20 e distancia 1", func() {

				So(espada.Distance, ShouldEqual, 1)
				So(espada.Damage, ShouldEqual, 20)
			})

			Convey("Criar lanca com dano 15 e distancia 3", func() {

				So(lanca.Distance, ShouldEqual, 3)
				So(lanca.Damage, ShouldEqual, 15)
			})

			Convey("Criar flecha com dano 10 e distancia 5", func() {
				flecha = NewWeapon("Flecha", 5, 10)

				So(flecha.Distance, ShouldEqual, 5)
				So(flecha.Damage, ShouldEqual, 10)
			})

		})

		Convey("Criação de arena", func() {
			Convey("Ao criar uma arena 30x25", func() {

				arena = NewArena(30, 25)

				Convey("O comprimento deve ser 30", func() {
					So(arena.Width, ShouldEqual, 30)
				})
				Convey("O altura deve ser 25", func() {
					So(arena.Height, ShouldEqual, 25)
				})

				Reset(func() {
					arena = NewArena(20, 20)
				})
			})
		})

		Convey("Criar personagem", func() {
			leonidas := NewCharacter("Leonidas", 200, espada)

			So(leonidas.Name, ShouldEqual, "Leonidas")
			So(leonidas.Hp, ShouldEqual, 200)
			Convey("Ao iniciar com uma arma, ele deve estar equipado com a mesma", func() {
				So(leonidas.Weapon, ShouldEqual, espada)
			})

			Convey("A posição inicial deve ser 0,0", func() {
				So(leonidas.Position.X, ShouldEqual, 0)
				So(leonidas.Position.Y, ShouldEqual, 0)
			})

		})

		Convey("Verificar se estao em combate", func() {

			leonidas := NewCharacter("Leonidas", 200, espada)
			hulk := NewCharacter("Hulk", 800, lanca)

			Convey("Colocar os dois personagens proximos", func() {
				arena.PutCombatent(leonidas, 4, 6)
				arena.PutCombatent(hulk, 5, 7)

				Convey("Eles devem estar em combate", func() {
					inCombate := arena.InCombate(leonidas, hulk)
					So(inCombate, ShouldBeTrue)
				})

			})

			Convey("Colocar os dois personagens distantes", func() {
				arena.PutCombatent(leonidas, 10, 15)
				arena.PutCombatent(hulk, 1, 3)

				Convey("Eles não devem estar em combate", func() {
					inCombate := arena.InCombate(leonidas, hulk)
					So(inCombate, ShouldBeFalse)
				})

			})

			Reset(func() {
				arena = NewArena(20, 20)
			})

		})

		Convey("Verificar se as posições são validas", func() {

			stark := NewCharacter("Tony Stark", 150, flecha)

			Convey("Colocar stark em posição invállida", func() {

				putActionResult := arena.PutCombatent(stark, 21, 10)
				inArena := arena.InArena(stark)

				Convey("Não deve conseguir adicionar a arena", func() {
					So(putActionResult, ShouldBeFalse)
				})

				Convey("Não deve estar na arena", func() {
					So(inArena, ShouldBeFalse)
				})

			})

			Reset(func() {
				arena = NewArena(20, 20)
			})

		})

		Convey("Deve mover personagem", func() {

			leonidas := NewCharacter("Leonidas", 200, espada)

			arena.PutCombatent(leonidas, 10, 12)
			So(leonidas.Position.X, ShouldEqual, 10)
			So(leonidas.Position.Y, ShouldEqual, 12)

			arena.Move(leonidas, 6, 7)
			So(leonidas.Position.X, ShouldEqual, 6)
			So(leonidas.Position.Y, ShouldEqual, 7)

			Reset(func() {
				arena = NewArena(20, 20)
			})

		})

	})

}
