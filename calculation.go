package simplemoduleproject

import (
	"math"
)

type Hitung interface {
	luas() float64
	keliling() float64
}

type Persegi struct {
	sisi float64
}

func Sum(x int, y int) int {
	return x + y
}

func CalculateCircle(diameter float64) (float64, float64) {
	var luas = math.Phi * math.Pow(diameter/2, 2)
	var keliling = math.Phi * diameter
	return luas, keliling
}

func (p Persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p Persegi) keliling() float64 {
	return p.sisi * 4
}
