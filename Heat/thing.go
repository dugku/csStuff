package main

import (
	"fmt"
	"math"
)

type vectors struct {
	x, y, z float64
}

func main() {
	dustPoint := make(map[string]map[string]vectors)

	dustPoint["outside long"] = map[string]vectors{
		"p1": {100.346130, 287.698090, 1.515508},
		"p2": {447.968506, 299.959595, 4.944149},
		"p3": {747.964172, 235.968674, 9.031250},
		"p4": {758.814636, -395.971558, 68.031250},
		"p5": {116.031776, -372.452881, 2.015634},
	}

	for _, i := range dustPoint {
		for _, x := range i {
			fmt.Println(x.x, x.y, x.x+1, x.y+1)
			//fmt.Println(distance(x.x, x.y, x.x+1, x.y+1))
		}
	}
}

//

func normalize() {

}

func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
}
