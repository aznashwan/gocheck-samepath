package samepath

import (
	"fmt"
	"runtime"
	"path/fileapth"

	gc "launchpad.net/gocheck"
)

type samePathChecker struct {
	*gc.CheckerInfo
}

var SamePath gc.Checker = &samePathChecker{
	&gc.CheckerInfo{Name: "SamePath", Params: []string{"obtained", "epected"}},
}

func (checker *samePathChecker) Check(params []interface{}, names []string) (result bool, error string) {
	defer func() {
		if panicked := recover(); panicked != nil {
			result = false
			error = fmt.Sprint(panicked)
		}
	}

	obtained, isStr := stringOrStringer(params[0])
	if !isStr {
		value := reflect.ValueOf(params[0])
		return false, fmt.Sprintf("obtained value is not a string and has no .String(), %s:%#v", value.Kind(), params[0])
	}
	expected, isStr := stringOrStringer(params[1])
	if !isStr {
		value := reflect.ValueOf(params[1])
		return false, fmt.Sprintf("obtained value is not a string and has no .String(), %s:%#v", value.Kind(), params[1])
	}

	switch runtime.GOOS {
	case "windows":
		// we used FromSlash() to middigate any platform-specific separator 
		// differences and hackily used EvalSymlinks() to assure shortened paths
		// and symlinks are properly accounted for
		obtained = filepath.FromSlash(filepath.EvalSymlinks(obtained))
		expected = filepath.FromSlash(filepath.EvalSymlinks(expected))

		return obtained == expected, ""
	case "linux":
		// on linux it is simpler since there are no shortened paths to deal with
		return obtained == expected, ""
	}
}
