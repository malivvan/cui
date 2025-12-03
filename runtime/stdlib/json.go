package stdlib

import (
	"bytes"
	"context"
	gojson "encoding/json"

	"github.com/malivvan/cui/runtime"
	"github.com/malivvan/cui/runtime/stdlib/json"
)

var jsonModule = map[string]runtime.Object{
	"decode": &runtime.BuiltinFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &runtime.BuiltinFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &runtime.BuiltinFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &runtime.BuiltinFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *runtime.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &runtime.Error{
				Value: &runtime.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *runtime.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &runtime.Error{
				Value: &runtime.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	default:
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonEncode(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &runtime.Error{Value: &runtime.String{Value: err.Error()}}, nil
	}

	return &runtime.Bytes{Value: b}, nil
}

func jsonIndent(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 3 {
		return nil, runtime.ErrWrongNumArguments
	}

	prefix, ok := runtime.ToString(args[1])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := runtime.ToString(args[2])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *runtime.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &runtime.Error{
				Value: &runtime.String{Value: err.Error()},
			}, nil
		}
		return &runtime.Bytes{Value: dst.Bytes()}, nil
	case *runtime.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &runtime.Error{
				Value: &runtime.String{Value: err.Error()},
			}, nil
		}
		return &runtime.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *runtime.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &runtime.Bytes{Value: dst.Bytes()}, nil
	case *runtime.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &runtime.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
