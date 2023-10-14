package nearest

import (
	"math"
)

type Point interface {
	GetX() float64
	GetY() float64
}

type Node struct {
	Index    int
	Distance float64
}

func GetSortedDistances(clientX, clientY float64, points Point) float64 {
	return math.Pow(points.GetX()-clientX, 2) + math.Pow(points.GetY()-clientY, 2)
}
