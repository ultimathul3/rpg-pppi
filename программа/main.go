package main

import "game/game"

type Game struct {
	Saves   game.Saves
	Dialogs game.Dialogs
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
