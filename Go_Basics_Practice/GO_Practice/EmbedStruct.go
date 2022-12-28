package main

import (
	"fmt"
)

// struct is like class and create methods
// embed like classes  extends another class so that it can access it super parent level methods too
// create a new object using classs or new keyword

func (s *SpecialPosition) MoveSpecial(x, y float64) {
	fmt.Println(s)
	s.x *= x
	s.y *= y
	fmt.Println(s.x, s.y)
}

type Position struct {
	x, y float64
}
type SpecialPosition struct {
	Position
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
	fmt.Println(p.x, p.y)
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
	fmt.Println(p.x, p.y)
}

// embed one struct in to another (static composition)
// so that it can access its super level methods
type Player struct {
	*Position
}

// create a new object using classs or new keyword
// 
func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*SpecialPosition
}

// create a new object using classs or new keyword
func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func main() {
	player := NewPlayer()
	player.Move(2, 3)
	player.Move(2, 3)
	player.Teleport(3, 6)
	fmt.Println(*player.Position)

	enemy := NewEnemy()
	enemy.Move(3, 4)
	enemy.MoveSpecial(2, 4)
	//fmt.Println(enemy.SpecialPosition)
	enemy.Teleport(8, 9)

}
