package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type hailstone struct {
	x0, y0, z0 int64
	vx, vy, vz int64

	line line
}

func (h hailstone) Before(x float64, y float64) bool {
	if h.vx < 0 && x > float64(h.x0) {
		return false
	}
	if h.vx > 0 && x < float64(h.x0) {
		return false
	}
	if h.vy > 0 && y < float64(h.y0) {
		return false
	}
	if h.vy < 0 && y > float64(h.y0) {
		return false
	}
	return true
}

func main() {
	file, _ := os.Open("./twentyfour/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	hailstones := make([]*hailstone, 0)
	for scanner.Scan() {
		l := scanner.Text()
		hailstones = append(hailstones, getHs(l))
	}

	now := time.Now()
	fmt.Printf("Part One %d in [%s]\n", getIntersections(hailstones, 200000000000000, 200000000000000, 400000000000000, 400000000000000), time.Now().Sub(now))
	preintEqs(hailstones)
	// part two see solve.py
}

func preintEqs(hss []*hailstone) {
	for i, h := range hss {
		if i == 4 {
			break
		}
		fmt.Printf("(X - %d)*(%d - Vy) = (%d - Vx)*(Y - %d),\n", h.x0, h.vy, h.vx, h.y0)
		fmt.Printf("(X - %d)*(%d - Vz) = (%d - Vx)*(Z - %d),\n", h.x0, h.z0, h.vx, h.z0)
	}
}

func getIntersections(hs []*hailstone, minX, minY, maxX, maxY float64) int {
	count := 0
	for i := 0; i < len(hs)-1; i++ {
		for j := i + 1; j < len(hs); j++ {
			hailOne, hailTwo := hs[i], hs[j]
			lineOne, lineTwo := hs[i].line, hs[j].line
			x, y, _ := lineOne.intersectsXY(lineTwo)
			//res := "nil"
			if x >= minX && x <= maxX && y <= maxY && y >= minY {
				// check intersection after both starting points
				if hailOne.Before(x, y) && hailTwo.Before(x, y) {
					count++
					//fmt.Printf("line1: %f*x + %f;  line2: %f*x + %f\n", lineOne.m, lineOne.n, lineTwo.m, lineTwo.n)
					//fmt.Printf("start1 (%d, %d) intersects at (%f, %f)\n", hailOne.x0, hailOne.y0, x, y)
					//fmt.Printf("start2 (%d, %d) intersects at (%f, %f)\n", hailTwo.x0, hailTwo.y0, x, y)
					//fmt.Println()
					//res = "inside"
				} else {
					//res = "before"
				}
			} else {
				//res = "outside"
			}
			//fmt.Printf("Line [%d] and line [%d] interesct at (%f, %f) -- %s\n", i+1, j+1, x, y, res)
		}
	}

	return count
}

func getHs(l string) *hailstone {
	sp := strings.Split(l, " @ ")

	pos := strings.Split(sp[0], ", ")
	x, _ := strconv.ParseInt(strings.ReplaceAll(pos[0], " ", ""), 10, 64)
	y, _ := strconv.ParseInt(strings.ReplaceAll(pos[1], " ", ""), 10, 64)
	z, _ := strconv.ParseInt(strings.ReplaceAll(pos[2], " ", ""), 10, 64)

	velos := strings.Split(sp[1], ", ")
	vx, _ := strconv.ParseInt(strings.ReplaceAll(velos[0], " ", ""), 10, 64)
	vy, _ := strconv.ParseInt(strings.ReplaceAll(velos[1], " ", ""), 10, 64)
	vz, _ := strconv.ParseInt(strings.ReplaceAll(velos[2], " ", ""), 10, 64)

	hs := &hailstone{
		x0:   x,
		y0:   y,
		z0:   z,
		vx:   vx,
		vy:   vy,
		vz:   vz,
		line: getXYLinearFit(x, y, vx, vy),
	}
	return hs
}

type line struct {
	lineType string
	m, n     float64 /// (y = mx +n)
	vx, vy   int
	y, x     int
}

func (l line) intersectsXY(l2 line) (float64, float64, float64) {
	if l.lineType == "x" && l2.lineType == "x" {
		log.Fatal("not handled")
	}
	if l.lineType == "y" && l2.lineType == "y" {
		log.Fatal("not handled")
	}
	if l.lineType == "xy" && l2.lineType == "xy" {
		// solve for xy
		// check if start pos is before

		x := (l2.n - l.n) / (l.m - l2.m)
		y := x*l.m + l.n
		//t := (x-)
		return x, y, 0
	}

	xyLine, constLine := l2, l
	if l.lineType == "xy" {
		xyLine = l
		constLine = l2
	}

	switch constLine.lineType {
	case "x":
		x := float64(constLine.x)
		y := xyLine.m*x + xyLine.n
		return x, y, 0
	case "y":
		y := float64(constLine.y)
		x := (y - xyLine.n) / xyLine.m
		return x, y, 0
	}

	return 0, 0, 0
}

func getXYLinearFit(x0 int64, y0 int64, vx int64, vy int64) line {
	if vx == 0 {
		return line{
			lineType: "x",
			x:        int(x0),
			vy:       int(vy),
		}
	}
	if vy == 0 {
		return line{
			lineType: "y",
			y:        int(y0),
			vx:       int(vx),
		}
	}
	x1, y1 := x0+vx, y0+vy
	m := float64(y0-y1) / float64(x0-x1)
	n := float64(y0) - m*float64(x0)
	return line{
		lineType: "xy",
		m:        m,
		n:        n,
	}
}
