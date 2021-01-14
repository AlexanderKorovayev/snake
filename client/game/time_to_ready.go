/*
package core
модуль time_to_ready
содержит объекты, которые змейка может съесть
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//CreateTimeObj создать отрисовку обратного отсчёта
func CreateTimeObj(text string) *timeToReady {
	timeObj := new(timeToReady)
	timeObj.Text = termloop.NewText(3, (high/2)-1,
		text,
		termloop.ColorWhite,
		termloop.ColorDefault)
	return timeObj
}
