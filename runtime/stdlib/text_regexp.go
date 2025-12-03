package stdlib

import (
	"context"
	"regexp"

	"github.com/malivvan/cui/runtime"
)

func makeTextRegexp(re *regexp.Regexp) *runtime.ImmutableMap {
	return &runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			// match(text) => bool
			"match": &runtime.BuiltinFunction{
				Value: func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
					if len(args) != 1 {
						err = runtime.ErrWrongNumArguments
						return
					}

					s1, ok := runtime.ToString(args[0])
					if !ok {
						err = runtime.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if re.MatchString(s1) {
						ret = runtime.TrueValue
					} else {
						ret = runtime.FalseValue
					}

					return
				},
			},

			// find(text) 			=> array(array({text:,begin:,end:}))/undefined
			// find(text, maxCount) => array(array({text:,begin:,end:}))/undefined
			"find": &runtime.BuiltinFunction{
				Value: func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = runtime.ErrWrongNumArguments
						return
					}

					s1, ok := runtime.ToString(args[0])
					if !ok {
						err = runtime.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if numArgs == 1 {
						m := re.FindStringSubmatchIndex(s1)
						if m == nil {
							ret = runtime.UndefinedValue
							return
						}

						arr := &runtime.Array{}
						for i := 0; i < len(m); i += 2 {
							arr.Value = append(arr.Value,
								&runtime.ImmutableMap{
									Value: map[string]runtime.Object{
										"text": &runtime.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &runtime.Int{
											Value: int64(m[i]),
										},
										"end": &runtime.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						ret = &runtime.Array{Value: []runtime.Object{arr}}

						return
					}

					i2, ok := runtime.ToInt(args[1])
					if !ok {
						err = runtime.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}
					m := re.FindAllStringSubmatchIndex(s1, i2)
					if m == nil {
						ret = runtime.UndefinedValue
						return
					}

					arr := &runtime.Array{}
					for _, m := range m {
						subMatch := &runtime.Array{}
						for i := 0; i < len(m); i += 2 {
							subMatch.Value = append(subMatch.Value,
								&runtime.ImmutableMap{
									Value: map[string]runtime.Object{
										"text": &runtime.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &runtime.Int{
											Value: int64(m[i]),
										},
										"end": &runtime.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						arr.Value = append(arr.Value, subMatch)
					}

					ret = arr

					return
				},
			},

			// replace(src, repl) => string
			"replace": &runtime.BuiltinFunction{
				Value: func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
					if len(args) != 2 {
						err = runtime.ErrWrongNumArguments
						return
					}

					s1, ok := runtime.ToString(args[0])
					if !ok {
						err = runtime.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					s2, ok := runtime.ToString(args[1])
					if !ok {
						err = runtime.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "string(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}

					s, ok := doTextRegexpReplace(re, s1, s2)
					if !ok {
						return nil, runtime.ErrStringLimit
					}

					ret = &runtime.String{Value: s}

					return
				},
			},

			// split(text) 			 => array(string)
			// split(text, maxCount) => array(string)
			"split": &runtime.BuiltinFunction{
				Value: func(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = runtime.ErrWrongNumArguments
						return
					}

					s1, ok := runtime.ToString(args[0])
					if !ok {
						err = runtime.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					var i2 = -1
					if numArgs > 1 {
						i2, ok = runtime.ToInt(args[1])
						if !ok {
							err = runtime.ErrInvalidArgumentType{
								Name:     "second",
								Expected: "int(compatible)",
								Found:    args[1].TypeName(),
							}
							return
						}
					}

					arr := &runtime.Array{}
					for _, s := range re.Split(s1, i2) {
						arr.Value = append(arr.Value,
							&runtime.String{Value: s})
					}

					ret = arr

					return
				},
			},
		},
	}
}

// Size-limit checking implementation of regexp.ReplaceAllString.
func doTextRegexpReplace(re *regexp.Regexp, src, repl string) (string, bool) {
	idx := 0
	out := ""
	for _, m := range re.FindAllStringSubmatchIndex(src, -1) {
		var exp []byte
		exp = re.ExpandString(exp, repl, src, m)
		if len(out)+m[0]-idx+len(exp) > runtime.MaxStringLen {
			return "", false
		}
		out += src[idx:m[0]] + string(exp)
		idx = m[1]
	}
	if idx < len(src) {
		if len(out)+len(src)-idx > runtime.MaxStringLen {
			return "", false
		}
		out += src[idx:]
	}
	return out, true
}
