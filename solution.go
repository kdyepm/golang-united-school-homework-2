package square

import (
	"math"
)

type sidesNum int

const (
	SidesCircle   sidesNum = 0
	SidesTriangle sidesNum = 3
	SidesSquare   sidesNum = 4
)

func CalcSquare(sideLen float64, sidesNum sidesNum) float64 {

	switch sidesNum {
	case 0:
		square := math.Pi * sideLen
		return square
	case 4:
		square := float64(sidesNum) * sideLen
		return square
	case 3:
		square := math.Sqrt(float64(sidesNum)) / 4 * math.Pow(sideLen, 2)
		return square
	}
	return 0
}
