package stdlib

import (
	"context"
	"math/rand"

	"github.com/malivvan/cui/runtime"
)

var randModule = map[string]runtime.Object{
	"int": &runtime.BuiltinFunction{
		Name:  "int",
		Value: FuncARI64(rand.Int63),
	},
	"float": &runtime.BuiltinFunction{
		Name:  "float",
		Value: FuncARF(rand.Float64),
	},
	"intn": &runtime.BuiltinFunction{
		Name:  "intn",
		Value: FuncAI64RI64(rand.Int63n),
	},
	"exp_float": &runtime.BuiltinFunction{
		Name:  "exp_float",
		Value: FuncARF(rand.ExpFloat64),
	},
	"norm_float": &runtime.BuiltinFunction{
		Name:  "norm_float",
		Value: FuncARF(rand.NormFloat64),
	},
	"perm": &runtime.BuiltinFunction{
		Name:  "perm",
		Value: FuncAIRIs(rand.Perm),
	},
	"seed": &runtime.BuiltinFunction{
		Name:  "seed",
		Value: FuncAI64R(rand.Seed),
	},
	"read": &runtime.BuiltinFunction{
		Name: "read",
		Value: func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
			if len(args) != 1 {
				return nil, runtime.ErrWrongNumArguments
			}
			y1, ok := args[0].(*runtime.Bytes)
			if !ok {
				return nil, runtime.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "bytes",
					Found:    args[0].TypeName(),
				}
			}
			res, err := rand.Read(y1.Value)
			if err != nil {
				ret = wrapError(err)
				return
			}
			return &runtime.Int{Value: int64(res)}, nil
		},
	},
	"rand": &runtime.BuiltinFunction{
		Name: "rand",
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
			src := rand.NewSource(i1)
			return randRand(rand.New(src)), nil
		},
	},
}

func randRand(r *rand.Rand) *runtime.ImmutableMap {
	return &runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			"int": &runtime.BuiltinFunction{
				Name:  "int",
				Value: FuncARI64(r.Int63),
			},
			"float": &runtime.BuiltinFunction{
				Name:  "float",
				Value: FuncARF(r.Float64),
			},
			"intn": &runtime.BuiltinFunction{
				Name:  "intn",
				Value: FuncAI64RI64(r.Int63n),
			},
			"exp_float": &runtime.BuiltinFunction{
				Name:  "exp_float",
				Value: FuncARF(r.ExpFloat64),
			},
			"norm_float": &runtime.BuiltinFunction{
				Name:  "norm_float",
				Value: FuncARF(r.NormFloat64),
			},
			"perm": &runtime.BuiltinFunction{
				Name:  "perm",
				Value: FuncAIRIs(r.Perm),
			},
			"seed": &runtime.BuiltinFunction{
				Name:  "seed",
				Value: FuncAI64R(r.Seed),
			},
			"read": &runtime.BuiltinFunction{
				Name: "read",
				Value: func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
					if len(args) != 1 {
						return nil, runtime.ErrWrongNumArguments
					}
					y1, ok := args[0].(*runtime.Bytes)
					if !ok {
						return nil, runtime.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "bytes",
							Found:    args[0].TypeName(),
						}
					}
					res, err := r.Read(y1.Value)
					if err != nil {
						ret = wrapError(err)
						return
					}
					return &runtime.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
