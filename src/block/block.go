package block

import "strings"

type (
	LineInfo struct {
		Line string
		Col  int
	}
	Lines []LineInfo
)

func FindAll(lines []string, with, next string) (l Lines) {
	l = make(Lines, 0)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, with) {
			col := i + 1
			if !strings.Contains(line, next) {
				l = append(l, LineInfo{Line: line, Col: col})
			} else {
				start := lines[i:]
				for j, s := range start {
					l = append(l, LineInfo{Line: s, Col: col + j})
					if !strings.Contains(s, next) {
						break
					}
				}
			}
		}
	}
	return
}
