package slag

import "unicode"

func toSname(s string) string {
	buf, pu := make([]rune, 0), false
	for i, c := range []rune(s) {
		u := unicode.IsUpper(c)
		if i > 0 && !pu && u {
			buf = append(buf, '_')
		}
		pu = u
		buf = append(buf, unicode.ToLower(c))
	}
	return string(buf)
}
