package stdlib

import (
	"context"
	"os/exec"

	"github.com/malivvan/cui/runtime"
)

func makeOSExecCommand(cmd *exec.Cmd) *runtime.ImmutableMap {
	return &runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			// combined_output() => bytes/error
			"combined_output": &runtime.BuiltinFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &runtime.BuiltinFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &runtime.BuiltinFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &runtime.BuiltinFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &runtime.BuiltinFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &runtime.BuiltinFunction{
				Name: "set_path",
				Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
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
					cmd.Path = s1
					return runtime.UndefinedValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &runtime.BuiltinFunction{
				Name: "set_dir",
				Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
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
					cmd.Dir = s1
					return runtime.UndefinedValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &runtime.BuiltinFunction{
				Name: "set_env",
				Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
					if len(args) != 1 {
						return nil, runtime.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *runtime.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *runtime.ImmutableArray:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					default:
						return nil, runtime.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "array",
							Found:    arg0.TypeName(),
						}
					}
					cmd.Env = env
					return runtime.UndefinedValue, nil
				},
			},
			// process() => imap(process)
			"process": &runtime.BuiltinFunction{
				Name: "process",
				Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
					if len(args) != 0 {
						return nil, runtime.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
