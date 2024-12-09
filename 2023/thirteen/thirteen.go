package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./thirteen/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	m := make([][]rune, 0)
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			total += getReflections(m)
			m = make([][]rune, 0)
		} else {
			m = append(m, []rune(l))
		}
	}
	total += getReflections(m)
	fmt.Println(total)
}

func getReflections(m [][]rune) int {
	//part 1 flipped=true
	v := getVertical(m, false)
	h := 100 * getHorizontal(m, false)

	if h > 0 && v > 0 {
		printRunes(m)
	}

	return h + v
}

func printRunes(m [][]rune) {
	for _, s := range m {
		fmt.Println(string(s))
	}
	fmt.Println("")
}

func getVertical(m [][]rune, flipped bool) int {
	ffS := flipped
	for i := 0; i < len(m[0])-1; i++ {
		j := i + 1
		if ok, ff := collEq(m, i, j, ffS); ok {
			if okk, f := checkColl(m, i-1, j+1, ff); okk && f {
				return i + 1
			}
		}
	}
	return 0
}

func collEq(m [][]rune, i, j int, flipped bool) (bool, bool) {
	diff := 0
	res := true
	for k := 0; k < len(m) && diff < 2; k++ {
		if m[k][i] != m[k][j] {
			res = false
			diff++
		}
	}

	if res {
		return true, flipped
	}
	if diff == 1 && !flipped {
		return true, true
	}
	return false, flipped
}

func getHorizontal(m [][]rune, flipped bool) int {
	ffS := flipped
	for i := 0; i < len(m)-1; i++ {
		j := i + 1
		if okk, ff := compLine(m[i], m[j], ffS); okk {
			if ok, f := checkLine(m, i-1, j+1, ff); ok && f {
				return i + 1
			}
		}
	}
	return 0
}

func checkColl(m [][]rune, i int, j int, flipped bool) (bool, bool) {
	if i < 0 || j >= len(m[0]) {
		return true, flipped
	}
	if ok, ff := collEq(m, i, j, flipped); ok {
		return checkColl(m, i-1, j+1, ff)
	}
	return false, flipped
}

func checkLine(m [][]rune, i int, j int, flipped bool) (bool, bool) {
	if i < 0 || j >= len(m) {
		return true, flipped
	}
	if ok, ff := compLine(m[i], m[j], flipped); ok {
		return checkLine(m, i-1, j+1, ff)
	}
	return false, false
}

func compLine(one, two []rune, flipped bool) (bool, bool) {
	diff := 0
	res := true
	for i := 0; i < len(one) && diff < 2; i++ {
		if one[i] != two[i] {
			res = false
			diff++
		}
	}

	if res {
		return true, flipped
	}
	if diff == 1 && !flipped {
		return true, true
	}
	return false, flipped
}
