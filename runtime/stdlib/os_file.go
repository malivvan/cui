package stdlib

import (
	"context"
	"os"

	"github.com/malivvan/cui/runtime"
)

func makeOSFile(file *os.File) *runtime.ImmutableMap {
	return &runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			// chdir() => true/error
			"chdir": &runtime.BuiltinFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &runtime.BuiltinFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &runtime.BuiltinFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &runtime.BuiltinFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &runtime.BuiltinFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &runtime.BuiltinFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &runtime.BuiltinFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &runtime.BuiltinFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &runtime.BuiltinFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &runtime.BuiltinFunction{
				Name: "chmod",
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
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &runtime.BuiltinFunction{
				Name: "seek",
				Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
					if len(args) != 2 {
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
					i2, ok := runtime.ToInt(args[1])
					if !ok {
						return nil, runtime.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &runtime.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &runtime.BuiltinFunction{
				Name: "stat",
				Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
					if len(args) != 0 {
						return nil, runtime.ErrWrongNumArguments
					}
					return osStat(ctx, &runtime.String{Value: file.Name()})
				},
			},
		},
	}
}
