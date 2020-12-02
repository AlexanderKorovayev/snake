/*
package core
модуль types
модуль содержит новые типы, которые необходимы для удобства в работе
*/

package game

//coordinates координаты
type coordinates struct {
	x int
	y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)
