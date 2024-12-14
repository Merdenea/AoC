package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Robot struct {
	px, py, vx, vy int64
}

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	robots := make([]Robot, 0)
	for scanner.Scan() {
		line := scanner.Text()
		st := strings.Split(line, " ")

		sp := strings.Split(strings.Trim(st[0], "p="), ",")
		px, _ := strconv.ParseInt(sp[0], 10, 64)
		py, _ := strconv.ParseInt(sp[1], 10, 64)

		sp = strings.Split(strings.Trim(st[1], "v="), ",")
		vx, _ := strconv.ParseInt(sp[0], 10, 64)
		vy, _ := strconv.ParseInt(sp[1], 10, 64)

		robots = append(robots, Robot{px, py, vx, vy})
	}

	fmt.Println(robots)

	now := time.Now()
	one := partOne(robots)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(robots)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne(robots []Robot) int64 {
	bound_x, bound_y := int64(101), int64(103)
	simTime := int64(8000)

	q1, q2, q3, q4 := 0, 0, 0, 0

	for i := 0; i < int(simTime); i++ {
		printRobots(i, robots)
		for j, r := range robots {
			nx := (r.vx + r.px)
			ny := (r.vy + r.py)

			if nx < 0 {
				nx += bound_x
			} else if nx >= bound_x {
				nx = nx % bound_x
			}
			if ny < 0 {
				ny += bound_y
			} else if ny >= bound_y {
				ny = ny % bound_y
			}
			robots[j].px = nx
			robots[j].py = ny
		}
	}

	for _, r := range robots {
		nx, ny := r.px, r.py
		if nx < bound_x/2 && ny < bound_y/2 {
			q1++
		}
		if nx > bound_x/2 && ny < bound_y/2 {
			q2++
		}
		if nx < bound_x/2 && ny > bound_y/2 {
			q3++
		}
		if nx > bound_x/2 && ny > bound_y/2 {
			q4++
		}
	}

	return int64(q1 * q2 * q3 * q4)
}

func printRobots(itter int, r []Robot) {
	rmap := make(map[[2]int]int)
	for _, rr := range r {
		rmap[[2]int{int(rr.px), int(rr.py)}]++
	}

	density := false
	for i := 0; i < 101; i++ {
		if density {
			break
		}
		for j := 0; j < 103; j++ {
			if rmap[[2]int{i, j}] > 0 {
				if i-1 >= 0 && j-1 >= 0 && i+1 < 101 && j+1 < 103 {
					if rmap[[2]int{i - 1, j}] > 0 && rmap[[2]int{i + 1, j}] > 0 && rmap[[2]int{i, j + 1}] > 0 && rmap[[2]int{i, j - 1}] > 0 {
						density = true
						break
					}
				}
			}
		}
	}

	if !density {
		return
	}
	fmt.Println("itter: ", itter)

	for i := 0; i < 101; i++ {
		for j := 0; j < 103; j++ {
			if rmap[[2]int{i, j}] > 0 {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func partTwo(robots []Robot) int64 {
	return 0
}
