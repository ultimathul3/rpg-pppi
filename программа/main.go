package main

import "game/engine"

type Game struct {
	Audio engine.Audio
	Draw  engine.Drawer
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
