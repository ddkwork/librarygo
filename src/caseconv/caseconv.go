package caseconv

import (
	"fmt"
	"github.com/dc0d/caseconv"
	"strings"
)

func ToCamel(s string) string {
	return fmt.Sprintf("%-50s", caseconv.ToCamel(s)) + "//" + s
}
func ToCamelUpper(s string) string {
	camel := ToCamel(s)
	camel = strings.TrimSpace(camel)
	return strings.ToUpper(string(camel[0])) + camel[1:]
}
