package stdlib

import (
	"context"
	"fmt"

	"github.com/malivvan/cui/runtime"
)

var fmtModule = map[string]runtime.Object{
	"print":   &runtime.BuiltinFunction{Value: fmtPrint},
	"printf":  &runtime.BuiltinFunction{Value: fmtPrintf},
	"println": &runtime.BuiltinFunction{Value: fmtPrintln},
	"sprintf": &runtime.BuiltinFunction{Name: "sprintf", Value: fmtSprintf},
}

func fmtPrint(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	vm := ctx.Value(runtime.ContextKey("vm")).(*runtime.VM)
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	fmt.Fprint(vm.Out, printArgs...)
	return nil, nil
}

func fmtPrintf(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	vm := ctx.Value(runtime.ContextKey("vm")).(*runtime.VM)
	numArgs := len(args)
	if numArgs == 0 {
		return nil, runtime.ErrWrongNumArguments
	}

	format, ok := args[0].(*runtime.String)
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		fmt.Fprint(vm.Out, format)
		return nil, nil
	}

	s, err := runtime.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	fmt.Fprint(vm.Out, s)
	return nil, nil
}

func fmtPrintln(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	vm := ctx.Value(runtime.ContextKey("vm")).(*runtime.VM)
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	printArgs = append(printArgs, "\n")
	fmt.Fprint(vm.Out, printArgs...)
	return nil, nil
}

func fmtSprintf(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, runtime.ErrWrongNumArguments
	}

	format, ok := args[0].(*runtime.String)
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := runtime.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &runtime.String{Value: s}, nil
}

func getPrintArgs(args ...runtime.Object) ([]interface{}, error) {
	var printArgs []interface{}
	l := 0
	for _, arg := range args {
		s, _ := runtime.ToString(arg)
		slen := len(s)
		// make sure length does not exceed the limit
		if l+slen > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		l += slen
		printArgs = append(printArgs, s)
	}
	return printArgs, nil
}
