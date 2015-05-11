package slag

import "unicode"

func toSnake(s string) string {
	t := []rune(s)
	l := len(t)
	buf := make([]rune, 0, l)
	prevLow := false
	for i, c := range t {
		u := unicode.IsUpper(c)
		if u {
			c = unicode.ToLower(c)
			if i > 0 && (prevLow || (i+1 < l && unicode.IsLower(t[i+1]))) {
				buf = append(buf, '_')
			}
		}
		prevLow = !u
		buf = append(buf, c)
	}
	return string(buf)
}
