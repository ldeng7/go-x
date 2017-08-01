package sql_builder

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	ArgTypeString = iota
	ArgTypeQuotedString
	ArgTypeStringArray
	ArgTypeQuotedStringArray
	ArgTypeMap
	ArgTypeKVArrays
	ArgTypeArrayArray
)

type Arg struct {
	Type  int
	Value interface{}
}

func Sql(sql string, args map[string]Arg) string {
	m := parseArgs(args)
	return reg.ReplaceAllStringFunc(sql, genReplFunc(m))
}

var reg *regexp.Regexp

func init() {
	reg = regexp.MustCompile("#\\{([\\w]+)\\}")
}

func quote(s string) string {
	bytes := []byte(s)
	out := make([]byte, len(bytes)*2+2)

	out[0] = '\''
	i := 1

	for _, b := range bytes {
		switch b {
		case 0:
			out[i] = '\\'
			out[i+1] = '0'
			i = i + 2
		case '\b':
			out[i] = '\\'
			out[i+1] = 'b'
			i = i + 2
		case '\n':
			out[i] = '\\'
			out[i+1] = 'n'
			i = i + 2
		case '\r':
			out[i] = '\\'
			out[i+1] = 'r'
			i = i + 2
		case '\t':
			out[i] = '\\'
			out[i+1] = 't'
			i = i + 2
		case 26:
			out[i] = '\\'
			out[i+1] = 'Z'
			i = i + 2
		case '\\':
			out[i] = '\\'
			out[i+1] = '\\'
			i = i + 2
		case '\'':
			out[i] = '\\'
			out[i+1] = '\''
			i = i + 2
		case '"':
			out[i] = '\\'
			out[i+1] = '"'
			i = i + 2
		default:
			out[i] = b
			i = i + 1
		}
	}

	out[i] = '\''
	i = i + 1

	return string(out[:i])
}

func parseArgs(args map[string]Arg) map[string]string {
	out := make(map[string]string)
	for ka, a := range args {
		switch a.Type {
		case ArgTypeString:
			s, ok := a.Value.(string)
			if ok {
				out[ka] = s
			}
		case ArgTypeQuotedString:
			s, ok := a.Value.(string)
			if ok {
				out[ka] = quote(s)
			}
		case ArgTypeStringArray:
			arr, ok := a.Value.([]string)
			if ok {
				out[ka] = strings.Join(arr, ", ")
			}
		case ArgTypeQuotedStringArray:
			arr, ok := a.Value.([]string)
			if ok {
				arrO := make([]string, len(arr))
				for i, s := range arr {
					arrO[i] = quote(s)
				}
				out[ka] = strings.Join(arrO, ", ")
			}
		case ArgTypeMap:
			m, ok := a.Value.(map[string]string)
			if ok {
				arrO := make([]string, len(m))
				i := 0
				for k, v := range m {
					arrO[i] = fmt.Sprintf("%s = %s", k, quote(v))
				}
				out[ka] = strings.Join(arrO, ", ")
			}
		case ArgTypeKVArrays:
			arr, ok := a.Value.([][]string)
			if ok && (len(arr) >= 2) {
				arrK, arrV := arr[0], arr[1]
				if len(arrV) >= len(arrK) {
					arrO := make([]string, len(arrK))
					for i, k := range arrK {
						arrO[i] = fmt.Sprintf("%s = %s", k, quote(arrV[i]))
					}
					out[ka] = strings.Join(arrO, ", ")
				}
			}
		case ArgTypeArrayArray:
			arr, ok := a.Value.([][]string)
			if ok {
				arrO := make([]string, len(arr))
				for i, arrE := range arr {
					arrOE := make([]string, len(arrE))
					for j, e := range arrE {
						arrOE[j] = quote(e)
					}
					arrO[i] = strings.Join(arrOE, ", ")
				}
				out[ka] = fmt.Sprintf("(%s)", strings.Join(arrO, "), ("))
			}
		}
	}
	return out
}

func genReplFunc(m map[string]string) func(string) string {
	return func(src string) string {
		key := src[2 : len(src)-1]
		repl, found := m[key]
		if !found {
			repl = key
		}
		return repl
	}
}
