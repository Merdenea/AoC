package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open("./fourteen/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]rune, 0)
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, []rune(l))
	}

	//part one
	//tiltNorth(m)
	//fmt.Println(getLoad(m))

	now := time.Now()

	mm := make(map[string]loadPair)
	loads := make([]int, 0)
	cycles := 1000000000
	cycleStart := -1
	for i := 0; i < cycles; i++ {
		tiltNorth(m)
		tiltEast(m)
		tiltSouth(m)
		tiltWest(m)

		key := getKey(m)
		if v, ok := mm[key]; ok {
			cycleStart = v.index
			break
		}

		l := getLoad(m)
		mm[getKey(m)] = loadPair{l, i}
		loads = append(loads, l)
	}

	fmt.Println(loads)
	fmt.Println(cycleStart)

	loads = loads[cycleStart:]
	fmt.Println(loads)

	fmt.Println(loads[(cycles-cycleStart-1)%len(loads)])

	fmt.Println(time.Now().Sub(now))
}

type loadPair struct {
	load  int
	index int
}

func tiltNorth(m [][]rune) {
	//printRunes(m)
	for j := 0; j < len(m[0]); j++ {
		roll := 0
		for i := 0; i < len(m); i++ {
			if m[i][j] == '#' {
				roll = 0
			} else if m[i][j] == '.' {
				roll++
			} else { // O
				if roll > 0 {
					m[i-roll][j] = 'O'
					m[i][j] = '.'
				}
			}
		}
	}

	//printRunes(m)
}

func getKey(m [][]rune) string {
	sb := strings.Builder{}

	for _, s := range m {
		sb.WriteString(string(s))
	}
	return sb.String()
}

func tiltSouth(m [][]rune) {
	for j := 0; j < len(m[0]); j++ {
		roll := 0
		for i := len(m) - 1; i >= 0; i-- {
			if m[i][j] == '#' {
				roll = 0
			} else if m[i][j] == '.' {
				roll++
			} else { // O
				if roll > 0 {
					m[i+roll][j] = 'O'
					m[i][j] = '.'
				}
			}
		}
	}
	//printRunes(m)
}

func tiltWest(m [][]rune) {
	for i := 0; i < len(m); i++ {
		roll := 0
		for j := len(m[0]) - 1; j >= 0; j-- {
			if m[i][j] == '#' {
				roll = 0
			} else if m[i][j] == '.' {
				roll++
			} else { // O
				if roll > 0 {
					m[i][j+roll] = 'O'
					m[i][j] = '.'
				}
			}
		}
	}
	//printRunes(m)
}

func tiltEast(m [][]rune) {
	for i := 0; i < len(m); i++ {
		roll := 0
		for j := 0; j < len(m); j++ {
			if m[i][j] == '#' {
				roll = 0
			} else if m[i][j] == '.' {
				roll++
			} else { // O
				if roll > 0 {
					m[i][j-roll] = 'O'
					m[i][j] = '.'
				}
			}
		}
	}
	//printRunes(m)
}

func getLoad(m [][]rune) int {
	total := 0
	for j := 0; j < len(m[0]); j++ {
		for i := 0; i < len(m); i++ {
			if m[i][j] == 'O' {
				total += len(m) - i
			}
		}
	}
	return total
}

func printRunes(m [][]rune) {
	for _, s := range m {
		fmt.Println(string(s))
	}
	fmt.Println("")
}
