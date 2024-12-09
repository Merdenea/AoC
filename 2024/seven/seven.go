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

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make(map[int64][]int64)
	for scanner.Scan() {
		line := scanner.Text()
		st := strings.Split(line, ": ")
		target, _ := strconv.ParseInt(st[0], 10, 64)
		sp := strings.Split(st[1], " ")
		nums := make([]int64, 0)
		for _, s := range sp {
			n, _ := strconv.ParseInt(s, 10, 64)
			nums = append(nums, n)
		}
		m[target] = nums
	}

	now := time.Now()
	one := partOne(m)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(m)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne(m map[int64][]int64) int64 {
	total := int64(0)
	for k, v := range m {
		if isOK(k, v[0], MUL, 1, v) || isOK(k, v[0], ADD, 1, v) {
			total += k
		}
	}
	return total
}

const (
	MUL = 1
	ADD = 2
	CON = 3
)

// var (
// 	seen map[]
// )

func isOK(target, currentVal int64, op, i int, vals []int64) bool {
	if i == len(vals) {
		// fmt.Println(target, currentVal)
		return currentVal == target
	}
	// fmt.Println(target, currentVal, vals[i])

	switch op {
	case MUL:
		currentVal *= vals[i]
	case ADD:
		currentVal += vals[i]
	}

	return isOK(target, currentVal, MUL, i+1, vals) || isOK(target, currentVal, ADD, i+1, vals)
}

func isOK2(target, currentVal int64, op, i int, vals []int64) bool {
	if i == len(vals) {
		// fmt.Println(target, currentVal)
		return currentVal == target
	}
	// fmt.Println(target, currentVal, vals[i])

	switch op {
	case MUL:
		currentVal *= vals[i]
	case ADD:
		currentVal += vals[i]
	case CON:
		currentVal = int64(math.Pow10(len(strconv.Itoa(int(vals[i])))))*int64(currentVal) + vals[i]
	}

	return isOK2(target, currentVal, MUL, i+1, vals) || isOK2(target, currentVal, ADD, i+1, vals) || isOK2(target, currentVal, CON, i+1, vals)
}

func partTwo(m map[int64][]int64) int64 {
	total := int64(0)
	for k, v := range m {
		if isOK2(k, v[0], MUL, 1, v) || isOK2(k, v[0], ADD, 1, v) || isOK2(k, v[0], CON, 1, v) {
			total += k
		}
	}
	return total
}
