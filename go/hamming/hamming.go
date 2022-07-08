package hamming

import "errors"

func Distance(a, b string) (int, error) {
	diff := 0

	if a == "" && b == "" {
		return 0, nil
	}

	if len(a) != len(b) {
		return 0, errors.New("cannot compare strands with not equal length")
	}


	aa := []rune(a)
	bb := []rune(b)

	for i := 0; i < len(a); i++ {
		if aa[i] != bb[i] {
			diff += 1
		}
	}

	return diff, nil
}
