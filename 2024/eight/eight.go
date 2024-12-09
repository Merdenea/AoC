package main

import (
	"aoc/common"
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
		m = append(m, []rune(line))
	}

	now := time.Now()
	one := partOne(m)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(m)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne(m [][]rune) int64 {
	freqs := make(map[rune][]common.Pair)

	for i, l := range m {
		for j, r := range l {
			if r == '.' {
				continue
			}
			if _, ok := freqs[r]; !ok {
				freqs[r] = make([]common.Pair, 0)
			}
			freqs[r] = append(freqs[r], common.Pair{int64(i), int64(j)})
		}
	}

	pos := make(map[string]int)

	for _, v := range freqs {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				dx, dy := v[j].I-v[i].I, v[j].J-v[i].J

				if inBounds(int(v[j].I+dx), int(v[j].J+dy), m) {
					pos[key(int(v[j].I+dx), int(v[j].J+dy))]++
				}
				if inBounds(int(v[i].I-dx), int(v[i].J-dy), m) {
					pos[key(int(v[i].I-dx), int(v[i].J-dy))]++
				}
			}
		}
	}
	return int64(len(pos))
}

func inBounds(i, j int, m [][]rune) bool {
	return i >= 0 && j >= 0 && i < len(m) && j < len(m[0])
}

func key(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func partTwo(m [][]rune) int64 {
	freqs := make(map[rune][]common.Pair)

	for i, l := range m {
		for j, r := range l {
			if r == '.' {
				continue
			}
			if _, ok := freqs[r]; !ok {
				freqs[r] = make([]common.Pair, 0)
			}
			freqs[r] = append(freqs[r], common.Pair{int64(i), int64(j)})
		}
	}

	pos := make(map[string]int)

	for _, v := range freqs {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				dx, dy := int(v[j].I-v[i].I), int(v[j].J-v[i].J)

				for c := 0; inBounds(int(dx)*c, j, m) || inBounds(i, c*int(dy), m); c++ {
					if inBounds(int(v[j].I)+dx*c, int(v[j].J)+dy*c, m) {
						pos[key(int(v[j].I)+dx*c, int(v[j].J)+dy*c)]++
					}
					if inBounds(int(v[i].I)-dx*c, int(v[i].J)-dy*c, m) {
						pos[key(int(v[i].I)-dx*c, int(v[i].J)-dy*c)]++
					}
				}
			}
		}
	}

	// for i, l := range m {
	// 	for j, r := range l {
	// 		if r != '.' {
	// 			fmt.Print(string(r) + " ")
	// 		} else if pos[key(i, j)] > 0 {
	// 			fmt.Print("# ")
	// 		} else {
	// 			fmt.Print(string(r) + " ")
	// 		}
	// 	}
	// 	fmt.Println()

	// }
	return int64(len(pos))
}
