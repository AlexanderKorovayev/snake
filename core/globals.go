/*
package core
модуль globals
хранит объекты для общего доступа
*/

package core

//GameScreen глобальная переменная которая хранит основные объекты уровня
//в начале игры мы передаём эту переменную в termloop и меняя сзначеия этой
//переменной мы можем менять происходящее на уровне
var GameScreen *Game

//Width ширина поля
const Width int = 35

//High высота поля
const High int = 15
