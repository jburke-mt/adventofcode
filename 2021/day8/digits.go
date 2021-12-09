package main

type Digit struct {
	value uint
}

func (d *Digit) IsSegmentOn(segment rune) bool {
	const runeStart = 97
	segmentAsInt := int(segment - runeStart)
	return d.value&(1<<segmentAsInt) != 0
}

//  aaa
// b   c
// b   c
//  ddd
// e   f
// e   f
//  ggg

func RuneToInt(r rune) uint {
	const runeStart = 97
	return uint(r - runeStart)
}

func SetDifference(first string, second string) string {
	m := make(map[rune]bool)
	var diff_slice []rune

	for _, r := range first {
		m[r] = true
	}

	for _, r := range second {
		if _, ok := m[r]; !ok {
			diff_slice = append(diff_slice, r)
		}
	}

	return string(diff_slice)
}

func CreateDecoding(one string, four string, seven string, eight string) map[rune]rune {
	decoding := make(map[rune]rune)
	decoding['a'] = rune(SetDifference(seven, one)[0])
	return decoding
}

// create constants corresponding to the segments for digits 0-9

// store decoding in a map (above, [d:a, a:c, b:f])

// then decode the output and match to the correct digit