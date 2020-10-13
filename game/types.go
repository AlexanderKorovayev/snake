// модуль содержит нестордартные типы
package game

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
