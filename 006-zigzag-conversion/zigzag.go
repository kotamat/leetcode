package _06_zigzag_conversion

func convert(s string, numRows int) string {
	l := len(s)
	if l == 0 || numRows == 1 || l <= numRows {
		return s
	}
	roopMod := numRows * 2 - 2

	result := []byte{}
	width := l / numRows

	// pattern 0
	for j := 0; j <= width; j++ {
		if l > j * roopMod {
			result = append(result, s[j * roopMod])
		}
	}

	for i := 1; i < numRows - 1; i++ {
		for j := 0; j <= width; j++ {
			// pattern even
			if l > j * roopMod + i {
				result = append(result, s[j * roopMod + i])
				//pp.Printf("added %s\n", string(s[j*roopMod + i]))
			}
			// pattern odd
			if l > (j + 1) * roopMod - i {
				result = append(result, s[(j + 1) * roopMod - i])
				//pp.Printf("added %s\n", string(s[j*roopMod + roopMod - i]))
			}
		}
	}
	// pattern center
	for j := 0; j <= width; j++ {
		if l > j * roopMod + numRows -1 {
			result = append(result, s[j * roopMod + numRows -1])
		}
	}
	return string(result)
}
