package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

//Put all character in alphabetic order in a string
func stringtoalpha(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

//Compare all strings inside a file and look for anagrams
func comparestrings(file io.Reader) {
	m := make(map[int]string)
	i := 1
	var tab []string

	scanner := bufio.NewScanner(file)
	tab = append(tab, "")
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
		m[i] = stringtoalpha(scanner.Text())
		i++
	}

	for k := 1; k <= i; k++ {
		for a := k + 1; a <= i; a++ {
			if strings.Compare(m[k], m[a]) == 0 {
				fmt.Printf("(%d / %s) - (%d / %s)\n", k, tab[k], a, tab[a])
			}
		}
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	comparestrings(file)
}
