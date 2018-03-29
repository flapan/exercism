//Package hamming contains a solution to the hamming exercism exercise
package hamming

// Distance finds and returns the hamming distance between two [DNA sequence] strings
func Distance(a, b string) (int, error) {
	var distance = 0
	sequenceA := []rune(a)
	sequenceB := []rune(b)
	for i := 0; i < len(a); i++ {
		if sequenceA[i] != sequenceB[i] {
			distance++
		}
	}

	return distance, nil
}
