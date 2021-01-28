/*
package core
модуль globals
хранит объекты для общего доступа
*/

package core

// Width ширина поля, по факту граница рисуется на 34 пикселе, и змейке достаётся 33 в ширину
var Width int

// High высота поля, по факту граница рисуется на 14 пикселе, и змейке достаётся 13 в ширину
var High int

// MaxObjectsCount максимальное колличество объектов
var MaxObjectsCount int

// TimeCount переменная для обратного отсчёта
var TimeCount int

// Colors Возможные цвета змейки
var Colors []string

// MainObjects контейнер для хранения координат змеек и еды
var MainObjects map[string][]Coordinates = map[string][]Coordinates{}

// ColorMap хранит в себе цвет для каждой змейки
// так же используем для хранения колличества подключённых клиентов
var ColorMap map[string]string = map[string]string{}
