package runtime_test

import (
	"testing"

	"github.com/malivvan/cui/runtime"
	"github.com/malivvan/cui/runtime/require"
	"github.com/malivvan/cui/runtime/token"
)

func TestObject_TypeName(t *testing.T) {
	var o runtime.Object = &runtime.Int{}
	require.Equal(t, "int", o.TypeName())
	o = &runtime.Float{}
	require.Equal(t, "float", o.TypeName())
	o = &runtime.Char{}
	require.Equal(t, "char", o.TypeName())
	o = &runtime.String{}
	require.Equal(t, "string", o.TypeName())
	o = &runtime.Bool{}
	require.Equal(t, "bool", o.TypeName())
	o = &runtime.Array{}
	require.Equal(t, "array", o.TypeName())
	o = &runtime.Map{}
	require.Equal(t, "map", o.TypeName())
	o = &runtime.ArrayIterator{}
	require.Equal(t, "array-iterator", o.TypeName())
	o = &runtime.StringIterator{}
	require.Equal(t, "string-iterator", o.TypeName())
	o = &runtime.MapIterator{}
	require.Equal(t, "map-iterator", o.TypeName())
	o = &runtime.BuiltinFunction{Name: "fn"}
	require.Equal(t, "builtin-function:fn", o.TypeName())
	o = &runtime.CompiledFunction{}
	require.Equal(t, "compiled-function", o.TypeName())
	o = &runtime.Undefined{}
	require.Equal(t, "undefined", o.TypeName())
	o = &runtime.Error{}
	require.Equal(t, "error", o.TypeName())
	o = &runtime.Bytes{}
	require.Equal(t, "bytes", o.TypeName())
}

func TestObject_IsFalsy(t *testing.T) {
	var o runtime.Object = &runtime.Int{Value: 0}
	require.True(t, o.IsFalsy())
	o = &runtime.Int{Value: 1}
	require.False(t, o.IsFalsy())
	o = &runtime.Float{Value: 0}
	require.False(t, o.IsFalsy())
	o = &runtime.Float{Value: 1}
	require.False(t, o.IsFalsy())
	o = &runtime.Char{Value: ' '}
	require.False(t, o.IsFalsy())
	o = &runtime.Char{Value: 'T'}
	require.False(t, o.IsFalsy())
	o = &runtime.String{Value: ""}
	require.True(t, o.IsFalsy())
	o = &runtime.String{Value: " "}
	require.False(t, o.IsFalsy())
	o = &runtime.Array{Value: nil}
	require.True(t, o.IsFalsy())
	o = &runtime.Array{Value: []runtime.Object{nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &runtime.Map{Value: nil}
	require.True(t, o.IsFalsy())
	o = &runtime.Map{Value: map[string]runtime.Object{"a": nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &runtime.StringIterator{}
	require.True(t, o.IsFalsy())
	o = &runtime.ArrayIterator{}
	require.True(t, o.IsFalsy())
	o = &runtime.MapIterator{}
	require.True(t, o.IsFalsy())
	o = &runtime.BuiltinFunction{}
	require.False(t, o.IsFalsy())
	o = &runtime.CompiledFunction{}
	require.False(t, o.IsFalsy())
	o = &runtime.Undefined{}
	require.True(t, o.IsFalsy())
	o = &runtime.Error{}
	require.True(t, o.IsFalsy())
	o = &runtime.Bytes{}
	require.True(t, o.IsFalsy())
	o = &runtime.Bytes{Value: []byte{1, 2}}
	require.False(t, o.IsFalsy())
}

func TestObject_String(t *testing.T) {
	var o runtime.Object = &runtime.Int{Value: 0}
	require.Equal(t, "0", o.String())
	o = &runtime.Int{Value: 1}
	require.Equal(t, "1", o.String())
	o = &runtime.Float{Value: 0}
	require.Equal(t, "0", o.String())
	o = &runtime.Float{Value: 1}
	require.Equal(t, "1", o.String())
	o = &runtime.Char{Value: ' '}
	require.Equal(t, " ", o.String())
	o = &runtime.Char{Value: 'T'}
	require.Equal(t, "T", o.String())
	o = &runtime.String{Value: ""}
	require.Equal(t, `""`, o.String())
	o = &runtime.String{Value: " "}
	require.Equal(t, `" "`, o.String())
	o = &runtime.Array{Value: nil}
	require.Equal(t, "[]", o.String())
	o = &runtime.Map{Value: nil}
	require.Equal(t, "{}", o.String())
	o = &runtime.Error{Value: nil}
	require.Equal(t, "error", o.String())
	o = &runtime.Error{Value: &runtime.String{Value: "error 1"}}
	require.Equal(t, `error: "error 1"`, o.String())
	o = &runtime.StringIterator{}
	require.Equal(t, "<string-iterator>", o.String())
	o = &runtime.ArrayIterator{}
	require.Equal(t, "<array-iterator>", o.String())
	o = &runtime.MapIterator{}
	require.Equal(t, "<map-iterator>", o.String())
	o = &runtime.Undefined{}
	require.Equal(t, "<undefined>", o.String())
	o = &runtime.Bytes{}
	require.Equal(t, "", o.String())
	o = &runtime.Bytes{Value: []byte("foo")}
	require.Equal(t, "foo", o.String())
}

func TestObject_BinaryOp(t *testing.T) {
	var o runtime.Object = &runtime.Char{}
	_, err := o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.Bool{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.Map{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.ArrayIterator{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.StringIterator{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.MapIterator{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.BuiltinFunction{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.CompiledFunction{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.Undefined{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
	o = &runtime.Error{}
	_, err = o.BinaryOp(token.Add, runtime.UndefinedValue)
	require.Error(t, err)
}

func TestArray_BinaryOp(t *testing.T) {
	testBinaryOp(t, &runtime.Array{Value: nil}, token.Add,
		&runtime.Array{Value: nil}, &runtime.Array{Value: nil})
	testBinaryOp(t, &runtime.Array{Value: nil}, token.Add,
		&runtime.Array{Value: []runtime.Object{}}, &runtime.Array{Value: nil})
	testBinaryOp(t, &runtime.Array{Value: []runtime.Object{}}, token.Add,
		&runtime.Array{Value: nil}, &runtime.Array{Value: []runtime.Object{}})
	testBinaryOp(t, &runtime.Array{Value: []runtime.Object{}}, token.Add,
		&runtime.Array{Value: []runtime.Object{}},
		&runtime.Array{Value: []runtime.Object{}})
	testBinaryOp(t, &runtime.Array{Value: nil}, token.Add,
		&runtime.Array{Value: []runtime.Object{
			&runtime.Int{Value: 1},
		}}, &runtime.Array{Value: []runtime.Object{
			&runtime.Int{Value: 1},
		}})
	testBinaryOp(t, &runtime.Array{Value: nil}, token.Add,
		&runtime.Array{Value: []runtime.Object{
			&runtime.Int{Value: 1},
			&runtime.Int{Value: 2},
			&runtime.Int{Value: 3},
		}}, &runtime.Array{Value: []runtime.Object{
			&runtime.Int{Value: 1},
			&runtime.Int{Value: 2},
			&runtime.Int{Value: 3},
		}})
	testBinaryOp(t, &runtime.Array{Value: []runtime.Object{
		&runtime.Int{Value: 1},
		&runtime.Int{Value: 2},
		&runtime.Int{Value: 3},
	}}, token.Add, &runtime.Array{Value: nil},
		&runtime.Array{Value: []runtime.Object{
			&runtime.Int{Value: 1},
			&runtime.Int{Value: 2},
			&runtime.Int{Value: 3},
		}})
	testBinaryOp(t, &runtime.Array{Value: []runtime.Object{
		&runtime.Int{Value: 1},
		&runtime.Int{Value: 2},
		&runtime.Int{Value: 3},
	}}, token.Add, &runtime.Array{Value: []runtime.Object{
		&runtime.Int{Value: 4},
		&runtime.Int{Value: 5},
		&runtime.Int{Value: 6},
	}}, &runtime.Array{Value: []runtime.Object{
		&runtime.Int{Value: 1},
		&runtime.Int{Value: 2},
		&runtime.Int{Value: 3},
		&runtime.Int{Value: 4},
		&runtime.Int{Value: 5},
		&runtime.Int{Value: 6},
	}})
}

func TestError_Equals(t *testing.T) {
	err1 := &runtime.Error{Value: &runtime.String{Value: "some error"}}
	err2 := err1
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = &runtime.Error{Value: &runtime.String{Value: "some error"}}
	require.False(t, err1.Equals(err2))
	require.False(t, err2.Equals(err1))
}

func TestFloat_BinaryOp(t *testing.T) {
	// float + float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Add,
				&runtime.Float{Value: r}, &runtime.Float{Value: l + r})
		}
	}

	// float - float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Sub,
				&runtime.Float{Value: r}, &runtime.Float{Value: l - r})
		}
	}

	// float * float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Mul,
				&runtime.Float{Value: r}, &runtime.Float{Value: l * r})
		}
	}

	// float / float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			if r != 0 {
				testBinaryOp(t, &runtime.Float{Value: l}, token.Quo,
					&runtime.Float{Value: r}, &runtime.Float{Value: l / r})
			}
		}
	}

	// float < float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Less,
				&runtime.Float{Value: r}, boolValue(l < r))
		}
	}

	// float > float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Greater,
				&runtime.Float{Value: r}, boolValue(l > r))
		}
	}

	// float <= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &runtime.Float{Value: l}, token.LessEq,
				&runtime.Float{Value: r}, boolValue(l <= r))
		}
	}

	// float >= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &runtime.Float{Value: l}, token.GreaterEq,
				&runtime.Float{Value: r}, boolValue(l >= r))
		}
	}

	// float + int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Add,
				&runtime.Int{Value: r}, &runtime.Float{Value: l + float64(r)})
		}
	}

	// float - int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Sub,
				&runtime.Int{Value: r}, &runtime.Float{Value: l - float64(r)})
		}
	}

	// float * int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Mul,
				&runtime.Int{Value: r}, &runtime.Float{Value: l * float64(r)})
		}
	}

	// float / int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &runtime.Float{Value: l}, token.Quo,
					&runtime.Int{Value: r},
					&runtime.Float{Value: l / float64(r)})
			}
		}
	}

	// float < int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Less,
				&runtime.Int{Value: r}, boolValue(l < float64(r)))
		}
	}

	// float > int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Float{Value: l}, token.Greater,
				&runtime.Int{Value: r}, boolValue(l > float64(r)))
		}
	}

	// float <= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Float{Value: l}, token.LessEq,
				&runtime.Int{Value: r}, boolValue(l <= float64(r)))
		}
	}

	// float >= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Float{Value: l}, token.GreaterEq,
				&runtime.Int{Value: r}, boolValue(l >= float64(r)))
		}
	}
}

func TestInt_BinaryOp(t *testing.T) {
	// int + int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Add,
				&runtime.Int{Value: r}, &runtime.Int{Value: l + r})
		}
	}

	// int - int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Sub,
				&runtime.Int{Value: r}, &runtime.Int{Value: l - r})
		}
	}

	// int * int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Mul,
				&runtime.Int{Value: r}, &runtime.Int{Value: l * r})
		}
	}

	// int / int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &runtime.Int{Value: l}, token.Quo,
					&runtime.Int{Value: r}, &runtime.Int{Value: l / r})
			}
		}
	}

	// int % int
	for l := int64(-4); l <= 4; l++ {
		for r := -int64(-4); r <= 4; r++ {
			if r == 0 {
				testBinaryOp(t, &runtime.Int{Value: l}, token.Rem,
					&runtime.Int{Value: r}, &runtime.Int{Value: l % r})
			}
		}
	}

	// int & int
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.And, &runtime.Int{Value: 0},
		&runtime.Int{Value: int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.And, &runtime.Int{Value: 0},
		&runtime.Int{Value: int64(1) & int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.And, &runtime.Int{Value: 1},
		&runtime.Int{Value: int64(0) & int64(1)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.And, &runtime.Int{Value: 1},
		&runtime.Int{Value: int64(1)})
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.And, &runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(0) & int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.And, &runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(1) & int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: int64(0xffffffff)}, token.And,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: 1984}, token.And,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(1984) & int64(0xffffffff)})
	testBinaryOp(t, &runtime.Int{Value: -1984}, token.And,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(-1984) & int64(0xffffffff)})

	// int | int
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.Or, &runtime.Int{Value: 0},
		&runtime.Int{Value: int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.Or, &runtime.Int{Value: 0},
		&runtime.Int{Value: int64(1) | int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.Or, &runtime.Int{Value: 1},
		&runtime.Int{Value: int64(0) | int64(1)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.Or, &runtime.Int{Value: 1},
		&runtime.Int{Value: int64(1)})
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.Or, &runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(0) | int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.Or, &runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(1) | int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: int64(0xffffffff)}, token.Or,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: 1984}, token.Or,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(1984) | int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: -1984}, token.Or,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(-1984) | int64(0xffffffff)})

	// int ^ int
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.Xor, &runtime.Int{Value: 0},
		&runtime.Int{Value: int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.Xor, &runtime.Int{Value: 0},
		&runtime.Int{Value: int64(1) ^ int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.Xor, &runtime.Int{Value: 1},
		&runtime.Int{Value: int64(0) ^ int64(1)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.Xor, &runtime.Int{Value: 1},
		&runtime.Int{Value: int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.Xor, &runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(0) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.Xor, &runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(1) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: int64(0xffffffff)}, token.Xor,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 1984}, token.Xor,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(1984) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: -1984}, token.Xor,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(-1984) ^ int64(0xffffffff)})

	// int &^ int
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.AndNot, &runtime.Int{Value: 0},
		&runtime.Int{Value: int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.AndNot, &runtime.Int{Value: 0},
		&runtime.Int{Value: int64(1) &^ int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.AndNot,
		&runtime.Int{Value: 1}, &runtime.Int{Value: int64(0) &^ int64(1)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.AndNot, &runtime.Int{Value: 1},
		&runtime.Int{Value: int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 0}, token.AndNot,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(0) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: 1}, token.AndNot,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(1) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: int64(0xffffffff)}, token.AndNot,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(0)})
	testBinaryOp(t,
		&runtime.Int{Value: 1984}, token.AndNot,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(1984) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&runtime.Int{Value: -1984}, token.AndNot,
		&runtime.Int{Value: int64(0xffffffff)},
		&runtime.Int{Value: int64(-1984) &^ int64(0xffffffff)})

	// int << int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&runtime.Int{Value: 0}, token.Shl, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(0) << uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: 1}, token.Shl, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(1) << uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: 2}, token.Shl, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(2) << uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: -1}, token.Shl, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(-1) << uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: -2}, token.Shl, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(-2) << uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: int64(0xffffffff)}, token.Shl,
			&runtime.Int{Value: s},
			&runtime.Int{Value: int64(0xffffffff) << uint(s)})
	}

	// int >> int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&runtime.Int{Value: 0}, token.Shr, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(0) >> uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: 1}, token.Shr, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(1) >> uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: 2}, token.Shr, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(2) >> uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: -1}, token.Shr, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(-1) >> uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: -2}, token.Shr, &runtime.Int{Value: s},
			&runtime.Int{Value: int64(-2) >> uint(s)})
		testBinaryOp(t,
			&runtime.Int{Value: int64(0xffffffff)}, token.Shr,
			&runtime.Int{Value: s},
			&runtime.Int{Value: int64(0xffffffff) >> uint(s)})
	}

	// int < int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Less,
				&runtime.Int{Value: r}, boolValue(l < r))
		}
	}

	// int > int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Greater,
				&runtime.Int{Value: r}, boolValue(l > r))
		}
	}

	// int <= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Int{Value: l}, token.LessEq,
				&runtime.Int{Value: r}, boolValue(l <= r))
		}
	}

	// int >= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &runtime.Int{Value: l}, token.GreaterEq,
				&runtime.Int{Value: r}, boolValue(l >= r))
		}
	}

	// int + float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Add,
				&runtime.Float{Value: r},
				&runtime.Float{Value: float64(l) + r})
		}
	}

	// int - float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Sub,
				&runtime.Float{Value: r},
				&runtime.Float{Value: float64(l) - r})
		}
	}

	// int * float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Mul,
				&runtime.Float{Value: r},
				&runtime.Float{Value: float64(l) * r})
		}
	}

	// int / float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			if r != 0 {
				testBinaryOp(t, &runtime.Int{Value: l}, token.Quo,
					&runtime.Float{Value: r},
					&runtime.Float{Value: float64(l) / r})
			}
		}
	}

	// int < float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Less,
				&runtime.Float{Value: r}, boolValue(float64(l) < r))
		}
	}

	// int > float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &runtime.Int{Value: l}, token.Greater,
				&runtime.Float{Value: r}, boolValue(float64(l) > r))
		}
	}

	// int <= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &runtime.Int{Value: l}, token.LessEq,
				&runtime.Float{Value: r}, boolValue(float64(l) <= r))
		}
	}

	// int >= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &runtime.Int{Value: l}, token.GreaterEq,
				&runtime.Float{Value: r}, boolValue(float64(l) >= r))
		}
	}
}

func TestMap_Index(t *testing.T) {
	m := &runtime.Map{Value: make(map[string]runtime.Object)}
	k := &runtime.Int{Value: 1}
	v := &runtime.String{Value: "abcdef"}
	err := m.IndexSet(k, v)

	require.NoError(t, err)

	res, err := m.IndexGet(k)
	require.NoError(t, err)
	require.Equal(t, v, res)
}

func TestString_BinaryOp(t *testing.T) {
	lstr := "abcde"
	rstr := "01234"
	for l := 0; l < len(lstr); l++ {
		for r := 0; r < len(rstr); r++ {
			ls := lstr[l:]
			rs := rstr[r:]
			testBinaryOp(t, &runtime.String{Value: ls}, token.Add,
				&runtime.String{Value: rs},
				&runtime.String{Value: ls + rs})

			rc := []rune(rstr)[r]
			testBinaryOp(t, &runtime.String{Value: ls}, token.Add,
				&runtime.Char{Value: rc},
				&runtime.String{Value: ls + string(rc)})
		}
	}
}

func testBinaryOp(
	t *testing.T,
	lhs runtime.Object,
	op token.Token,
	rhs runtime.Object,
	expected runtime.Object,
) {
	t.Helper()
	actual, err := lhs.BinaryOp(op, rhs)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func boolValue(b bool) runtime.Object {
	if b {
		return runtime.TrueValue
	}
	return runtime.FalseValue
}
