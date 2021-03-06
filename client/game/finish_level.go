/*
package game
модуль start_level
отвечает за отрисовку стартового меню.
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//finishLevel стартовый уровень игры
type finishLevel struct {
	termloop.Level
	finishMenu *finishMenu
}

//finishMenu объект финишное окно
type finishMenu struct {
	*termloop.Text
}

//createFinishMenu финишное состояние игры
func createFinishMenu() *finishMenu {
	finishObj := new(finishMenu)
	finishText := "Game over///press Enter to srart game///press cntrl+c to exit"
	finishObj.Text = termloop.NewText(0, 0, finishText,
		termloop.ColorWhite,
		termloop.ColorDefault)
	return finishObj
}

//Tick отслеживаем события
func (s *finishMenu) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		if event.Key == termloop.KeyEnter {
			// надо вставить правильную функцию
			//level := startSnakeLevel()
			//TermloopGame.Screen().SetLevel(level)
		}
	}
}
