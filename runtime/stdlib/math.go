package stdlib

import (
	"math"

	"github.com/malivvan/cui/runtime"
)

var mathModule = map[string]runtime.Object{
	"e":       &runtime.Float{Value: math.E},
	"pi":      &runtime.Float{Value: math.Pi},
	"phi":     &runtime.Float{Value: math.Phi},
	"sqrt2":   &runtime.Float{Value: math.Sqrt2},
	"sqrtE":   &runtime.Float{Value: math.SqrtE},
	"sqrtPi":  &runtime.Float{Value: math.SqrtPi},
	"sqrtPhi": &runtime.Float{Value: math.SqrtPhi},
	"ln2":     &runtime.Float{Value: math.Ln2},
	"log2E":   &runtime.Float{Value: math.Log2E},
	"ln10":    &runtime.Float{Value: math.Ln10},
	"log10E":  &runtime.Float{Value: math.Log10E},
	"abs": &runtime.BuiltinFunction{
		Name:  "abs",
		Value: FuncAFRF(math.Abs),
	},
	"acos": &runtime.BuiltinFunction{
		Name:  "acos",
		Value: FuncAFRF(math.Acos),
	},
	"acosh": &runtime.BuiltinFunction{
		Name:  "acosh",
		Value: FuncAFRF(math.Acosh),
	},
	"asin": &runtime.BuiltinFunction{
		Name:  "asin",
		Value: FuncAFRF(math.Asin),
	},
	"asinh": &runtime.BuiltinFunction{
		Name:  "asinh",
		Value: FuncAFRF(math.Asinh),
	},
	"atan": &runtime.BuiltinFunction{
		Name:  "atan",
		Value: FuncAFRF(math.Atan),
	},
	"atan2": &runtime.BuiltinFunction{
		Name:  "atan2",
		Value: FuncAFFRF(math.Atan2),
	},
	"atanh": &runtime.BuiltinFunction{
		Name:  "atanh",
		Value: FuncAFRF(math.Atanh),
	},
	"cbrt": &runtime.BuiltinFunction{
		Name:  "cbrt",
		Value: FuncAFRF(math.Cbrt),
	},
	"ceil": &runtime.BuiltinFunction{
		Name:  "ceil",
		Value: FuncAFRF(math.Ceil),
	},
	"copysign": &runtime.BuiltinFunction{
		Name:  "copysign",
		Value: FuncAFFRF(math.Copysign),
	},
	"cos": &runtime.BuiltinFunction{
		Name:  "cos",
		Value: FuncAFRF(math.Cos),
	},
	"cosh": &runtime.BuiltinFunction{
		Name:  "cosh",
		Value: FuncAFRF(math.Cosh),
	},
	"dim": &runtime.BuiltinFunction{
		Name:  "dim",
		Value: FuncAFFRF(math.Dim),
	},
	"erf": &runtime.BuiltinFunction{
		Name:  "erf",
		Value: FuncAFRF(math.Erf),
	},
	"erfc": &runtime.BuiltinFunction{
		Name:  "erfc",
		Value: FuncAFRF(math.Erfc),
	},
	"exp": &runtime.BuiltinFunction{
		Name:  "exp",
		Value: FuncAFRF(math.Exp),
	},
	"exp2": &runtime.BuiltinFunction{
		Name:  "exp2",
		Value: FuncAFRF(math.Exp2),
	},
	"expm1": &runtime.BuiltinFunction{
		Name:  "expm1",
		Value: FuncAFRF(math.Expm1),
	},
	"floor": &runtime.BuiltinFunction{
		Name:  "floor",
		Value: FuncAFRF(math.Floor),
	},
	"gamma": &runtime.BuiltinFunction{
		Name:  "gamma",
		Value: FuncAFRF(math.Gamma),
	},
	"hypot": &runtime.BuiltinFunction{
		Name:  "hypot",
		Value: FuncAFFRF(math.Hypot),
	},
	"ilogb": &runtime.BuiltinFunction{
		Name:  "ilogb",
		Value: FuncAFRI(math.Ilogb),
	},
	"inf": &runtime.BuiltinFunction{
		Name:  "inf",
		Value: FuncAIRF(math.Inf),
	},
	"is_inf": &runtime.BuiltinFunction{
		Name:  "is_inf",
		Value: FuncAFIRB(math.IsInf),
	},
	"is_nan": &runtime.BuiltinFunction{
		Name:  "is_nan",
		Value: FuncAFRB(math.IsNaN),
	},
	"j0": &runtime.BuiltinFunction{
		Name:  "j0",
		Value: FuncAFRF(math.J0),
	},
	"j1": &runtime.BuiltinFunction{
		Name:  "j1",
		Value: FuncAFRF(math.J1),
	},
	"jn": &runtime.BuiltinFunction{
		Name:  "jn",
		Value: FuncAIFRF(math.Jn),
	},
	"ldexp": &runtime.BuiltinFunction{
		Name:  "ldexp",
		Value: FuncAFIRF(math.Ldexp),
	},
	"log": &runtime.BuiltinFunction{
		Name:  "log",
		Value: FuncAFRF(math.Log),
	},
	"log10": &runtime.BuiltinFunction{
		Name:  "log10",
		Value: FuncAFRF(math.Log10),
	},
	"log1p": &runtime.BuiltinFunction{
		Name:  "log1p",
		Value: FuncAFRF(math.Log1p),
	},
	"log2": &runtime.BuiltinFunction{
		Name:  "log2",
		Value: FuncAFRF(math.Log2),
	},
	"logb": &runtime.BuiltinFunction{
		Name:  "logb",
		Value: FuncAFRF(math.Logb),
	},
	"max": &runtime.BuiltinFunction{
		Name:  "max",
		Value: FuncAFFRF(math.Max),
	},
	"min": &runtime.BuiltinFunction{
		Name:  "min",
		Value: FuncAFFRF(math.Min),
	},
	"mod": &runtime.BuiltinFunction{
		Name:  "mod",
		Value: FuncAFFRF(math.Mod),
	},
	"nan": &runtime.BuiltinFunction{
		Name:  "nan",
		Value: FuncARF(math.NaN),
	},
	"nextafter": &runtime.BuiltinFunction{
		Name:  "nextafter",
		Value: FuncAFFRF(math.Nextafter),
	},
	"pow": &runtime.BuiltinFunction{
		Name:  "pow",
		Value: FuncAFFRF(math.Pow),
	},
	"pow10": &runtime.BuiltinFunction{
		Name:  "pow10",
		Value: FuncAIRF(math.Pow10),
	},
	"remainder": &runtime.BuiltinFunction{
		Name:  "remainder",
		Value: FuncAFFRF(math.Remainder),
	},
	"signbit": &runtime.BuiltinFunction{
		Name:  "signbit",
		Value: FuncAFRB(math.Signbit),
	},
	"sin": &runtime.BuiltinFunction{
		Name:  "sin",
		Value: FuncAFRF(math.Sin),
	},
	"sinh": &runtime.BuiltinFunction{
		Name:  "sinh",
		Value: FuncAFRF(math.Sinh),
	},
	"sqrt": &runtime.BuiltinFunction{
		Name:  "sqrt",
		Value: FuncAFRF(math.Sqrt),
	},
	"tan": &runtime.BuiltinFunction{
		Name:  "tan",
		Value: FuncAFRF(math.Tan),
	},
	"tanh": &runtime.BuiltinFunction{
		Name:  "tanh",
		Value: FuncAFRF(math.Tanh),
	},
	"trunc": &runtime.BuiltinFunction{
		Name:  "trunc",
		Value: FuncAFRF(math.Trunc),
	},
	"y0": &runtime.BuiltinFunction{
		Name:  "y0",
		Value: FuncAFRF(math.Y0),
	},
	"y1": &runtime.BuiltinFunction{
		Name:  "y1",
		Value: FuncAFRF(math.Y1),
	},
	"yn": &runtime.BuiltinFunction{
		Name:  "yn",
		Value: FuncAIFRF(math.Yn),
	},
}
