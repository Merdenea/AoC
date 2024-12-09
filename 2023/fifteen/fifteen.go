package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./fifteen/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	l := scanner.Text()
	sp := strings.Split(l, ",")
	// Part One
	//total := 0
	//for _, s := range sp {
	//	total += getHash(s)
	//}
	//fmt.Println(total)

	re := regexp.MustCompile("-|=")

	mm := make(map[int][]*lens)
	for _, s := range sp {
		//fmt.Println(s)
		step := re.Split(s, -1)
		label := step[0]
		boxNo := getHash(label)
		if strings.Contains(s, "=") {
			cp, _ := strconv.ParseInt(step[1], 10, 64)
			if v, ok := mm[boxNo]; ok {
				mm[boxNo] = updateLenses(v, &lens{label, cp})
			} else {
				newL := make([]*lens, 0)
				newL = append(newL, &lens{label, cp})
				mm[boxNo] = newL
			}
		} else {
			if v, ok := mm[boxNo]; ok {
				mm[boxNo] = removeLens(v, label)
			}
		}

		//printMap(mm)
	}

	total := 0
	for k, v := range mm {
		total += getBoxPower(k, v)
	}
	fmt.Println(total)

}

func printMap(mm map[int][]*lens) {
	for k, v := range mm {
		fmt.Printf("Box [%d]: ", k)
		for _, i := range v {
			fmt.Printf("[%s %d]", i.label, i.focalL)
		}
		fmt.Println()
	}
	fmt.Println()

}

func getBoxPower(k int, ll []*lens) int {
	p := 0
	for i, l := range ll {
		p += (k + 1) * (i + 1) * int(l.focalL)
	}
	return p
}

func removeLens(lenses []*lens, label string) []*lens {
	index := -1
	for i, l := range lenses {
		if l.label == label {
			index = i
			break
		}
	}
	if index == -1 {
		return lenses
	}
	newL := make([]*lens, 0, len(lenses)-1)
	newL = append(lenses[:index], lenses[index+1:]...)
	return newL
}

func updateLenses(lenses []*lens, newL *lens) []*lens {
	for _, l := range lenses {
		if l.label == newL.label {
			l.focalL = newL.focalL
			return lenses
		}
	}
	return append(lenses, newL)
}

type lens struct {
	label  string
	focalL int64
}

func getHash(s string) int {
	c := 0
	for _, r := range []rune(s) {
		c += int(r)
		c *= 17
		c %= 256
	}
	return c
}
