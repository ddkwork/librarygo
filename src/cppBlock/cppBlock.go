package cppBlock

import (
	"strings"
)

type (
	LineInfo struct {
		Line string
		Col  int
	}
	Lines []LineInfo
)

func FindEnum(lines []string) (l Lines)   { return findAll(lines, `typedef enum`, "}") }
func FindStruct(lines []string) (l Lines) { return findAll(lines, `typedef struct`, "}") }
func findAll(lines []string, start string, end string) (l Lines) {
	l = make(Lines, 0)
	for i, line := range lines {
		if strings.Contains(line, start) {
			col := i + 1
			start := lines[i:]
			for j, s := range start {
				if s == "" {
					continue
				}
				l = append(l, LineInfo{Line: s, Col: col + j})
				if strings.Contains(s, end) {
					break
				}
			}
		}
	}
	return
}

func FindDefine(lines []string) (l Lines) {
	start, end := `#define`, `\`
	l = make(Lines, 0)
	for i, line := range lines {
		if strings.Contains(line, start) {
			col := i + 1
			block := lines[i:]
			for j, s := range block {
				if s == "" {
					break
				}
				l = append(l, LineInfo{Line: s, Col: col + j})
				if !strings.Contains(s, end) {
					break
				}
			}
		}
	}
	return
}

func FindExtern(lines []string) (l Lines) {
	l = make(Lines, 0)
	for i, line := range lines {
		if strings.Contains(line, `extern`) {
			col := i + 1
			if line == "" {
				continue
			}
			l = append(l, LineInfo{Line: line, Col: col})
		}
	}
	return
}
