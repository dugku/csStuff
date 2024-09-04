package main

import (
	"fmt"
	"math"
)

type vector struct {
	x, y, z float64
}

type vectors []vector

type singleVect struct {
	x, y, z float64
}

func main() {

	dustPoint := make(map[string]vectors)

	dustPoint["blue"] = vectors{
		{534.395020, 1094.046143, 1.701973},
		{965.420532, 1097.945068, 0.868401},
		{964.536804, 1195.968628, 36.031250},
		{534.042419, 1157.227661, 2.955454},
	}

	dustPoint["sidepit"] = vectors{
		{927.315674, 788.737549, 9.031250},
		{1227.969849, 780.331116, 9.031250},
		{1227.968628, 215.031250, 11.459801},
		{968.031250, 215.031342, 13.133325},
	}

	dustPoint["pit"] = vectors{
		{1292.029297, 777.767090, -4.984802},
		{1574.417114, 788.031067, -7.672363},
		{1571.968750, 201.031326, -179.261719},
		{1292.030029, 201.030350, -180.752197},
	}

	dustPoint["longdoors"] = vectors{
		{539.031311, 342.934204, 1.567197},
		{740.578064, 341.583191, 0.527790},
		{740.968689, 695.227600, 76.031250},
		{539.223206, 696.461975, 1.394940},
	}

	singlePoint := singleVect{
		x: 660,
		y: 550,
	}

	inter, where := raycast(singlePoint, dustPoint)

	fmt.Println(inter, where)
}

func raycast(point singleVect, edges map[string]vectors) (bool, string) {
	xpoint := point.x
	ypoint := point.y
	count := 0
	where := ""

	for josh, edge := range edges {

		for i := 0; i < len(edge)-1; i++ {

			x1 := edge[i].x
			y1 := edge[i].y
			x2 := edge[i+1].x
			y2 := edge[i+1].y

			if y1 == y2 {
				continue // Skip horizontal edges
			}
			if ypoint > math.Min(y1, y2) && ypoint <= math.Max(y1, y2) {

				// Check if the point lies to the left of the maximum x-value of the edge
				if xpoint <= math.Max(x1, x2) {

					// Check if the ray crosses the edge
					if ((ypoint < y1) != (ypoint < y2)) && (xpoint < x1+((ypoint-y1)/(y2-y1))*(x2-x1)) {

						// Store the current edge identifier
						where = josh

						// Increment the crossing count
						count = count + 1
					}
				}
			}
		}
	}
	fmt.Println(count)
	return count%2 == 1, where
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
