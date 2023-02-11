# Сценарии использования

1. Одиночная игра
```go
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
}
```
2. Многопользовательская игра с совместным прохождением
```go
type Game struct {
	Audio        engine.Audio
	Draw         engine.Drawer
	Shaders      engine.Shaders
	Text         engine.Text
	Menu         game.Menu
	Story        game.Story
	World        game.World
	Multiplayer  game.Singleplayer
	Combat       game.Combat
	Collisions   game.Collisions
	Achievements game.Achievements
	Dialogs      game.Dialogs
	Inventory    game.Inventory
	Equipment    game.Equipment
	Skills       game.Skills
	Quests       game.Quests
	Map          game.Map
	Chat         network.Chat
	Interaction  network.Interaction
}
```
3. Многопользовательская игра с режимом PvP (игрок против игрока)
```go
type Game struct {
	Audio        engine.Audio
	Draw         engine.Drawer
	Shaders      engine.Shaders
	Text         engine.Text
	Menu         game.Menu
	World        game.World
	Multiplayer  game.Singleplayer
	Combat       game.Combat
	Collisions   game.Collisions
	Inventory    game.Inventory
	Equipment    game.Equipment
	Skills       game.Skills
	Map          game.Map
	Chat         network.Chat
	Interaction  network.Interaction
}
```