package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Machine struct {
	A_dx,
	A_dy,
	B_dx,
	B_dy int64

	target_x, target_y int64
}

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ms := make([]Machine, 0)
	for scanner.Scan() {
		var adx, ady, bdx, bdy, tx, ty int64
		line := scanner.Text()
		st := strings.Split(line, "Button A: ")
		sp := strings.Split(st[1], ", ")

		str1 := sp[0][2:]
		str2 := sp[1][2:]

		adx, _ = strconv.ParseInt(str1, 10, 64)
		ady, _ = strconv.ParseInt(str2, 10, 64)

		scanner.Scan()
		line = scanner.Text()
		st = strings.Split(line, "Button B: ")
		sp = strings.Split(st[1], ", ")

		str1 = sp[0][2:]
		str2 = sp[1][2:]

		bdx, _ = strconv.ParseInt(str1, 10, 64)
		bdy, _ = strconv.ParseInt(str2, 10, 64)

		scanner.Scan()
		line = scanner.Text()
		st = strings.Split(line, "Prize: ")
		sp = strings.Split(st[1], ", ")

		str1 = sp[0][2:]
		str2 = sp[1][2:]

		tx, _ = strconv.ParseInt(str1, 10, 64)
		ty, _ = strconv.ParseInt(str2, 10, 64)

		ms = append(ms, Machine{
			A_dx:     adx,
			A_dy:     ady,
			B_dx:     bdx,
			B_dy:     bdy,
			target_x: tx,
			target_y: ty,
		})
		scanner.Scan()
	}

	now := time.Now()
	one := partOne(ms)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(ms)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne(ms []Machine) int64 {
	total := 0
	costA := 3
	costB := 1
	for _, m := range ms {
		a, b := 1, 1
		for a*int(m.A_dx) < int(m.target_x) {
			b = 1
			found := false
			for a*int(m.A_dx)+b*int(m.B_dx) <= int(m.target_x) && a*int(m.A_dy)+b*int(m.B_dy) <= int(m.target_y) {
				if a*int(m.A_dx)+b*int(m.B_dx) == int(m.target_x) && a*int(m.A_dy)+b*int(m.B_dy) == int(m.target_y) {
					found = true
					total += a*costA + b*costB
					break
				}
				b++
			}
			if found {
				break
			}
			a++
		}
	}
	return int64(total)
}

func partTwo(ms []Machine) int64 {
	total := float64(0)
	costA := 3
	costB := 1

	var a, b float64

	for _, m := range ms {
		m.target_x += 10000000000000
		m.target_y += 10000000000000
		one := float64(m.A_dx*m.target_y)/float64(m.A_dy) - float64(m.target_x)
		two := float64(m.A_dx*m.B_dy)/float64(m.A_dy) - float64(m.B_dx)

		b = one / two
		a = (float64(m.target_y) - b*float64(m.B_dy)) / float64(m.A_dy)

		// fmt.Println(a, b)

		if isInt(a) && isInt(b) {
			total += a*float64(costA) + b*float64(costB)
			// fmt.Println(i)
		}
	}
	return int64(total)
}

func isInt(a float64) bool {
	epsilon := 1e-3 // Margin of error
	if _, frac := math.Modf(math.Abs(a)); frac < epsilon || frac > 1.0-epsilon {
		return true
	}
	return false
}
