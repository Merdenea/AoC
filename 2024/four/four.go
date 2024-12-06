package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		l := make([]rune, 0)
		for _, r := range line {
			l = append(l, r)
		}
		m = append(m, l)
	}

	now := time.Now()
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), partOne(m))

	now = time.Now()
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), partTwo(m))

}

func partOne(m [][]rune) int64 {
	total := int64(0)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 'X' {
				total += checkX(i, j, m)
			}
		}
	}

	return total
}

func checkX(i, j int, m [][]rune) int64 {
	total := int64(0)
	//RIGHT
	if j+3 < len(m[0]) {
		if m[i][j+1] == 'M' && m[i][j+2] == 'A' && m[i][j+3] == 'S' {
			total++
		}
	}

	//LEFFT
	if j-3 >= 0 {
		if m[i][j-1] == 'M' && m[i][j-2] == 'A' && m[i][j-3] == 'S' {
			total++
		}
	}

	//UP
	if i-3 >= 0 {
		if m[i-1][j] == 'M' && m[i-2][j] == 'A' && m[i-3][j] == 'S' {
			total++
		}
	}
	//DOWN
	if i+3 < len(m) {
		if m[i+1][j] == 'M' && m[i+2][j] == 'A' && m[i+3][j] == 'S' {
			total++
		}
	}

	//DIAG 1
	if i+3 < len(m) && j+3 < len(m[0]) {
		if m[i+1][j+1] == 'M' && m[i+2][j+2] == 'A' && m[i+3][j+3] == 'S' {
			total++
		}
	}
	//DIAG 2
	if i+3 < len(m) && j-3 >= 0 {
		if m[i+1][j-1] == 'M' && m[i+2][j-2] == 'A' && m[i+3][j-3] == 'S' {
			total++
		}
	}
	//DIAG 3
	if i-3 >= 0 && j-3 >= 0 {
		if m[i-1][j-1] == 'M' && m[i-2][j-2] == 'A' && m[i-3][j-3] == 'S' {
			total++
		}
	}
	//DIAG 4
	if i-3 >= 0 && j+3 < len(m[0]) {
		if m[i-1][j+1] == 'M' && m[i-2][j+2] == 'A' && m[i-3][j+3] == 'S' {
			total++
		}
	}
	return total
}

func checkA(i, j int, m [][]rune) int64 {
	total := int64(0)

	if i-1 < 0 || j-1 < 0 || i+1 >= len(m) || j+1 >= len(m[0]) {
		return 0
	}

	if ((m[i-1][j-1] == 'M' && m[i+1][j+1] == 'S') || (m[i-1][j-1] == 'S' && m[i+1][j+1] == 'M')) && ((m[i-1][j+1] == 'M' && m[i+1][j-1] == 'S') || (m[i-1][j+1] == 'S' && m[i+1][j-1] == 'M')) {
		total++
	}
	return total
}

func partTwo(m [][]rune) int64 {
	total := int64(0)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 'A' {
				total += checkA(i, j, m)
			}
		}
	}

	return total
}
