package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]rune, 0)
	m2 := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		m = append(m, []rune(line))
		m2 = append(m2, []rune(line))
	}

	st := make([]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		st = append(st, []rune(line)...)
	}

	now := time.Now()
	one := partOne(m, st)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(m2, st)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne(m [][]rune, st []rune) int64 {
	si, sj := findStart(m)

	for _, dir := range st {
		si, sj = update(m, si, sj, dir)
		// print(m, dir)
	}

	total := 0
	for i, l := range m {
		for j, r := range l {
			if r == 'O' {
				total += 100*i + j
			}
		}
	}

	return int64(total)
}

func print(m [][]rune, dir rune) {
	fmt.Println("Dir", string(dir))

	for _, l := range m {
		for _, r := range l {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func update(m [][]rune, i, j int, dir rune) (int, int) {
	dir_i, dir_j := -1, -1
	switch dir {
	case 'v':
		dir_i, dir_j = 1, 0
	case '^':
		dir_i, dir_j = -1, 0
	case '<':
		dir_i, dir_j = 0, -1
	case '>':
		dir_i, dir_j = 0, 1
	}

	if ok, endI, endJ := canMove(m, i, j, dir_i, dir_j); ok {
		c := m[i+dir_i][j+dir_j]      // O
		m[i+dir_i][j+dir_j] = m[i][j] // @
		m[i][j] = '.'
		if i+dir_i != endI || j+dir_j != endJ {
			m[endI][endJ] = c
		}
		i, j = i+dir_i, j+dir_j
	}

	return i, j
}

func updatePart2(m [][]rune, i, j int, dir rune) (int, int) {
	dir_i, dir_j := -1, -1
	switch dir {
	case 'v':
		dir_i, dir_j = 1, 0
	case '^':
		dir_i, dir_j = -1, 0
	case '<':
		dir_i, dir_j = 0, -1
	case '>':
		dir_i, dir_j = 0, 1
	}

	if ok, endI, endJ := canMovePart2(m, i, j, dir_i, dir_j, dir); ok {
		//left/right
		for x, y := endI, endJ; x != i || y != j; {
			m[x][y], m[x-dir_i][y-dir_j] = m[x-dir_i][y-dir_j], m[x][y]
			x -= dir_i
			y -= dir_j
		}
		i, j = i+dir_i, j+dir_j
	} else {
		if i != -1 {
			return i, j
		}
	}
	return i, j
}

func canMovePart2(m [][]rune, i, j, dir_i, dir_j int, dir rune) (bool, int, int) {
	if dir != '<' && dir != '>' {
		for i >= 0 && j >= 0 && i < len(m) && j < len(m[0]) {
			if m[i][j] == '.' {
				return true, i, j
			}
			if m[i][j] == '#' {
				return false, -1, -1
			}
			if m[i][j] == '[' || m[i][j] == ']' {
				boxes = make(map[[2]int]rune)
				getBoxes(m, i, j, dir_i)
				if checkBoxes(m, dir_i) {
					moveBoxesUpDown(m, dir_i)
					m[i-dir_i][j] = '.'
					m[i][j] = '@'
				}
				return false, -1, -1
			}
			i += dir_i
			j += dir_j
		}

		return false, -1, -1
	}

	//left,right
	for i >= 0 && j >= 0 && i < len(m) && j < len(m[0]) {
		if m[i][j] == '.' {
			return true, i, j
		}
		if m[i][j] == '#' {
			return false, -1, -1
		}
		i += dir_i
		j += dir_j
	}
	return false, -1, -1
}

func moveBoxesUpDown(m [][]rune, dir_i int) {
	boxesArr := make([][2]int, 0)
	for k, _ := range boxes {
		boxesArr = append(boxesArr, k)
	}

	if dir_i == -1 {
		sort.Slice(boxesArr, func(i, j int) bool {
			return boxesArr[i][0] < boxesArr[j][0]
		})
		for _, box := range boxesArr {
			i, j := box[0], box[1]
			m[i-1][j] = m[i][j]
			m[i][j] = '.'
		}
	} else {
		sort.Slice(boxesArr, func(i, j int) bool {
			return boxesArr[i][0] > boxesArr[j][0]
		})
		for _, box := range boxesArr {
			i, j := box[0], box[1]
			m[i+1][j] = m[i][j]
			m[i][j] = '.'
		}
	}
}

var boxes = make(map[[2]int]rune)

func getBoxes(m [][]rune, i, j, dir_i int) {
	if m[i][j] == '#' || m[i][j] == '.' {
		return
	}

	if m[i][j] == ']' {
		boxes[[2]int{i, j}] = ']'
		boxes[[2]int{i, j - 1}] = '['

		getBoxes(m, i+dir_i, j, dir_i)
		getBoxes(m, i+dir_i, j-1, dir_i)
	} else if m[i][j] == '[' {
		boxes[[2]int{i, j}] = '['
		boxes[[2]int{i, j + 1}] = ']'

		getBoxes(m, i+dir_i, j, dir_i)
		getBoxes(m, i+dir_i, j+1, dir_i)
	}
}

func checkBoxes(m [][]rune, dir_i int) bool {
	for k, _ := range boxes {
		i, j := k[0], k[1]
		if m[i+dir_i][j] != '.' && m[i+dir_i][j] != '[' && m[i+dir_i][j] != ']' {
			return false
		}
	}

	return true
}

func canMove(m [][]rune, i, j, dir_i, dir_j int) (bool, int, int) {
	for i >= 0 && j >= 0 && i < len(m) && j < len(m[0]) {
		if m[i][j] == '.' {
			return true, i, j
		}
		if m[i][j] == '#' {
			return false, -1, -1
		}
		i += dir_i
		j += dir_j
	}
	return false, -1, -1
}

func findStart(m [][]rune) (int, int) {
	for i, l := range m {
		for j, r := range l {
			if r == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func partTwo(oldM [][]rune, st []rune) int64 {
	m := make([][]rune, 0)
	for _, l := range oldM {
		newL := make([]rune, 0)
		for _, r := range l {
			if r == '#' {
				newL = append(newL, '#', '#')
			}
			if r == 'O' {
				newL = append(newL, '[', ']')
			}
			if r == '.' {
				newL = append(newL, '.', '.')
			}
			if r == '@' {
				newL = append(newL, '@', '.')
			}
		}
		m = append(m, newL)
	}

	for _, dir := range st {
		si, sj := findStart(m) // fix this maybe
		updatePart2(m, si, sj, dir)
		print(m, dir)
	}

	total := 0
	for i, l := range m {
		for j, r := range l {
			if r == '[' {
				total += 100*i + j
			}
		}
	}

	return int64(total)
}
