/*
package core
модуль time_to_ready
содержит объекты, которые змейка может съесть
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//createTimeObj создать отрисовку обратного отсчёта
func createTimeObj(text string) *timeToReady {
	timeObj := new(timeToReady)
	timeObj.Text = termloop.NewText(7, (high/2)-1,
		text,
		termloop.ColorWhite,
		termloop.ColorDefault)
	return timeObj
}
