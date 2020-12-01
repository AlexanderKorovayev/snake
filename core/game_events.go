/*
package core
модуль game_events
описывает основные игровые события
*/

package core

import (
	"io/ioutil"
	"runtime"

	"github.com/JoelOtter/termloop"
)

//gameOver функция для отображения конца игры
func gameOver() {
	GameScreen.Level.RemoveEntity(GameScreen.Snake)
	GameScreen.Level.RemoveEntity(GameScreen.GameArea)
	GameScreen.Level.RemoveEntity(GameScreen.GameFood)

	defaultText := termloop.NewText(0, 0, "Game over, pleace press cntrl+c",
		termloop.ColorWhite, termloop.ColorDefault)

	if runtime.GOOS == "windows" {
		dat, err := ioutil.ReadFile(gameOverLogo)
		if err != nil {
			// если картинка не открылась, то отобразим хотя бы текст
			GameScreen.AddEntity(defaultText)
		}
		canv := termloop.NewEntityFromCanvas(1, 1,
			termloop.CanvasFromString(string(dat)))
		GameScreen.AddEntity(canv)
	} else {
		//в линуксе почему-то лого не открывается, поэтому вставляем текст
		GameScreen.AddEntity(defaultText)
	}
}
