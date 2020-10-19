package game

import "github.com/JoelOtter/termloop"

//CreateArea создать арену, по которой будет перемещаться змейка
func CreateArea() *Area {
	areaObj := new(Area)
	areaObj.Entity = termloop.NewEntity(1, 1, 1, 1)
	var border map[coordinates]int
	border = make(map[coordinates]int)
	fillAreaBorder(35, 20, &border)
	areaObj.areaBorder = border
	return areaObj
}

func fillAreaBorder(imax, jmax int, border *map[coordinates]int) {
	starti := 0
	startj := 0
	for i := starti; i < imax; i++ {
		for j := startj; j < jmax; j++ {
			coord := coordinates{i, j}
			if i == starti || i == imax-1 {
				(*border)[coord] = 1
			} else if j == startj || j == jmax-1 {
				(*border)[coord] = 1
			} else {
				(*border)[coord] = 0
			}
		}
	}
}

//Draw отвечает за отрисовку змеи на дисплее
func (area *Area) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	for k, v := range area.areaBorder {
		if v == 1 {
			screen.RenderCell(k.x, k.y, &termloop.Cell{Fg: termloop.ColorWhite,
				Bg: termloop.ColorWhite})
		}
	}
}
