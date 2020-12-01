/*
package core
модуль globals
хранит объекты для общего доступа
*/

package core

import (
	"github.com/JoelOtter/termloop"
)

//GameScreen глобальная переменная которая хранит основные объекты уровня
//в начале игры мы передаём эту переменную в termloop и меняя сзначеия этой
//переменной мы можем менять происходящее на уровне
var GameScreen *Game

//TermloopGame переменная игры нам нужна для динамической смены левелов
var TermloopGame *termloop.Game

const gameOverLogo string = "gameover-logo.txt"

//width ширина поля, по факту граница рисуется на 34 пикселе, и змейке достаётся 33 в ширину
const width int = 35

//high высота поля, по факту граница рисуется на 14 пикселе, и змейке достаётся 13 в ширину
const high int = 15

//максимальное колличество игроков
const maxPlayer int = 4
