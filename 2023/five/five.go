package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./five/in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	seedsL := getSeeds(scanner.Text())
	scanner.Scan()
	scanner.Text()

	seedToSoil := getMap(scanner)
	soilToFert := getMap(scanner)
	fertToWatter := getMap(scanner)
	waterToLight := getMap(scanner)
	lightToTemp := getMap(scanner)
	tempToHum := getMap(scanner)
	humidityToLoc := getMap(scanner)

	minLoc := int64(math.MaxInt64)
	oldApproach := false

	//locationRanges := make([]Pair, 0)
	//for _, v := range humidityToLoc {
	//	locationRanges = append(locationRanges, Pair{from: v.to, to: v.to+v.rrange})
	//}

	now := time.Now()
	if oldApproach {
		//not optimal generates all seeds
		seeds := make([]int64, 0)

		for i := 0; i < len(seedsL)-1; i += 2 {
			seeds = append(seeds, getSeedsFromL(seedsL[i], seedsL[i+1])...)
		}
		fmt.Println("done generating: ", time.Now().Sub(now).Seconds())
		now = time.Now()
		for _, seed := range seeds {
			s2s := seedToSoil.GetValue(seed, false)
			s2f := soilToFert.GetValue(s2s, false)
			f2w := fertToWatter.GetValue(s2f, false)
			w2l := waterToLight.GetValue(f2w, false)
			l2t := lightToTemp.GetValue(w2l, false)
			t2h := tempToHum.GetValue(l2t, false)
			h2l := humidityToLoc.GetValue(t2h, false)

			if h2l < minLoc {
				minLoc = h2l
			}
		}
		fmt.Println("done processing: ", time.Now().Sub(now).Seconds())
		fmt.Println(minLoc)
	} else {
		//fmt.Println(largestLoc)

		largestLoc := getLargestLoc(humidityToLoc)
		fmt.Println(largestLoc)
		for l := int64(0); l <= largestLoc; l++ {
			h2l := humidityToLoc.GetValue(int64(l), true)
			t2h := tempToHum.GetValue(h2l, true)
			l2t := lightToTemp.GetValue(t2h, true)
			w2l := waterToLight.GetValue(l2t, true)
			f2w := fertToWatter.GetValue(w2l, true)
			s2f := soilToFert.GetValue(f2w, true)
			s2s := seedToSoil.GetValue(s2f, true)

			//fmt.Println(l, h2l, t2h, l2t, w2l, f2w, s2f, s2s)

			if checkSeed(seedsL, s2s) {
				//fmt.Println(s2s)
				fmt.Println(l)
				break
			}
		}
	}
	fmt.Println("time taken ", time.Now().Sub(now))
}

type Pair struct {
	from int64
	to   int64
}

func getLargestLoc(loc Map) int64 {
	max := int64(0)
	for _, l := range loc {
		if l.to > max {
			max = l.to
		}
	}
	return max
}

func checkSeed(seedsL []int64, rStart int64) bool {
	for i := 0; i < len(seedsL)-1; i += 2 {
		if rStart >= seedsL[i] && rStart < seedsL[i]+seedsL[i+1] {
			return true
		}
	}
	return false
}

func getSeedsFromL(from, r int64) []int64 {
	res := make([]int64, 0)

	for i := int64(0); i < r; i++ {
		res = append(res, from+i)
	}
	return res
}

type Map []mmap

type mmap struct {
	from   int64
	to     int64
	rrange int64
}

func (m Map) GetValue(key int64, reverse bool) int64 {
	for _, mm := range m {
		v := mm.GetValue(key, reverse)
		if v != -1 {
			return v
		}
	}
	return key
}
func (m mmap) GetValue(key int64, reverse bool) int64 {
	if reverse {
		// key is the value, find the key (start of range)
		if key >= m.to && key < m.to+m.rrange {
			return m.from + key - m.to
		}
		return -1
	}
	if key >= m.from && key < m.from+m.rrange {
		return key - m.from + m.to
	}
	return -1
}

func getMap(scanner *bufio.Scanner) Map {
	list := make([]mmap, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		} else if strings.Contains(line, "map:") {
			continue
		} else {
			list = append(list, getMapValues(line))
		}
	}
	scanner.Scan()
	scanner.Text()

	//sort
	return list
}

func getMapValues(line string) mmap {
	sp := strings.Split(line, " ")
	to := getNumber(sp[0])
	from := getNumber(sp[1])
	rr := getNumber(sp[2])
	return mmap{
		from:   from,
		to:     to,
		rrange: rr,
	}
}

func getSeeds(l string) []int64 {
	seeds := make([]int64, 0)

	sp := strings.Split(l, ":")
	numStr := strings.Split(sp[1], " ")
	for _, s := range numStr {
		if s == "" || s == " " {
			continue
		}
		seeds = append(seeds, getNumber(s))
	}
	return seeds
}

func getNumber(s string) int64 {
	i := strings.LastIndex(s, " ")
	v, _ := strconv.ParseInt(s[i+1:], 10, 64)
	return v
}
