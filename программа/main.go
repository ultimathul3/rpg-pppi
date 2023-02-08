package main

import "game/game"

type Game struct {
	World        game.World
	Singleplayer game.Singleplayer
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
