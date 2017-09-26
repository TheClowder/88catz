//https://www.codewars.com/kata/pick-peaks

package kata

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	peaks := []int{}
	positions := []int{}
	peak := 0
	peakpos := 0
	for pos, i := range array {
		if pos != 0 {
			if i > array[pos-1] {
				peak = i
				peakpos = pos
			} else if i < array[pos-1] && peak != 0 {
				positions, peaks = append(positions, peakpos), append(peaks, peak)
				peakpos = 0
				peak = 0
			}

		}
	}
	return PosPeaks{positions, peaks}
}
