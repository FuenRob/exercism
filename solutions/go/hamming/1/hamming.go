package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) == len(b) {
		var count int
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count, nil
	}

	return 0, errors.New("two DNA strands are different")
}
