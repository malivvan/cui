package stdlib

import (
	"context"
	"os"
	"syscall"

	"github.com/malivvan/cui/runtime"
)

func makeOSProcessState(state *os.ProcessState) *runtime.ImmutableMap {
	return &runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			"exited": &runtime.BuiltinFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &runtime.BuiltinFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &runtime.BuiltinFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &runtime.BuiltinFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *runtime.ImmutableMap {
	return &runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			"kill": &runtime.BuiltinFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &runtime.BuiltinFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &runtime.BuiltinFunction{
				Name: "signal",
				Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
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
					return wrapError(proc.Signal(syscall.Signal(i1))), nil
				},
			},
			"wait": &runtime.BuiltinFunction{
				Name: "wait",
				Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
					if len(args) != 0 {
						return nil, runtime.ErrWrongNumArguments
					}
					state, err := proc.Wait()
					if err != nil {
						return wrapError(err), nil
					}
					return makeOSProcessState(state), nil
				},
			},
		},
	}
}
