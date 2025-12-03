package stdlib

import (
	"context"
	"fmt"

	"github.com/malivvan/cui/runtime"
)

// FuncAR transform a function of 'func()' signature into CallableFunc type.
func FuncAR(fn func()) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		fn()
		return runtime.UndefinedValue, nil
	}
}

// FuncARI transform a function of 'func() int' signature into CallableFunc
// type.
func FuncARI(fn func() int) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		return &runtime.Int{Value: int64(fn())}, nil
	}
}

// FuncARI64 transform a function of 'func() int64' signature into CallableFunc
// type.
func FuncARI64(fn func() int64) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		return &runtime.Int{Value: fn()}, nil
	}
}

// FuncAI64RI64 transform a function of 'func(int64) int64' signature into
// CallableFunc type.
func FuncAI64RI64(fn func(int64) int64) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}

		i1, ok := runtime.ToInt64(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &runtime.Int{Value: fn(i1)}, nil
	}
}

// FuncAI64R transform a function of 'func(int64)' signature into CallableFunc
// type.
func FuncAI64R(fn func(int64)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}

		i1, ok := runtime.ToInt64(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		fn(i1)
		return runtime.UndefinedValue, nil
	}
}

// FuncARB transform a function of 'func() bool' signature into CallableFunc
// type.
func FuncARB(fn func() bool) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		if fn() {
			return runtime.TrueValue, nil
		}
		return runtime.FalseValue, nil
	}
}

// FuncARE transform a function of 'func() error' signature into CallableFunc
// type.
func FuncARE(fn func() error) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		return wrapError(fn()), nil
	}
}

// FuncARS transform a function of 'func() string' signature into CallableFunc
// type.
func FuncARS(fn func() string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		s := fn()
		if len(s) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		return &runtime.String{Value: s}, nil
	}
}

// FuncARSE transform a function of 'func() (string, error)' signature into
// CallableFunc type.
func FuncARSE(fn func() (string, error)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		res, err := fn()
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		return &runtime.String{Value: res}, nil
	}
}

// FuncARYE transform a function of 'func() ([]byte, error)' signature into
// CallableFunc type.
func FuncARYE(fn func() ([]byte, error)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		res, err := fn()
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > runtime.MaxBytesLen {
			return nil, runtime.ErrBytesLimit
		}
		return &runtime.Bytes{Value: res}, nil
	}
}

// FuncARF transform a function of 'func() float64' signature into CallableFunc
// type.
func FuncARF(fn func() float64) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		return &runtime.Float{Value: fn()}, nil
	}
}

// FuncARSs transform a function of 'func() []string' signature into
// CallableFunc type.
func FuncARSs(fn func() []string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		arr := &runtime.Array{}
		for _, elem := range fn() {
			if len(elem) > runtime.MaxStringLen {
				return nil, runtime.ErrStringLimit
			}
			arr.Value = append(arr.Value, &runtime.String{Value: elem})
		}
		return arr, nil
	}
}

// FuncARIsE transform a function of 'func() ([]int, error)' signature into
// CallableFunc type.
func FuncARIsE(fn func() ([]int, error)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 0 {
			return nil, runtime.ErrWrongNumArguments
		}
		res, err := fn()
		if err != nil {
			return wrapError(err), nil
		}
		arr := &runtime.Array{}
		for _, v := range res {
			arr.Value = append(arr.Value, &runtime.Int{Value: int64(v)})
		}
		return arr, nil
	}
}

// FuncAIRIs transform a function of 'func(int) []int' signature into
// CallableFunc type.
func FuncAIRIs(fn func(int) []int) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		i1, ok := runtime.ToInt(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res := fn(i1)
		arr := &runtime.Array{}
		for _, v := range res {
			arr.Value = append(arr.Value, &runtime.Int{Value: int64(v)})
		}
		return arr, nil
	}
}

// FuncAFRF transform a function of 'func(float64) float64' signature into
// CallableFunc type.
func FuncAFRF(fn func(float64) float64) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		f1, ok := runtime.ToFloat64(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &runtime.Float{Value: fn(f1)}, nil
	}
}

// FuncAIR transform a function of 'func(int)' signature into CallableFunc type.
func FuncAIR(fn func(int)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		i1, ok := runtime.ToInt(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		fn(i1)
		return runtime.UndefinedValue, nil
	}
}

// FuncAIRF transform a function of 'func(int) float64' signature into
// CallableFunc type.
func FuncAIRF(fn func(int) float64) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		i1, ok := runtime.ToInt(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &runtime.Float{Value: fn(i1)}, nil
	}
}

// FuncAFRI transform a function of 'func(float64) int' signature into
// CallableFunc type.
func FuncAFRI(fn func(float64) int) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		f1, ok := runtime.ToFloat64(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &runtime.Int{Value: int64(fn(f1))}, nil
	}
}

// FuncAFFRF transform a function of 'func(float64, float64) float64' signature
// into CallableFunc type.
func FuncAFFRF(fn func(float64, float64) float64) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		f1, ok := runtime.ToFloat64(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := runtime.ToFloat64(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return &runtime.Float{Value: fn(f1, f2)}, nil
	}
}

// FuncAIFRF transform a function of 'func(int, float64) float64' signature
// into CallableFunc type.
func FuncAIFRF(fn func(int, float64) float64) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		i1, ok := runtime.ToInt(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := runtime.ToFloat64(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return &runtime.Float{Value: fn(i1, f2)}, nil
	}
}

// FuncAFIRF transform a function of 'func(float64, int) float64' signature
// into CallableFunc type.
func FuncAFIRF(fn func(float64, int) float64) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		f1, ok := runtime.ToFloat64(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := runtime.ToInt(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return &runtime.Float{Value: fn(f1, i2)}, nil
	}
}

// FuncAFIRB transform a function of 'func(float64, int) bool' signature
// into CallableFunc type.
func FuncAFIRB(fn func(float64, int) bool) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		f1, ok := runtime.ToFloat64(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := runtime.ToInt(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		if fn(f1, i2) {
			return runtime.TrueValue, nil
		}
		return runtime.FalseValue, nil
	}
}

// FuncAFRB transform a function of 'func(float64) bool' signature
// into CallableFunc type.
func FuncAFRB(fn func(float64) bool) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		f1, ok := runtime.ToFloat64(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		if fn(f1) {
			return runtime.TrueValue, nil
		}
		return runtime.FalseValue, nil
	}
}

// FuncASRS transform a function of 'func(string) string' signature into
// CallableFunc type. User function will return 'true' if underlying native
// function returns nil.
func FuncASRS(fn func(string) string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s := fn(s1)
		if len(s) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		return &runtime.String{Value: s}, nil
	}
}

// FuncASRSs transform a function of 'func(string) []string' signature into
// CallableFunc type.
func FuncASRSs(fn func(string) []string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res := fn(s1)
		arr := &runtime.Array{}
		for _, elem := range res {
			if len(elem) > runtime.MaxStringLen {
				return nil, runtime.ErrStringLimit
			}
			arr.Value = append(arr.Value, &runtime.String{Value: elem})
		}
		return arr, nil
	}
}

// FuncASRSE transform a function of 'func(string) (string, error)' signature
// into CallableFunc type. User function will return 'true' if underlying
// native function returns nil.
func FuncASRSE(fn func(string) (string, error)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(s1)
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		return &runtime.String{Value: res}, nil
	}
}

// FuncASRE transform a function of 'func(string) error' signature into
// CallableFunc type. User function will return 'true' if underlying native
// function returns nil.
func FuncASRE(fn func(string) error) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return wrapError(fn(s1)), nil
	}
}

// FuncASSRE transform a function of 'func(string, string) error' signature
// into CallableFunc type. User function will return 'true' if underlying
// native function returns nil.
func FuncASSRE(fn func(string, string) error) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := runtime.ToString(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return wrapError(fn(s1, s2)), nil
	}
}

// FuncASSRSs transform a function of 'func(string, string) []string'
// signature into CallableFunc type.
func FuncASSRSs(fn func(string, string) []string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := runtime.ToString(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		arr := &runtime.Array{}
		for _, res := range fn(s1, s2) {
			if len(res) > runtime.MaxStringLen {
				return nil, runtime.ErrStringLimit
			}
			arr.Value = append(arr.Value, &runtime.String{Value: res})
		}
		return arr, nil
	}
}

// FuncASSIRSs transform a function of 'func(string, string, int) []string'
// signature into CallableFunc type.
func FuncASSIRSs(fn func(string, string, int) []string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 3 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := runtime.ToString(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		i3, ok := runtime.ToInt(args[2])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		arr := &runtime.Array{}
		for _, res := range fn(s1, s2, i3) {
			if len(res) > runtime.MaxStringLen {
				return nil, runtime.ErrStringLimit
			}
			arr.Value = append(arr.Value, &runtime.String{Value: res})
		}
		return arr, nil
	}
}

// FuncASSRI transform a function of 'func(string, string) int' signature into
// CallableFunc type.
func FuncASSRI(fn func(string, string) int) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := runtime.ToString(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &runtime.Int{Value: int64(fn(s1, s2))}, nil
	}
}

// FuncASSRS transform a function of 'func(string, string) string' signature
// into CallableFunc type.
func FuncASSRS(fn func(string, string) string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := runtime.ToString(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		s := fn(s1, s2)
		if len(s) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		return &runtime.String{Value: s}, nil
	}
}

// FuncASSRB transform a function of 'func(string, string) bool' signature
// into CallableFunc type.
func FuncASSRB(fn func(string, string) bool) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := runtime.ToString(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		if fn(s1, s2) {
			return runtime.TrueValue, nil
		}
		return runtime.FalseValue, nil
	}
}

// FuncASsSRS transform a function of 'func([]string, string) string' signature
// into CallableFunc type.
func FuncASsSRS(fn func([]string, string) string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		var ss1 []string
		switch arg0 := args[0].(type) {
		case *runtime.Array:
			for idx, a := range arg0.Value {
				as, ok := runtime.ToString(a)
				if !ok {
					return nil, runtime.ErrInvalidArgumentType{
						Name:     fmt.Sprintf("first[%d]", idx),
						Expected: "string(compatible)",
						Found:    a.TypeName(),
					}
				}
				ss1 = append(ss1, as)
			}
		case *runtime.ImmutableArray:
			for idx, a := range arg0.Value {
				as, ok := runtime.ToString(a)
				if !ok {
					return nil, runtime.ErrInvalidArgumentType{
						Name:     fmt.Sprintf("first[%d]", idx),
						Expected: "string(compatible)",
						Found:    a.TypeName(),
					}
				}
				ss1 = append(ss1, as)
			}
		default:
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "array",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := runtime.ToString(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		s := fn(ss1, s2)
		if len(s) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		return &runtime.String{Value: s}, nil
	}
}

// FuncASI64RE transform a function of 'func(string, int64) error' signature
// into CallableFunc type.
func FuncASI64RE(fn func(string, int64) error) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := runtime.ToInt64(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return wrapError(fn(s1, i2)), nil
	}
}

// FuncAIIRE transform a function of 'func(int, int) error' signature
// into CallableFunc type.
func FuncAIIRE(fn func(int, int) error) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		i1, ok := runtime.ToInt(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := runtime.ToInt(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return wrapError(fn(i1, i2)), nil
	}
}

// FuncASIRS transform a function of 'func(string, int) string' signature
// into CallableFunc type.
func FuncASIRS(fn func(string, int) string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 2 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := runtime.ToInt(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		s := fn(s1, i2)
		if len(s) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		return &runtime.String{Value: s}, nil
	}
}

// FuncASIIRE transform a function of 'func(string, int, int) error' signature
// into CallableFunc type.
func FuncASIIRE(fn func(string, int, int) error) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 3 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := runtime.ToInt(args[1])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		i3, ok := runtime.ToInt(args[2])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		return wrapError(fn(s1, i2, i3)), nil
	}
}

// FuncAYRIE transform a function of 'func([]byte) (int, error)' signature
// into CallableFunc type.
func FuncAYRIE(fn func([]byte) (int, error)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		y1, ok := runtime.ToByteSlice(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "bytes(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(y1)
		if err != nil {
			return wrapError(err), nil
		}
		return &runtime.Int{Value: int64(res)}, nil
	}
}

// FuncAYRS transform a function of 'func([]byte) string' signature into
// CallableFunc type.
func FuncAYRS(fn func([]byte) string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		y1, ok := runtime.ToByteSlice(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "bytes(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res := fn(y1)
		return &runtime.String{Value: res}, nil
	}
}

// FuncASRIE transform a function of 'func(string) (int, error)' signature
// into CallableFunc type.
func FuncASRIE(fn func(string) (int, error)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(s1)
		if err != nil {
			return wrapError(err), nil
		}
		return &runtime.Int{Value: int64(res)}, nil
	}
}

// FuncASRYE transform a function of 'func(string) ([]byte, error)' signature
// into CallableFunc type.
func FuncASRYE(fn func(string) ([]byte, error)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		s1, ok := runtime.ToString(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(s1)
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > runtime.MaxBytesLen {
			return nil, runtime.ErrBytesLimit
		}
		return &runtime.Bytes{Value: res}, nil
	}
}

// FuncAIRSsE transform a function of 'func(int) ([]string, error)' signature
// into CallableFunc type.
func FuncAIRSsE(fn func(int) ([]string, error)) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		i1, ok := runtime.ToInt(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(i1)
		if err != nil {
			return wrapError(err), nil
		}
		arr := &runtime.Array{}
		for _, r := range res {
			if len(r) > runtime.MaxStringLen {
				return nil, runtime.ErrStringLimit
			}
			arr.Value = append(arr.Value, &runtime.String{Value: r})
		}
		return arr, nil
	}
}

// FuncAIRS transform a function of 'func(int) string' signature into
// CallableFunc type.
func FuncAIRS(fn func(int) string) runtime.CallableFunc {
	return func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
		if len(args) != 1 {
			return nil, runtime.ErrWrongNumArguments
		}
		i1, ok := runtime.ToInt(args[0])
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s := fn(i1)
		if len(s) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		return &runtime.String{Value: s}, nil
	}
}
