package stdlib

import (
	"github.com/malivvan/cui/runtime"
)

func wrapError(err error) runtime.Object {
	if err == nil {
		return runtime.TrueValue
	}
	return &runtime.Error{Value: &runtime.String{Value: err.Error()}}
}
