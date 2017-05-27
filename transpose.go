package transpose

const testVersion = 1

func Transpose(in []string) []string {
	if len(in) == 0 {
		return make([]string, 0)
	}
	ou := make([][]rune, 1000)
	for i, str := range in {
		for j, r := range str {
			ou[j][i] = r
		}
	}
	output := make([]string, len(ou))
	for i, o := range ou {
		output[i] = string(o)
	}
	return output
}
