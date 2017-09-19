//https://www.codewars.com/kata/disease-spread/train/go

package kata

func Epidemic(tm, n, s0, i0 int, b, a float64) int {
	susceptible := make([]float64, n)
	infecteds := make([]float64, n)
	recovered := make([]float64, n)
	susceptible[0] = float64(s0)
	infecteds[0] = float64(i0)
	recovered[0] = 0.0
	deltaT := float64(tm) / float64(n)
	for i := 1; i < n-1; i++ {
		susceptible[i] = susceptible[i-1] - deltaT*b*susceptible[i-1]*infecteds[i-1]
		infecteds[i] = infecteds[i-1] + deltaT*(b*susceptible[i-1]*infecteds[i-1]-a*infecteds[i-1])
		recovered[i] = recovered[i-1] + deltaT*infecteds[i-1]*a
	}

	max := 0.0
	for i := 1; i < n; i++ {
		if max < infecteds[i] {
			max = infecteds[i]
		}
	}

	return int(max)
}
