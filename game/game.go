package game

import (
	tl "github.com/JoelOtter/termloop"
)

// Rungame основным объектом является игра
func Rungame() {
	tlgame := tl.NewGame()
	tlgame.Screen().SetFps(30)

	level := Getarea()

	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)
	// установим созданный уровень на экран
	tlgame.Screen().SetLevel(level)
	// запустим игру
	tlgame.Start()
}
