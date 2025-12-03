package stdlib

import (
	"encoding/hex"
	"github.com/malivvan/cui/runtime"
)

var hexModule = map[string]runtime.Object{
	"encode": &runtime.BuiltinFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &runtime.BuiltinFunction{Value: FuncASRYE(hex.DecodeString)},
}
