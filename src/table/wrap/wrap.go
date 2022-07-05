package wrap

import "strings"

func wrap(s string, n int, b *strings.Builder) {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		wrap(s[:i], n, b)
		b.WriteByte('\n')
		wrap(s[i+1:], n, b)
		return
	}

	if len(s) > n {
		b.WriteString(s[:n])
		b.WriteByte('\n')
		wrap(s[n:], n, b)
		return
	}

	b.WriteString(s)
}

// Wrap wraps the given string forcefully with the given length.
func Wrap(s string, n int) string {
	b := strings.Builder{}
	wrap(s, n, &b)
	return b.String()
}

// Force wraps the given string without checking for newline characters.
func Force(s string, n int) string {
	b := strings.Builder{}
	i := 0
	for i+n < len(s) {
		b.WriteString(s[i : i+n])
		b.WriteByte('\n')
		i += n
	}
	b.WriteString(s[i:])
	return b.String()
}

func at(s string, c byte, n int, b *strings.Builder) {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		at(s[:i], c, n, b)
		b.WriteByte('\n')
		at(s[i+1:], c, n, b)
		return
	}

	if len(s) > n {
		if i := strings.LastIndexByte(s[:n], c); 0 <= i && i <= n {
			b.WriteString(s[:i])
			b.WriteByte('\n')
			at(s[i+1:], c, n, b)
			return
		}

		if i := strings.IndexByte(s, c); i >= 0 {
			b.WriteString(s[:i])
			b.WriteByte('\n')
			at(s[i+1:], c, n, b)
			return
		}
	}

	b.WriteString(s)
}

// At wraps the given string at the given byte. If possible, the string will be
// wrapped at the first appearance of the given byte before the given length.
// If the given byte only appears after the given length, the string will be
// wrapped at the first appearance. The string will be unaltered otherwise.
func At(s string, c byte, n int) string {
	b := strings.Builder{}
	at(s, c, n, &b)
	return b.String()
}

// Space wraps the given string at a whitespace character.
func Space(s string, n int) string { return At(s, ' ', n) }
