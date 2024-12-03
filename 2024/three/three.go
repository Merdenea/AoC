package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	now := time.Now()
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), partOne(lines))

	now = time.Now()
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), partTwo(lines))

}

func parseLine(lines []string) []string {
	sts := make([]string, 0)
	re := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)|don't|do")

	for _, s := range lines {
		matches := re.FindAll([]byte(s), -1)
		for _, m := range matches {
			sts = append(sts, string(m))
		}
	}

	return sts
}

func partOne(lines []string) int64 {
	mls := parseLine(lines)
	total := int64(0)

	for _, m := range mls {
		if strings.HasPrefix(m, "mul(") {
			s := strings.Trim(m, "mul(")
			s = strings.Trim(s, ")")
			st := strings.Split(s, ",")

			one, _ := strconv.ParseInt(st[0], 10, 64)
			two, _ := strconv.ParseInt(st[1], 10, 64)

			total += one * two
		}
	}

	return int64(total)
}

func partTwo(lines []string) int64 {
	mls := parseLine(lines)
	total := int64(0)

	shouldMultiply := true
	for _, m := range mls {
		if m == "don't" {
			shouldMultiply = false
		}
		if m == "do" {
			shouldMultiply = true
		}
		if strings.HasPrefix(m, "mul(") && shouldMultiply {
			s := strings.Trim(m, "mul(")
			s = strings.Trim(s, ")")
			st := strings.Split(s, ",")

			one, _ := strconv.ParseInt(st[0], 10, 64)
			two, _ := strconv.ParseInt(st[1], 10, 64)

			total += one * two
		}
	}

	return int64(total)
}
