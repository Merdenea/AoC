package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

type gear struct {
	one int
	two int
	total int
}

var gears = make(map[string]gear)

func main() {
	file, err := os.Open("./three/in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mm := make([][]rune, 0)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		mm = append(mm, []rune(scanner.Text()))
	}

	start, end := -1, -1

	sum := int64(0)

	for i := 0; i < len(mm); i++ {
		for j :=0 ; j < len(mm[i]); j++ {
			if unicode.IsDigit(mm[i][j]) {
				if start == -1 {
					start = j
					end = j
				} else {
					end = j
				}
			} else {
				sum += check(mm, i, start, end)
				start, end = -1, -1
			}

			if j == len(mm[i])-1 {
				sum += check(mm, i, start, end)
				start, end = -1, -1
			}
		}
	}

	fmt.Println(sum)
	sum = 0
	for _, v := range gears{
		if v.total == 2 {
			sum += int64(v.one * v.two)
		}
	}
	fmt.Println(sum)

	for k, v := range gears {
		if v.total >2 {
			fmt.Println(k,v)
		}
	}
}


func check(mm [][]rune, index, start, end int) int64{
	if start == -1 || end == -1 {
		return 0
	}
	// check left
	if start -1 >= 0 {
		if isSymbol(mm[index][start-1], index, start-1, int(getNum(mm[index], start, end))) {
			return getNum(mm[index], start, end)
		}
	}
	// check right
	if end + 1 < len(mm[index]) {
		if isSymbol(mm[index][end+1], index, end+1, int(getNum(mm[index], start, end))) {
			return getNum(mm[index], start, end)
		}
	}

	//above
	if checkLine(mm, index -1, start -1, end+1, index, start, end) {
		return getNum(mm[index], start, end)
	}
	//below
	if checkLine(mm, index +1, start -1, end+1, index, start, end) {
		return getNum(mm[index], start, end)
	}
	return 0
}

func checkLine(mm [][]rune, i, start, end, originalIndex, oS, oE int) bool{
	if i < 0 || i >= len(mm){
		return false
	}

	for j := start; j <= end; j ++ {
		if j < 0 {continue}
		if j >= len(mm[i]) {continue}

		if isSymbol(mm[i][j], i, j, int(getNum(mm[originalIndex], oS, oE))) {
			return true
		}
	}

	return false
}

func isSymbol(r rune, i, j, number int)bool {
	key := fmt.Sprintf("%d+%d", i, j)
	if !unicode.IsDigit(r) && r != '.' {
		if r == '*' {
			current := gears[key]
			switch current.total {
			case 0:
				current.one = number
				current.total=1
				gears[key] = current
				break
			case 1:
				current.two = number
				current.total = 2
				gears[key] = current
				break
			default:
				current.total = current.total+1
				gears[key] = current
			}
		}
	}
	return false
}

func getNum(line []rune, start, end int) int64 {
	val :=int64(0)
	pow :=float64(0)
	for i := end; i >= start; i-- {
		v,_ :=  strconv.ParseInt(string(line[i]), 10, 64)
		val += v * int64(math.Pow(10, pow))
		pow++
	}
	return val
}




