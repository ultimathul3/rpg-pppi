package main

import "game/engine"

type Game struct {
	Shaders engine.Shaders
	Text    engine.Text
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
