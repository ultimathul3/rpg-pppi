package main

import "game/game"

type Game struct {
	Menu  game.Menu
	Story game.Story
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
