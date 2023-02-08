package main

import "game/engine"

type Game struct {
	Audio   engine.Audio
	Draw    engine.Drawer
	Shaders engine.Shaders
	Text    engine.Text
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
