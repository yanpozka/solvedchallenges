package main

import "fmt"

func main() {
	undergroundSystem := Constructor()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)  // Customer 45 "Leyton" -> "Waterloo" in 15-3 = 12
	undergroundSystem.CheckOut(27, "Waterloo", 20)  // Customer 27 "Leyton" -> "Waterloo" in 20-10 = 10
	undergroundSystem.CheckOut(32, "Cambridge", 22) // Customer 32 "Paradise" -> "Cambridge" in 22-8 = 14
	// return 14.00000. One trip "Paradise" -> "Cambridge", (14) / 1 = 14
	fmt.Println(undergroundSystem.GetAverageTime("Paradise", "Cambridge"))

	// return 11.00000. Two trips "Leyton" -> "Waterloo", (10 + 12) / 2 = 11
	fmt.Println(undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
	undergroundSystem.CheckIn(10, "Leyton", 24)
	fmt.Println(undergroundSystem.GetAverageTime("Leyton", "Waterloo")) // return 11.00000
	undergroundSystem.CheckOut(10, "Waterloo", 38)                      // Customer 10 "Leyton" -> "Waterloo" in 38-24 = 14

	// return 12.00000. Three trips "Leyton" -> "Waterloo", (10 + 12 + 14) / 3 = 12
	fmt.Println(undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
}

type station struct {
	name string
	time int
}

type UndergroundSystem struct {
	checkIns   map[int]*station
	sumTimes   map[string]float64
	totalTimes map[string]float64
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkIns:   make(map[int]*station),
		sumTimes:   make(map[string]float64),
		totalTimes: make(map[string]float64),
	}
}

func (u *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	u.checkIns[id] = &station{
		name: stationName,
		time: t,
	}
}

func key(startStation, endStation string) string {
	return startStation + "-" + endStation
}

func (u *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	s := u.checkIns[id]
	k := key(s.name, stationName)
	u.sumTimes[k] += float64(t - s.time)
	u.totalTimes[k]++
}

func (u *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	k := key(startStation, endStation)
	return u.sumTimes[k] / u.totalTimes[k]
}
