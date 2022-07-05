package stream

import (
	"strconv"
	"strings"
)

func (b *Buffer) WriteXMakeBody(key string, values ...string) {
	isNewLineKey := len(values) > 1
	if strings.HasPrefix(values[0], "wdk") {
		isNewLineKey = false
	}
	b.WriteString(key)
	b.WriteString("(")
	if isNewLineKey {
		b.NewLine()
	}
	for i, value := range values {
		if isNewLineKey {
			b.WriteString("\t")
		}
		b.WriteString(strconv.Quote(value))
		if key == "add_includedirs" {
			b.WriteString(",{public=true}") //for deps add
		}
		if i+1 < len(values) {
			b.WriteString(",")
			if isNewLineKey {
				b.NewLine()
			}
		}
	}
	if isNewLineKey {
		b.NewLine()
	}
	b.WriteString(")")
	b.NewLine()
}
