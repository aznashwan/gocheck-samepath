package samepath

import (
	"os"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"path/filepath"

	gc "launchpad.net/gocheck"
	jc "github.com/juju/testing/checkers"
)

type samePathChecker struct {
	*gc.CheckerInfo
}

var SamePath gc.Checker = &samePathChecker{
	&gc.CheckerInfo{Name: "SamePath", Params: []string{"obtained", "epected"}},
}

func (checker *samePathChecker) Check(params []interface{}, names []string) (result bool, error string) {
	// Check for panics
	defer func() {
		if panicked := recover(); panicked != nil {
			result = false
			error = fmt.Sprint(panicked)
		}
	}()

	// Convert inputs to strings
	obtained, isStr := jc.StringOrStringer(params[0])
	if !isStr {
		value := reflect.ValueOf(params[0])
		return false, fmt.Sprintf("obtained value is not a string and has no .String(), %s:%#v", value.Kind(), params[0])
	}
	expected, isStr := jc.StringOrStringer(params[1])
	if !isStr {
		value := reflect.ValueOf(params[1])
		return false, fmt.Sprintf("obtained value is not a string and has no .String(), %s:%#v", value.Kind(), params[1])
	}

	// Convert paths to proper format
	obtained = filepath.FromSlash(obtained)
	expected = filepath.FromSlash(expected)

	// If running on Windows, paths will be case-insensitive and thus we
	// normalize the inputs to a default of all upper-case
	if runtime.GOOS == "windows" {
		obtained = strings.ToUpper(obtained)
		expected = strings.ToUpper(expected)
	}

	// Same path does not check further is the inputs are already equal
	if obtained == expected {
		return true, ""
	}

	// If it's not the same path, check if it points to the same file.
	// Thus, the cases with windows-shortened paths are accounted for
	// This will throw an error if it's not a file
	ob, err := os.Stat(obtained)
	if os.IsNotExist(err) {
		return false, fmt.Sprintf("file %s does not exist", obtained)
	} else if err != nil {
		return false, fmt.Sprintf("other stat error: %v", err)
	}

	ex, err := os.Stat(expected)
	if os.IsNotExist(err) {
		return false, fmt.Sprintf("file %s does not exist", expected)
	} else if err != nil {
		return false, fmt.Sprintf("other stat error: %v", err)
	}

	res := os.SameFile(ob, ex)
	if res {
		return true, ""
	}
	return false, ""
}
