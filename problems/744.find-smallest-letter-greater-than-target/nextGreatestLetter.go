package nextGreatestLetter

// Given a characters array letters that is sorted in non-decreasing order and a character target,
// return the smallest character in the array that is larger than target.

// Note that the letters wrap around.
// For example, if target == 'z' and letters == ['a', 'b'], the answer is 'a'.

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, b := range letters {
		if b > target {
			return b
		}
	}
	return letters[0]
}
