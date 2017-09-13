// https://www.codewars.com/kata/buying-a-car

package kata

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	months := 0
	priceOld := float64(startPriceOld)
	priceNew := float64(startPriceNew)

	for ; priceOld+float64(months*savingperMonth) < priceNew; months++ {
		if months%2 == 1 {
			percentLossByMonth += 0.5
		}
		priceOld -= priceOld * percentLossByMonth / 100.0
		priceNew -= priceNew * percentLossByMonth / 100.0
	}

	return [2]int{months, int(priceOld + float64(months*savingperMonth) - priceNew + 0.5)}
}
