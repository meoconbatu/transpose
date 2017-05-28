package transpose

const testVersion = 1

func Transpose(in []string) []string {
	if len(in) == 0 {
		return make([]string, 0)
	}
	ou := make([][]rune, 0)
	lenOu := len(ou)
	for i, str := range in {
		for j, r := range str {
			if j >= lenOu {
				x := make([]rune, i+1)
				x[i] = r
				for z := 0; z < i; z++ {
					x[z] = ' '
				}
				ou = append(ou, x)
				lenOu = len(ou)
			} else {
				for z := len(ou[j]); z < i; z++ {
					ou[j] = append(ou[j], ' ')
				}
				ou[j] = append(ou[j], r)
			}

		}
	}
	output := make([]string, len(ou))
	for i, str := range ou {
		output[i] = string(str)
	}
	return output
}
