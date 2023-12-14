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
	file, err := os.Open("./two/in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := int64(0)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		sum += validateGame(scanner.Text(), 12,13,14)
	}
	fmt.Println(sum)
}

func validateGame(line string, maxR, maxG, maxB int64) int64{
	split := strings.Split(line, ":")
	//id,_ := strconv.ParseInt(split[0][5:], 10, 64)

	play := strings.Split(split[1], ";")

	minR, minG, minB := int64(0), int64(0),int64(0)

	for _, p := range play {
		r,g,b := getRGB(p)
		//fmt.Println(r,g,b)

		// Part One

		//if r > maxR || g > maxG || b > maxB {
		//	fmt.Println("invalid: ", id)
		//	return 0
		//}

		// Part Two

		minR, minG, minB = ensureMax(minR, minG, minB, r,g,b)
	}
	//fmt.Println("val: ", minR * minG *minB)
	return minR * minG *minB
}

func ensureMax(r int64, g int64, b int64, r2 int64, g2 int64, b2 int64) (int64, int64, int64) {
	r = max(r, r2)
	g = max(g, g2)
	b = max(b, b2)
	return r,g,b
}

func max(a,b int64)int64 {
	if a >b {
		return a
	}
	return b
}

func getRGB(p string) (int64, int64, int64){
	red := "red"
	green := "green"
	blue := "blue"

	r,g,b :=int64(0), int64(0), int64(0)
	split := strings.Split(p, ",")
	for _, s := range split {
		if i := strings.Index(s, red); i !=-1 {
			r,_ = strconv.ParseInt(s[1:i-1], 10, 64)
		} else if j := strings.Index(s, green); j !=-1 {
			g,_ = strconv.ParseInt(s[1:j-1], 10, 64)
		} else if k := strings.Index(s, blue); k !=-1 {
			b,_ = strconv.ParseInt(s[1:k-1], 10, 64)
		}
	}
	return r,g,b
}