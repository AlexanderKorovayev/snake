/*
package game
модуль start_menu
отвечает за отричовку старта игры.
*/

package game

import (
	"github.com/AlexanderKorovaev/snake/core"
	"github.com/JoelOtter/termloop"
)

//StartLevel основной уровень игры
type StartLevel struct {
	termloop.Level
	StartMenu *StartMenu
}

//StartMenu объект стартовое окно
type StartMenu struct {
	*termloop.Text
}

//CreateStartMenu стартовое состояние игры
func CreateStartMenu() *StartMenu {
	startObj := new(StartMenu)
	startObj.Text = termloop.NewText(0, 0, "press Enter", termloop.ColorWhite,
		termloop.ColorDefault)
	return startObj
}

//Tick отслеживаем события
func (s *StartMenu) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		if event.Key == termloop.KeyEnter {
			level := snakeLevel()
			core.TermloopGame.Screen().SetLevel(level)
		}
	}
}

// что бы исключить взаимоимпорты надо переносить всё в гейм и создавать новый отдельный уровень
