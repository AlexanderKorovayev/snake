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

//gameOver функция для отображения конца игры
func gameOver() {
	GameScreen.Level.RemoveEntity(GameScreen.Snake)
	GameScreen.Level.RemoveEntity(GameScreen.GameArea)
	GameScreen.Level.RemoveEntity(GameScreen.GameFood)

	dat, _ := ioutil.ReadFile("gameover-logo.txt")
	e := termloop.NewEntityFromCanvas(1, 1, termloop.CanvasFromString(string(dat)))
	GameScreen.AddEntity(e)
}
