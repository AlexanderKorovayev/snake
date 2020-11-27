/*
package core
модуль game_events
описывает основные игровые события
*/

package core

import (
	"io/ioutil"

	"github.com/JoelOtter/termloop"
)

//GameOver функция для отображения конца игры
func GameOver() {
	GameScreen.Level.RemoveEntity(GameScreen.snake)
	GameScreen.Level.RemoveEntity(GameScreen.gameArea)
	GameScreen.Level.RemoveEntity(GameScreen.gameFood)

	dat, _ := ioutil.ReadFile("gameover-logo.txt")
	e := termloop.NewEntityFromCanvas(1, 1, termloop.CanvasFromString(string(dat)))
	GameScreen.AddEntity(e)
}
