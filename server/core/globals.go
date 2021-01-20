/*
package core
модуль globals
хранит объекты для общего доступа
*/

package core

import (
	"github.com/JoelOtter/termloop"
)

// TermloopGame переменная игры нам нужна для динамической смены левелов
var TermloopGame *termloop.Game

// Width ширина поля, по факту граница рисуется на 34 пикселе, и змейке достаётся 33 в ширину
const Width int = 46

// High высота поля, по факту граница рисуется на 14 пикселе, и змейке достаётся 13 в ширину
const High int = 15

// MaxObjectsCount максимальное колличество объектов
const MaxObjectsCount int = 4

// MainObjects контейнер для хранения координат змеек и еды
var MainObjects map[string][]Coordinates = map[string][]Coordinates{}

// ClientsCount переменная для хранения колличества подключённых клиентов
// будем использовать словарь с пустыми значениями, такой подход быстро
// позволяет проверять наличие елемента в себе
var ClientsCount map[string]string

// TimeCount переменная для обратного отсчёта
var TimeCount int = 5

// ColorMap хранит в себе цвет для каждой змейки
var ColorMap map[string]string = map[string]string{}

// Colors Возможные цвета змейки
var Colors []string = []string{"Blue", "Cyan", "Green", "Yellow"}
