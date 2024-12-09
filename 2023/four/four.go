package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var mm map[int64]int64

func main() {
	file, err := os.Open("./four/in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := int64(0)
	mm = make(map[int64]int64)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		val := getPoints(scanner.Text())
		//fmt.Println(val)
		total += val
	}
	fmt.Println(total)
	fmt.Println(mm)
	total = 0
	for _, v := range mm {
		total +=v
	}
	fmt.Println(total)

}

func getPoints(line string) int64 {
	sp1 := strings.Split(line, ":")
	cardNo := getCardNo(sp1[0])
	if cardNo == 0 {
		fmt.Println("wtf")
	}
	mm[cardNo]++
	//fmt.Println(sp1[0])
	sp2 := strings.Split(sp1[1], "|")


	wNumList := strings.Split(sp2[0], " ")
	hNumList := strings.Split(sp2[1], " ")

	wNum := make(map[int64]bool)
	for _, n := range wNumList {
		if n == "" || n == " "{
			continue
		}
		x, _ := strconv.ParseInt(n, 10, 64)
		wNum[x] = true
	}

	hNum := make([]int64, 0)
	for _, n := range hNumList {
		if n == "" || n == " "{
			continue
		}
		x, _ := strconv.ParseInt(n, 10, 64)
		hNum = append(hNum, x)
	}

	count := int64(0)
	for _, n := range hNum {
		if wNum[n] {
			count++
		}
	}

	currentCardVal := mm[cardNo]

	for i := int64(1); i <= count; i++ {
		mm[cardNo+i] += currentCardVal
	}

	if count >=1 {
		return int64(math.Pow(2,float64(count-1)))
	}
	return 0
}

func getCardNo(s string) int64 {
	i := strings.LastIndex(s, " ")
	if i != -1 {
		v, _ := strconv.ParseInt(s[i+1:], 10, 64)
		return v
	}
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}