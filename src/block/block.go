package block

import "strings"

type (
	LineInfo struct {
		Line string
		Col  int
	}
	Lines []LineInfo
)

func FindAll(lines []string, start, end string) (l Lines) {
	l = make(Lines, 0)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, start) {
			col := i + 1
			if strings.Contains(line, end) {
				start := lines[i:]
				for j, s := range start {
					l = append(l, LineInfo{Line: s, Col: col + j})
					if !strings.Contains(s, end) {
						break
					}
				}
			}
		}
	}
	return
}
