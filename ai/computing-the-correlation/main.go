package main

import (
	"fmt"
	"math"
)

var N int
var floatN, onedivNminsone float64

//
func main() {
	var mi, pi, chi int
	var m_sum, p_sum, ch_sum float64 = 0.0, 0.0, 0.0

	fmt.Scanf("%d", &N)
	floatN = float64(N)
	onedivNminsone = 1.0 / (floatN - 1)

	var math_notes []float64 = make([]float64, N)
	var phys_notes []float64 = make([]float64, N)
	var chm_notes []float64 = make([]float64, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d%d%d", &mi, &pi, &chi)

		var m, p, ch float64 = float64(mi), float64(pi), float64(chi)
		m_sum, p_sum, ch_sum = m_sum+m, p_sum+p, ch_sum+ch
		math_notes[i], phys_notes[i], chm_notes[i] = m, p, ch
	}

	// fmt.Println(chm_notes)
	// fmt.Println(len(chm_notes))

	var math_mean, phys_mean, chm_mean float64 = m_sum / floatN, p_sum / floatN, ch_sum / floatN

	var math_sd, phys_sd, chm_sd float64 = sDvt(math_notes, math_mean), sDvt(phys_notes, phys_mean), sDvt(chm_notes, chm_mean)

	mVSp := corelation(math_notes, phys_notes, math_mean, phys_mean, math_sd, phys_sd)

	pVSch := corelation(phys_notes, chm_notes, phys_mean, chm_mean, phys_sd, chm_sd)

	mVSch := corelation(math_notes, chm_notes, math_mean, chm_mean, math_sd, chm_sd)

	// fmt.Println(mVSp, pVSch, mVSch)

	fmt.Println(_round(mVSp, 2))
	fmt.Println(_round(pVSch, 2))
	fmt.Println(_round(mVSch, 2))
}

//
func corelation(notes_a, notes_b []float64, mean_a, mean_b, sd_a, sd_b float64) float64 {
	var sum_corelation float64 = 0.0

	for i := 0; i < N; i++ {
		sum_corelation += ((notes_a[i] - mean_a) / sd_a) * ((notes_b[i] - mean_b) / sd_b)
	}
	return sum_corelation * onedivNminsone
}

//
func sDvt(nums []float64, m float64) float64 {
	var sd float64 = 0.0
	for i := 0; i < N; i++ {
		sd += math.Pow(nums[i]-m, 2)
	}
	return math.Sqrt(sd / floatN)
}

//
func _round(val float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= .5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}
