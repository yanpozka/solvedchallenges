package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var N int

//
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	var nums []float64 = make([]float64, N)

	for ix := range parts {
		nums[ix], _ = strconv.ParseFloat(parts[ix], 32)
	}
	var Mean float64 = avg(nums)
	fmt.Println(Mean)
	fmt.Println(median(nums))
	fmt.Println(mode(nums))

	Sd := standardDeviation(nums, Mean)

	fmt.Println(_round(Sd, 1))

	ci1, ci2 := confidenceInterval(Mean, Sd)
	fmt.Println(_round(ci1, 1), _round(ci2, 1))
}

//
func confidenceInterval(m, sd float64) (float64, float64) {
	div := (sd * 1.96) / math.Sqrt(float64(N))
	return m - div, m + div
}

//
func standardDeviation(nums []float64, m float64) float64 {
	var sd float64 = 0.0
	for i := 0; i < N; i++ {
		sd += math.Pow(nums[i]-m, 2)
	}
	return math.Pow(sd/float64(N), 0.5)
}

//
func avg(nums []float64) float64 {
	var total float64 = 0
	for i := 0; i < N; i++ {
		total += nums[i]
	}
	return total / float64(N)
}

//
func median(nums []float64) float64 {
	var floatSlice sort.Float64Slice = nums
	floatSlice.Sort()
	lenslice := len(floatSlice)

	if lenslice%2 != 0 {
		return floatSlice[lenslice/2]
	}

	return (floatSlice[lenslice/2] + floatSlice[lenslice/2-1]) / 2.0
}

//
func mode(nums []float64) float64 {
	num_map := make(map[float64]int)

	for i := 0; i < N; i++ {
		if _, exists := num_map[nums[i]]; !exists {
			num_map[nums[i]] = 0
		}
		num_map[nums[i]]++
	}
	var min float64 = 9999999999
	var max_amount int = -1

	for current_n := range num_map {
		amount := num_map[current_n]
		if amount > max_amount {
			max_amount = amount
			min = current_n
		} else {
			if max_amount == amount {
				if current_n < min {
					min = current_n
				}
			}
		}
	}
	return min
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
