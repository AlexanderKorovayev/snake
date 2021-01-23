/*
package core
модуль other_snake
описывает объект змейка других игроков
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//createOtherSnake создать змейку
func createOtherSnake(body []coordinates, name string, color termloop.Attr) *otherSnake {
	snakeObj := new(otherSnake)
	snakeObj.Entity = termloop.NewEntity(1, 1, 1, 1)
	snakeObj.body = body
	snakeObj.name = name
	snakeObj.color = color
	return snakeObj
}

//Draw отвечает за отрисовку змеи на дисплее
func (snake *otherSnake) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	for _, v := range snake.body {
		screen.RenderCell(v.X, v.Y, &termloop.Cell{Fg: termloop.ColorWhite,
			Bg: snake.color})
	}

}
