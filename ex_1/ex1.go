package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int]string)
	i := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m[i] = scanner.Text()
		i++
	}

	for k := 1; k <= i; k++ {
		for a := k + 1; a <= i; a++ {
			if m[k] == m[a] {
				fmt.Printf("(%d / %s) - (%d / %s)\n", k, m[k], a, m[a])
			}
		}
	}

}
