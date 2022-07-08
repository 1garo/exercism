package scrabble

import "strings"

var possibilities = map[int]string{
	1: "A, E, I, O, U, L, N, R, S, T",
	2: "D, G",
	3: "B, C, M, P",
	4: "F, H, V, W, Y",
	5: "K",
	8: "J, X",
	10: "Q, Z",
}

func Score(word string) int {
	acc := 0
	rWord := []rune(word)

	for _, c := range rWord {
		for i, key := range possibilities {
			if strings.Contains(key, strings.ToUpper(string(c))) {
				multiplier := strings.Count(key, strings.ToUpper(string(c)))
				acc += (i * multiplier)
			}
		}
	}

	return acc
}
