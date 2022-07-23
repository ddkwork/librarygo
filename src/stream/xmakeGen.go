package stream

import (
	"strconv"
	"strings"
)

func (b2 *Buffer) WriteXMakeBody(key string, values ...string) {
	isNewLineKey := len(values) > 1
	if strings.HasPrefix(values[0], "wdk") {
		isNewLineKey = false
	}
	b2.WriteString(key)
	b2.WriteString("(")
	if isNewLineKey {
		b2.NewLine()
	}
	for i, value := range values {
		if isNewLineKey {
			b2.WriteString("\t")
		}
		b2.WriteString(strconv.Quote(value))
		if key == "add_includedirs" {
			b2.WriteString(",{public=true}") //for deps add
		}
		if i+1 < len(values) {
			b2.WriteString(",")
			if isNewLineKey {
				b2.NewLine()
			}
		}
	}
	if isNewLineKey {
		b2.NewLine()
	}
	b2.WriteString(")")
	b2.NewLine()
}
