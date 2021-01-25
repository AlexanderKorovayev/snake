/*
package core
модуль globals
хранит объекты для общего доступа
*/

package core

// Width ширина поля, по факту граница рисуется на 34 пикселе, и змейке достаётся 33 в ширину
const Width int = 46

// High высота поля, по факту граница рисуется на 14 пикселе, и змейке достаётся 13 в ширину
const High int = 15

// MaxObjectsCount максимальное колличество объектов
const MaxObjectsCount int = 4

// MainObjects контейнер для хранения координат змеек и еды
var MainObjects map[string][]Coordinates = map[string][]Coordinates{}

// TimeCount переменная для обратного отсчёта
var TimeCount int = 5

// ColorMap хранит в себе цвет для каждой змейки
// так же используем для хранения колличества подключённых клиентов
var ColorMap map[string]string = map[string]string{}

// Colors Возможные цвета змейки
var Colors []string = []string{"Blue", "Cyan", "Green", "Yellow"}
