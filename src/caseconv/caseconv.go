package caseconv

import (
	"fmt"
	"github.com/dc0d/caseconv"
	"strings"
)

func ToCamel(data string, isCommit bool) string {
	s := fmt.Sprintf("%-50s", caseconv.ToCamel(data))
	if isCommit {
		s += "//" + s
	}
	return s
}
func ToCamelUpper(s string, isCommit bool) string {
	camel := ToCamel(s, isCommit)
	camel = strings.TrimSpace(camel)
	return strings.ToUpper(string(camel[0])) + camel[1:]
}
