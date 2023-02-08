package main

import "game/game"

type Game struct {
	Inventory game.Inventory
	Equipment game.Equipment
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
