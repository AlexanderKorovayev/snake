/*
package game
модуль start_level
отвечает за отрисовку стартового меню.
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//startLevel стартовый уровень игры
type startLevel struct {
	termloop.Level
	startMenu *startMenu
}

//startMenu объект стартовое окно
type startMenu struct {
	*termloop.Text
}

//createStartMenu стартовое состояние игры
func createStartMenu() *startMenu {
	startObj := new(startMenu)
	startObj.Text = termloop.NewText(0, 0, "press Enter to srart game",
		termloop.ColorWhite,
		termloop.ColorDefault)
	return startObj
}

//Tick отслеживаем события
func (s *startMenu) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		// если нажали enter, то начинаем этап подготовки всех играков к игре
		if event.Key == termloop.KeyEnter {
			// сначала установим базовую часть уровня
			// и уже она будет по своему тику пинать сервер для ожидания начала игры
			level := startBaseSnakeLevel()
			TermloopGame.Screen().SetLevel(level)
		}
	}
}
