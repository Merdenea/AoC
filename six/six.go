package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./six/in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	/* PartOne
	scanner.Scan()
	timeL := scanner.Text()
	scanner.Scan()
	distanceL := scanner.Text()

	sp1 := strings.Split(timeL, ": ")
	sp2 := strings.Split(distanceL, ": ")

	sp11 := strings.Split(sp1[1], " ")
	sp21 := strings.Split(sp2[1], " ")

	time := make([]int64, 0)
	distance := make([]int64, 0)
	for i := 0; i < len(sp11); i++ {
		if sp11[i] != "" {
			time = append(time, getNumber(sp11[i]))
		}
	}

	for i := 0; i < len(sp21); i++ {
		if sp21[i] != "" {
			distance = append(distance, getNumber(sp21[i]))
		}
	}

	fmt.Println(time)
	fmt.Println(distance)

	total := int64(1)
	for i := 0; i < len(time); i++ {
		total *= getWins(time[i], distance[i])
	}
	fmt.Println(total)

	*/

	/* Part Two */

	scanner.Scan()
	timeL := scanner.Text()
	scanner.Scan()
	distanceL := scanner.Text()

	sp1 := strings.Split(timeL, ": ")
	sp2 := strings.Split(distanceL, ": ")

	n1 := strings.Join(strings.Fields(sp1[1]), "")
	n2 := strings.Join(strings.Fields(sp2[1]), "")

	time := getNumber(n1)
	distance := getNumber(n2)
	fmt.Println(getWins(time, distance))
}

func getWins(time, distance int64) int64 {
	count := int64(0)
	for i := int64(0); i <= time; i++ {
		cd := (time - i) * i
		if cd > distance {
			count++
		}
	}
	//fmt.Println(count)
	return count
}

func getNumber(s string) int64 {
	i := strings.LastIndex(s, " ")
	v, _ := strconv.ParseInt(s[i+1:], 10, 64)
	return v
}
