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

	for scanner.Scan() {
		// line := scanner.Text()
		// st := strings.Split(line, " ")

	}

	now := time.Now()
	one := partOne()
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo()
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne() int64 {
	return 0
}

func partTwo() int64 {
	return 0
}
