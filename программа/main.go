package main

import (
	"game/engine"
	"game/game"
)

type Game struct {
	Audio   engine.Audio
	Draw    engine.Drawer
	Shaders engine.Shaders
	Text    engine.Text
	Menu    game.Menu
	Story   game.Story
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
