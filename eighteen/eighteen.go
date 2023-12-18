package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	UP    = direction{-1, 0}
	DOWN  = direction{1, 0}
	LEFT  = direction{0, -1}
	RIGHT = direction{0, 1}
	NONE  = direction{0, 0}

	dirMap = map[string]direction{
		"U": UP,
		"D": DOWN,
		"L": LEFT,
		"R": RIGHT,

		"3": UP,
		"1": DOWN,
		"2": LEFT,
		"0": RIGHT,
	}
)

type direction struct {
	i, j int
}

type instruction struct {
	direction direction
	len       int
	rgb       string
}

func main() {
	file, _ := os.Open("./eighteen/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instructions := make([]instruction, 0)

	for scanner.Scan() {
		line := scanner.Text()
		sp := strings.Split(line, " ")
		ll, _ := strconv.ParseInt(sp[1], 10, 64)
		instructions = append(instructions, instruction{
			direction: dirMap[sp[0]],
			len:       int(ll),
			rgb:       sp[2],
		})
	}

	now := time.Now()
	fmt.Printf("Part One %d in [%s]\n", getFancyArea(instructions), time.Now().Sub(now))

	newIns := make([]instruction, 0)
	for _, i := range instructions {
		rgb := i.rgb
		//(#262512)
		rgb = rgb[2 : len(rgb)-1]
		ll, _ := strconv.ParseInt(rgb[:len(rgb)-1], 16, 64)
		in := instruction{
			direction: dirMap[rgb[len(rgb)-1:]],
			len:       int(ll),
		}
		newIns = append(newIns, in)
	}

	now = time.Now()
	fmt.Printf("Part Two %d in [%s]\n", getFancyArea(newIns), time.Now().Sub(now))
}

// Irregular polygon area because I'm lazy
func getFancyArea(instructions []instruction) int {
	i, j := 0, 0
	det := make([][]int, 0)

	det = append(det, []int{i, j})
	per := 0

	minI, minJ := 0, 0
	for _, in := range instructions {
		i, j = i+in.direction.i*in.len, j+in.direction.j*in.len
		if i < minI {
			minI = i
		}
		if j < minJ {
			minI = j
		}
		per += in.len
		det = append(det, []int{i, j})
	}

	area := 0
	minI, minJ = minI*-1, minJ*-1
	for i = 0; i < len(det)-1; i++ {
		x1, y1 := det[i][0]+minI, det[i][1]+minJ
		x2, y2 := det[i+1][0]+minI, det[i+1][1]+minJ
		/*
			sum |x1  y1|
				|x2  y2| ...
		*/
		area += x2*y1 - x1*y2
	}

	return (area+per)/2 + 1
}
