package errutils

import (
	"strings"
)

func JoinErrs(errs []error, delim string) string {
	strs := make([]string, len(errs))
	for i, err := range errs {
		strs[i] = err.Error()
	}
	return strings.Join(strs, delim)
}
