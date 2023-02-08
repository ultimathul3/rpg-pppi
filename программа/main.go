package main

import "game/network"

type Game struct {
	Chat        network.Chat
	Interaction network.Interaction
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
