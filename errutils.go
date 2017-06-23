package errutils

import (
	"errors"
	"strings"
)

func JoinErrs(delim string, errs... error) error {
	return errors.New(JoinErrsStr(delim, errs...))
}

func JoinErrsStr(delim string, errs... error) string {
	strs := make([]string, len(errs))
	for i, err := range errs {
		strs[i] = err.Error()
	}
	return strings.Join(strs, delim)
}
