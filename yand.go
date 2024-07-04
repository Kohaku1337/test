package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Word struct {
	word       string
	index, pop int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := 3
	q := 6
	t := ""
	var qin string
	arr := make([]Word, n)
	arr[0] = Word{
		word:  "yandex",
		index: 1,
		pop:   10,
	}
	arr[1] = Word{
		word:  "yacht",
		index: 2,
		pop:   1,
	}
	arr[2] = Word{
		word:  "yoghurt",
		index: 3,
		pop:   15,
	}
	for i := 0; i < q; i++ {
		scanner.Scan()
		qin = scanner.Text()
		t = query(qin, t, arr)
	}
	autocomplete(t, arr)
	//arr[3] = "japanese goblins", 3, 20
}

func autocomplete(s string, dict []Word) int {
	maxIndex := -1
	max := 0
	for _, value := range dict {
		if strings.HasPrefix(value.word, s) {
			if value.pop > max {
				max = value.pop
				maxIndex = value.index
			}
		}
	}
	return maxIndex
}

func query(qinput, t string, dict []Word) string {
	if qinput[0] == '+' {
		t += string(qinput[2])
	} else {
		t = t[:(len(t) - 1)]
	}
	fmt.Println(autocomplete(t, dict))
	return t
}
