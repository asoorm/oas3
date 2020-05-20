package oas3

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

type PathInfo struct {
	Pkg          string     `json:"pkg"`
	Func         string     `json:"func"`
	Comment      string     `json:"comment"`
	Operation    *Operation `json:",omitempty"`
	File         string     `json:"file,omitempty"`
	Line         int        `json:"line,omitempty"`
	Anonymous    bool       `json:"anonymous,omitempty"`
	Unresolvable bool       `json:"unresolvable,omitempty"`
}

func BuildPathInfo(i interface{}) PathInfo {
	fi := PathInfo{}
	frame := getCallerFrame(i)
	goPathSrc := filepath.Join(os.Getenv("GOPATH"), "src")

	if frame == nil {
		fi.Unresolvable = true
		return fi
	}

	pkgName := getPkgName(frame.File)
	if pkgName == "chi" {
		fi.Unresolvable = true
	}
	funcPath := frame.Func.Name()

	idx := strings.Index(funcPath, "/"+pkgName)
	if idx > 0 {
		fi.Pkg = funcPath[:idx+1+len(pkgName)]
		fi.Func = funcPath[idx+2+len(pkgName):]
	} else {
		fi.Func = funcPath
	}

	if strings.Index(fi.Func, ".func") > 0 {
		fi.Anonymous = true
	}

	fi.File = frame.File
	fi.Line = frame.Line
	if filepath.HasPrefix(fi.File, goPathSrc) {
		fi.File = fi.File[len(goPathSrc)+1:]
	}

	// Check if file info is unresolvable
	if strings.Index(funcPath, pkgName) < 0 {
		fi.Unresolvable = true
	}

	if !fi.Unresolvable {
		fi.Comment = getFuncComment(frame.File, frame.Line)
		fi.Operation = parseComment(fi.Comment)
	}

	return fi
}

func getCallerFrame(i interface{}) *runtime.Frame {
	pc := reflect.ValueOf(i).Pointer()
	frames := runtime.CallersFrames([]uintptr{pc})
	if frames == nil {
		return nil
	}
	frame, _ := frames.Next()
	if frame.Entry == 0 {
		return nil
	}
	return &frame
}

func getPkgName(file string) string {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, file, nil, parser.PackageClauseOnly)
	if err != nil {
		return ""
	}
	if astFile.Name == nil {
		return ""
	}
	return astFile.Name.Name
}

func getFuncComment(file string, line int) string {
	fset := token.NewFileSet()

	astFile, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return ""
	}

	if len(astFile.Comments) == 0 {
		return ""
	}

	for _, cmt := range astFile.Comments {
		if fset.Position(cmt.End()).Line+1 == line {
			return cmt.Text()
		}
	}

	return ""
}

const (
	oas3OperationId = "oas3.operationId"
	oas3Description = "oas3.description"
	oas3Parameter   = "oas3.parameter"
	oas3Summary     = "oas3.summary"
	oas3Response   = "oas3.response"
)

func parseComment(comment string) *Operation {
	o := &Operation{}

	lines := strings.Split(comment, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, oas3OperationId) {
			o.OperationId = strings.TrimPrefix(line, oas3OperationId+": ")
			continue
		}
		if strings.HasPrefix(line, oas3Description) {
			o.Description = strings.TrimPrefix(line, oas3Description+": ")
			continue
		}
		if strings.HasPrefix(line, oas3Parameter) {
			var param Parameter

			if err := json.Unmarshal([]byte(strings.TrimPrefix(line, oas3Parameter+": ")), &param); err != nil {
				// TODO: handle error properly rather than ignoring
				println("there is an error unmarshalling")
				continue
			}

			valid := false
			for _, v := range parameterIn {
				if param.In == v {
					valid = true
					break
				}
			}
			if !valid {
				println("in param not valid")
				continue
			}

			if o.Parameters == nil {
				o.Parameters = make([]*Parameter, 0)
			}
			o.Parameters = append(o.Parameters, &param)
			continue
		}
		if strings.HasPrefix(line, oas3Response) {
			var response Response

			res := strings.TrimPrefix(line, oas3Response+".")

			re := regexp.MustCompile(`^(?P<code>(\d{3}|default)): (?P<body>.*)`)
			values := re.FindStringSubmatch(res)
			if values == nil {
				println("error with: ", res)
				continue
			}
			if len(values) != 4 {
				println("error with comments: ", res, len(values))
				for _, v := range values {
					println("\t%s", v)
				}
				continue
			}

			if err := json.Unmarshal([]byte(values[3]), &response); err != nil {
				// TODO: handle error properly rather than ignoring
				println("there is an error unmarshalling")
				continue
			}
			if o.Responses == nil {
				o.Responses = make(map[string]Response, 0)
			}
			o.Responses[values[1]] = response
			continue
		}
		if strings.HasPrefix(line, oas3Summary) {
			o.Summary = strings.TrimPrefix(line, oas3Summary+": ")
			continue
		}
	}

	return o
}
