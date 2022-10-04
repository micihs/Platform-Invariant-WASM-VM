package utils

import (
	"fmt"
	"runtime/debug"
)

func CatchPanic(out *error) {
	if err := recover(); err != nil {
		*out = fmt.Errorf("Error: %s\n---GO TRACEBACK---\n%s", UnifyError(err), string(debug.Stack()))
	}
}

func UnifyError(e interface{}) error {
	switch et := e.(type) {
	case error:
		return et
	default:
		return fmt.Errorf("%+v", e)
	}
}