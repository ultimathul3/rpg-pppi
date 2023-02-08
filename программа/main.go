package main

import "game/game"

type Game struct {
	Combat       game.Combat
	Collisions   game.Collisions
	Achievements game.Achievements
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
