package sqlx

import (
	"fmt"
	"regexp"
	"strings"
)

type Quote string

func Sql(sql string, args map[string]interface{}) string {
	m := parseArgs(args)
	return reg.ReplaceAllStringFunc(sql, genReplFunc(m))
}

var reg *regexp.Regexp

func init() {
	reg = regexp.MustCompile("#\\{([\\w]+)\\}")
}

func quote(s Quote) string {
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

func parseArgs(args map[string]interface{}) map[string]string {
	out := make(map[string]string)
	for ka, a := range args {
		switch arg := a.(type) {
		case string:
			out[ka] = arg
		case Quote:
			out[ka] = quote(arg)
		case []string:
			out[ka] = strings.Join(arg, ", ")
		case []Quote:
			arr := make([]string, len(arg))
			for i, s := range arg {
				arr[i] = quote(s)
			}
			out[ka] = strings.Join(arr, ", ")
		case map[string]Quote:
			arr := make([]string, len(arg))
			i := 0
			for k, v := range arg {
				arr[i] = fmt.Sprintf("%s = %s", k, quote(v))
				i++
			}
			out[ka] = strings.Join(arr, ", ")
		case [][]Quote:
			arr := make([]string, len(arg))
			for i, arrE := range arg {
				arrOE := make([]string, len(arrE))
				for j, e := range arrE {
					arrOE[j] = quote(e)
				}
				arr[i] = strings.Join(arrOE, ", ")
			}
			out[ka] = fmt.Sprintf("(%s)", strings.Join(arr, "), ("))
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
