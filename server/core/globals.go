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

//Width ширина поля, по факту граница рисуется на 34 пикселе, и змейке достаётся 33 в ширину
const Width int = 35

//High высота поля, по факту граница рисуется на 14 пикселе, и змейке достаётся 13 в ширину
const High int = 15

//максимальное колличество объектов
const MaxObjectsCount int = 4

//MainObjects контейнер для хранения координат змеек и еды
var MainObjects map[string][]Coordinates
