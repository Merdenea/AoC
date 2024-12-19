package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	// startTime := time.Now()
	// totalTimeSeconds := 0
	for _, e := range entries {
		os.Chdir("/Users/vladistrate/Documents/aoc/2024/" + e.Name())

		// out, _ := exec.Command("pwd").Output()
		// fmt.Println(string(out))
		fmt.Println("Day", e.Name())
		cmd := exec.Command("go", "run", e.Name()+".go", "input")
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))

	}
}
