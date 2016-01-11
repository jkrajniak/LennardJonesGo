package sim

import (
	"github.com/quells/LennardJones/vector"
	"math"
	"math/rand"
)

func InitPositionCubic(N int, L float64) [][3]float64 {
	R := make([][3]float64, N)
	Ncube := 1
	for N > Ncube*Ncube*Ncube {
		Ncube++
	}
	rs := L / float64(Ncube)
	roffset := (L - rs) / 2
	i := 0
	for x := 0; x < Ncube; x++ {
		x := float64(x)
		for y := 0; y < Ncube; y++ {
			y := float64(y)
			for z := 0; z < Ncube; z++ {
				z := float64(z)
				pos := vector.Scale([3]float64{x, y, z}, rs)
				offset := [3]float64{roffset, roffset, roffset}
				R[i] = vector.Difference(pos, offset)
				i++
			}
		}
	}
	return R
}

func InitVelocity(N int, T0 float64, M float64) [][3]float64 {
	V := make([][3]float64, N)
	rand.Seed(1)
	netP := [3]float64{0, 0, 0}
	netE := 0.0
	for n := 0; n < N; n++ {
		for i := 0; i < 3; i++ {
			newP := rand.Float64() - 0.5
			netP[i] += newP
			netE += newP * newP
			V[n][i] = newP
		}
	}
	netP = vector.Scale(netP, 1.0/float64(N))
	vscale := math.Sqrt(3.0 * float64(N) * T0 / (M * netE))
	for i, v := range V {
		correctedV := vector.Scale(vector.Difference(v, netP), vscale)
		V[i] = correctedV
	}
	return V
}