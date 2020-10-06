package game

import (
	tl "github.com/JoelOtter/termloop"
)

// GetArea функция для получения базового уровня для игры
func GetArea() *tl.BaseLevel {
	// уровень так же является основным объектом,
	// именно его мы наполняем другими объектами
	// ячейка является основным объектом при создании уровней.
	// она представляет собой одну ячейку терминала.
	// Базовый уровень рпедставляет из себя основной фон, который мы
	// заполним одной ячейкой, размноженной по основному фону
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen, // цвет ячейки
		Fg: tl.ColorBlack, // цвет символа в ячейке
		Ch: 'v',           // сам символ
	})

	// наполнение уровня объектами
	level.AddEntity(tl.NewRectangle(5, 5, 60, 20, tl.ColorBlue))
	return level
}
