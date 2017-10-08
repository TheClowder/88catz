//https://www.codewars.com/kata/ball-upwards

package kata

func round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}

func MaxBall(v0 int) int {
	return int(round((float64(v0)/3.6)/0.981, 1.0))
}
