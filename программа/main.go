package main

import "game/game"

type Game struct {
	Skills      game.Skills
	Quests      game.Quests
	Map         game.Map
	Multiplayer game.Multiplayer
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
