package main

import "fmt"

type SpesialPosition struct { //расширение класса Position
	Position //наследование функциональности от класса (структуры реализующей класс) Position
}

func (sp *SpesialPosition) MoveSpecial(x, y float64) {
	sp.x = x * x
	sp.y = y * y
}

// базовый класс
type Position struct {
	x float64
	y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}
func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

// /Примеры создания экжемпляров класса
type Player struct {
	*Position
}

func NewPlayer() *Player { //инициализация
	return &Player{
		Position: &Position{},
	}
}

type Enumu struct {
	*Position
}

func NewEnumu() *Enumu { //инициализация
	return &Enumu{
		Position: &Position{},
	}
}

type SpecialPlayer struct {
	*SpesialPosition
}

func NewSpecialPlayer() *SpecialPlayer {
	return &SpecialPlayer{
		SpesialPosition: &SpesialPosition{},
	}
}

func main() {
	sid := NewPlayer()
	sid.Move(1, 2)
	sid.Teleport(3, 4)
	fmt.Println("Player :=", sid.Position)

	sib := NewEnumu()
	sib.Move(1, 2)
	sib.Teleport(3, 4)
	fmt.Println("Enumu :=", sid.Position)

	fmt.Println("SpecPlayer:")
	specsid := NewSpecialPlayer()
	specsid.Move(20, 24)
	fmt.Println(specsid.Position)
	specsid.Teleport(3333, 4444)
	fmt.Println(specsid.Position)
	specsid.MoveSpecial(20, 20)
	fmt.Println(specsid.Position)
}
