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
	square := 0.0
	fls := float64(sidesNum)
	switch sidesNum {
	case 0:
		square := math.Pi * sideLen
		return square
	case 4:
		square := fls * sideLen
		return square
	case 3:
		square := math.Sqrt(fls) / 4 * math.Pow(sideLen, 2)
		return square
	default:
		return square
	}
}
