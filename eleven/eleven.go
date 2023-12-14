package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./eleven/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]rune, 0)
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, []rune(l))
		//if isEmptyRow(l) {
		// m = append(m, []rune(l))
		//}
	}

	//for j := 0; j < len(m[j]); j++ {
	// colEmpty := true
	// for i := 0; i < len(m) && colEmpty; i++ {
	//    if m[i][j] == '#' {
	//       colEmpty = false
	//    }
	// }
	// if colEmpty {
	//    emptyCols[fmt.Sprintf("%d", j)] = true
	// }
	//}
	//
	//for i := 0; i < len(m); i++ {
	// fm = append(fm, make([]rune, len(m[i])+len(emptyCols)))
	//}
	//
	//colCount := 0
	//for j := 0; j < len(m[j-colCount])+len(emptyCols); j++ {
	// for i := 0; i < len(m); i++ {
	//    fm[i][j] = m[i][j-colCount]
	// }
	// if emptyCols[fmt.Sprintf("%d", j-colCount)] {
	//    for i := 0; i < len(m); i++ {
	//       fm[i][j+1] = m[i][j-colCount]
	//    }
	//    colCount++
	//    j++
	// }
	//}

	emptyCols := make(map[string]bool)
	emptyRows := make(map[string]bool)
	gs := make([]pair, 0)
	n := 1
	//galaxyPair := make(map[string]bool)
	for i := 0; i < len(m); i++ {
		rowEmpty := true
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == '#' {
				rowEmpty = false
				gs = append(gs, pair{n, i, j})
				n++
			}
		}
		if rowEmpty {
			emptyRows[fmt.Sprintf("%d", i)] = true
		}
	}

	for j := 0; j < len(m[0]); j++ {
		colEmpty := true
		for i := 0; i < len(m) && colEmpty; i++ {
			if m[i][j] == '#' {
				colEmpty = false
			}
		}
		if colEmpty {
			emptyCols[fmt.Sprintf("%d", j)] = true
		}
	}

	fmt.Println(emptyRows)
	fmt.Println(emptyCols)

	total := 0
	factor := 1000000

	for i := 0; i < len(gs); i++ {
		for j := i + 1; j < len(gs); j++ {
			iDist := int(math.Abs(float64(gs[j].a-gs[i].a))) + crosses(emptyRows, gs[j].a, gs[i].a)*(factor-1)
			jDist := int(math.Abs(float64(gs[j].b-gs[i].b))) + crosses(emptyCols, gs[j].b, gs[i].b)*(factor-1)
			total += iDist + jDist
		}
	}

	fmt.Println(total)

	//printRunes(fm)
}

func crosses(emptyR map[string]bool, a, b int) int {
	maxA := int64(max(a, b))
	minA := int64(min(a, b))

	count := 0
	for k, _ := range emptyR {
		i, _ := strconv.ParseInt(k, 10, 64)
		if i > minA && i < maxA {
			count++
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getKey(i, j int) string {
	if i < j {
		return fmt.Sprintf("%d+%d", i, j)
	}
	return fmt.Sprintf("%d+%d", j, i)
}

type pair struct {
	galaxyN int

	a, b int
}

func printRunes(rr [][]rune) {
	for i := 0; i < len(rr); i++ {
		for j := 0; j < len(rr[i]); j++ {
			fmt.Print(string(rr[i][j]))
		}
		fmt.Println("")
	}
}

func isEmptyRow(s string) bool {
	return !strings.Contains(s, "#")
}
