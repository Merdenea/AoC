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

	// Regex who?
	scanner.Scan()
	line := scanner.Text()
	sp := strings.Split(line, "Register A: ")
	valA, _ = strconv.ParseInt(sp[1], 10, 64)
	A = int(valA)

	scanner.Scan()
	line = scanner.Text()
	sp = strings.Split(line, "Register B: ")
	valB, _ = strconv.ParseInt(sp[1], 10, 64)
	B = int(valB)

	scanner.Scan()
	line = scanner.Text()
	sp = strings.Split(line, "Register C: ")
	valC, _ = strconv.ParseInt(sp[1], 10, 64)
	C = int(valC)

	scanner.Scan()
	scanner.Scan()
	line = scanner.Text()
	st := strings.Split(line, "Program: ")
	program := make([]int, 0)
	for _, v := range strings.Split(st[1], ",") {
		val, _ := strconv.ParseInt(v, 10, 64)
		program = append(program, int(val))
	}
	reset()

	now := time.Now()
	one := partOne(program)
	fmt.Printf("Part I in [%s]: %s\n", time.Since(now).String(), one)
	reset()

	now = time.Now()
	two := partTwo(program)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

var (
	valA, valB, valC int64
)

func reset() {
	A = int(valA)
	B = int(valB)
	C = int(valC)
	OUT = []int{}
	INS = 0
}

var instrutions = map[int]func(int) bool{
	0: adv,
	1: bxl,
	2: bst,
	3: jnz,
	4: bxc,
	5: out,
	6: bdv,
	7: cdv,
}

var (
	A, B, C int
	INS     int
	OUT     []int
)

func combo(val int) int {
	if val <= 3 {
		return val
	}
	if val == 4 {
		return A
	}
	if val == 5 {
		return B
	}
	if val == 6 {
		return C
	}
	return -1
}

func adv(op int) bool { // 0
	res := A / int((math.Pow(2, float64(combo(op)))))
	A = res
	return true
}

func bxl(op int) bool { // 1
	B ^= op
	return true

}

func bst(op int) bool { // 2
	B = combo(op) % 8
	return true

}

func jnz(op int) bool { // 3
	if A == 0 {
		return true
	}

	INS = op
	return false
}

func bxc(op int) bool { // 4
	B ^= C
	return true
}

func out(op int) bool {
	OUT = append(OUT, combo(op)%8)
	return true
}

func bdv(op int) bool {
	res := A / int((math.Pow(2, float64(combo(op)))))
	B = res
	return true
}

func cdv(op int) bool {
	res := A / int((math.Pow(2, float64(combo(op)))))
	C = res
	return true
}

func toString(sep string) string {
	res := strings.Builder{}
	for i, v := range OUT {
		res.WriteString(strconv.Itoa(v))
		if i != len(OUT)-1 {
			res.WriteString(sep)
		}
	}
	return res.String()
}

func partOne(program []int) string {
	run(program)
	return toString(",")
}

func run(program []int) {
	for INS < len(program) {
		ins, op := program[INS], program[INS+1]
		instruction := instrutions[ins]
		if instruction(op) {
			INS += 2
		}
	}
}

/*
bst: B = A % 8
bxl: B = B ^ 6
cdv: C = A >> B
bxc: B = B ^ C
bxl: B = B ^ 4
out: -> B % 8
adv: -> A = A >> 3
jnz -> jump to start
*/

func partTwo(program []int) int64 {
	next := 0
	target := ""
	for j := len(program) - 1; j >= 0; j-- {
		target = strconv.Itoa(program[j]) + target
		// shift << 3 because the "loop" jump instructions shifts >> 3
		next = findNext(next<<3, target, program)
	}
	return int64(next)
}

func findNext(start int, target string, program []int) int {
	for i := start; ; i++ {
		reset()
		A = i
		run(program)
		out := toString("")
		if out == target {
			return i
		}
	}
}
