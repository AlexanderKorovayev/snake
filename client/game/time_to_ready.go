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
func CreateTimeObj(val string) *timeToReady {
	timeObj := new(timeToReady)
	timeObj.Text = termloop.NewText((width/2)-1, (high/2)-1, val,
		termloop.ColorWhite,
		termloop.ColorDefault)
	return timeObj
}
