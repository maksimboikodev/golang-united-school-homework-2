package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type MyType int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
	Pi            = 3.1415926
)

func CalcSquare(sideLen float64, sidesNum MyType) float64 {
	if sidesNum == 4 {
		return (math.Pow(sideLen, 2))
	} else if sidesNum == 3 {
		return ((math.Pow(sideLen, 2) * math.Sqrt(3)) / 4)
	} else if sidesNum == 0 {
		return (math.Pi * math.Pow(sideLen, 2))
	} else {
		return 0
	}
}
