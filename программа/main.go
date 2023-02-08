package main

import (
	"game/engine"
	"game/game"
)

type Game struct {
	Audio        engine.Audio
	Draw         engine.Drawer
	Shaders      engine.Shaders
	Text         engine.Text
	Menu         game.Menu
	Story        game.Story
	World        game.World
	Singleplayer game.Singleplayer
	Combat       game.Combat
	Collisions   game.Collisions
	Achievements game.Achievements
	Saves        game.Saves
	Dialogs      game.Dialogs
	Inventory    game.Inventory
	Equipment    game.Equipment
	Skills       game.Skills
	Quests       game.Quests
	Map          game.Map
	Multiplayer  game.Multiplayer
}

func (g *Game) Run() {

}

func main() {
	game := &Game{}
	game.Run()
}
