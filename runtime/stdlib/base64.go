package stdlib

import (
	"encoding/base64"

	"github.com/malivvan/cui/runtime"
)

var base64Module = map[string]runtime.Object{
	"encode": &runtime.BuiltinFunction{
		Value: FuncAYRS(base64.StdEncoding.EncodeToString),
	},
	"decode": &runtime.BuiltinFunction{
		Value: FuncASRYE(base64.StdEncoding.DecodeString),
	},
	"raw_encode": &runtime.BuiltinFunction{
		Value: FuncAYRS(base64.RawStdEncoding.EncodeToString),
	},
	"raw_decode": &runtime.BuiltinFunction{
		Value: FuncASRYE(base64.RawStdEncoding.DecodeString),
	},
	"url_encode": &runtime.BuiltinFunction{
		Value: FuncAYRS(base64.URLEncoding.EncodeToString),
	},
	"url_decode": &runtime.BuiltinFunction{
		Value: FuncASRYE(base64.URLEncoding.DecodeString),
	},
	"raw_url_encode": &runtime.BuiltinFunction{
		Value: FuncAYRS(base64.RawURLEncoding.EncodeToString),
	},
	"raw_url_decode": &runtime.BuiltinFunction{
		Value: FuncASRYE(base64.RawURLEncoding.DecodeString),
	},
}
