package game

import (
	tl "github.com/JoelOtter/termloop"
)

// Rungame функция для инициализации основных объектов игры и её запуска
func Rungame() {
	tlgame := tl.NewGame()
	tlgame.Screen().SetFps(30)

	level := GetArea()

	snake := Snake{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}

	snake1 := Snake{
		Entity: tl.NewEntity(1, 2, 1, 1),
		level:  level,
	}

	snake2 := Snake{
		Entity: tl.NewEntity(1, 3, 1, 1),
		level:  level,
	}

	// Set the character at position (0, 0) on the entity.
	snake.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	snake1.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	snake2.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})

	level.AddEntity(&snake)
	level.AddEntity(&snake1)
	level.AddEntity(&snake2)
	// установим созданный уровень на экран
	tlgame.Screen().SetLevel(level)
	// запустим игру
	tlgame.Start()
}
