package cppBlock

//Warning, this package only tested with comments removed and formatted

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
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, start) {
			col := i + 1
			block := lines[i:]
			for j, s := range block {
				if s == "" {
					continue
				}
				l = append(l, LineInfo{Line: s, Col: col + j})
				if s[0] == '}' { //if not formatted it will be not right for find api end
					i += j
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
func FindMethod(lines []string) (l Lines) {
	/*
		static ZyanStatus ZydisInputPeek(ZydisDecoderContext* context,
		    ZydisDecodedInstruction* instruction, ZyanU8* value)
		{
		    ZYAN_ASSERT(context);
		    ZYAN_ASSERT(instruction);
		    ZYAN_ASSERT(value);

	*/
	start, end := `(`, `}`
	l = make(Lines, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, start) &&
			!strings.Contains(line, `#`) {
			col := i + 1
			block := lines[i:]
			isApi := false
			for _, s := range block {
				if strings.Contains(s, "{") {
					isApi = true
					break
				}
			}
			if !isApi {
				continue
			}
			for j, s := range block {
				if s == "" {
					continue
				}
				l = append(l, LineInfo{Line: s, Col: col + j})
				if s == end { //if not formatted it will be not right for find api end
					i += j
					break
				}
			}
		}
	}
	return
}
