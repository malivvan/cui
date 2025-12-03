package runtime_test

import (
	"strings"
	"testing"
	"time"

	"github.com/malivvan/cui/runtime"
	"github.com/malivvan/cui/runtime/parser"
	"github.com/malivvan/cui/runtime/require"
)

func TestInstructions_String(t *testing.T) {
	assertInstructionString(t,
		[][]byte{
			runtime.MakeInstruction(parser.OpConstant, 1),
			runtime.MakeInstruction(parser.OpConstant, 2),
			runtime.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 CONST   1    
0003 CONST   2    
0006 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			runtime.MakeInstruction(parser.OpBinaryOp, 11),
			runtime.MakeInstruction(parser.OpConstant, 2),
			runtime.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 CONST   2    
0005 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			runtime.MakeInstruction(parser.OpBinaryOp, 11),
			runtime.MakeInstruction(parser.OpGetLocal, 1),
			runtime.MakeInstruction(parser.OpConstant, 2),
			runtime.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 GETL    1    
0004 CONST   2    
0007 CONST   65535`)
}

func TestMakeInstruction(t *testing.T) {
	makeInstruction(t, []byte{parser.OpConstant, 0, 0},
		parser.OpConstant, 0)
	makeInstruction(t, []byte{parser.OpConstant, 0, 1},
		parser.OpConstant, 1)
	makeInstruction(t, []byte{parser.OpConstant, 255, 254},
		parser.OpConstant, 65534)
	makeInstruction(t, []byte{parser.OpPop}, parser.OpPop)
	makeInstruction(t, []byte{parser.OpTrue}, parser.OpTrue)
	makeInstruction(t, []byte{parser.OpFalse}, parser.OpFalse)
}

func TestNumObjects(t *testing.T) {
	testCountObjects(t, &runtime.Array{}, 1)
	testCountObjects(t, &runtime.Array{Value: []runtime.Object{
		&runtime.Int{Value: 1},
		&runtime.Int{Value: 2},
		&runtime.Array{Value: []runtime.Object{
			&runtime.Int{Value: 3},
			&runtime.Int{Value: 4},
			&runtime.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, runtime.TrueValue, 1)
	testCountObjects(t, runtime.FalseValue, 1)
	testCountObjects(t, &runtime.BuiltinFunction{}, 1)
	testCountObjects(t, &runtime.Bytes{Value: []byte("foobar")}, 1)
	testCountObjects(t, &runtime.Char{Value: '가'}, 1)
	testCountObjects(t, &runtime.CompiledFunction{}, 1)
	testCountObjects(t, &runtime.Error{Value: &runtime.Int{Value: 5}}, 2)
	testCountObjects(t, &runtime.Float{Value: 19.84}, 1)
	testCountObjects(t, &runtime.ImmutableArray{Value: []runtime.Object{
		&runtime.Int{Value: 1},
		&runtime.Int{Value: 2},
		&runtime.ImmutableArray{Value: []runtime.Object{
			&runtime.Int{Value: 3},
			&runtime.Int{Value: 4},
			&runtime.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			"k1": &runtime.Int{Value: 1},
			"k2": &runtime.Int{Value: 2},
			"k3": &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 3},
				&runtime.Int{Value: 4},
				&runtime.Int{Value: 5},
			}},
		}}, 7)
	testCountObjects(t, &runtime.Int{Value: 1984}, 1)
	testCountObjects(t, &runtime.Map{Value: map[string]runtime.Object{
		"k1": &runtime.Int{Value: 1},
		"k2": &runtime.Int{Value: 2},
		"k3": &runtime.Array{Value: []runtime.Object{
			&runtime.Int{Value: 3},
			&runtime.Int{Value: 4},
			&runtime.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &runtime.String{Value: "foo bar"}, 1)
	testCountObjects(t, &runtime.Time{Value: time.Now()}, 1)
	testCountObjects(t, runtime.UndefinedValue, 1)
}

func testCountObjects(t *testing.T, o runtime.Object, expected int) {
	require.Equal(t, expected, runtime.CountObjects(o))
}

func assertInstructionString(
	t *testing.T,
	instructions [][]byte,
	expected string,
) {
	concatted := make([]byte, 0)
	for _, e := range instructions {
		concatted = append(concatted, e...)
	}
	require.Equal(t, expected, strings.Join(
		runtime.FormatInstructions(concatted, 0), "\n"))
}

func makeInstruction(
	t *testing.T,
	expected []byte,
	opcode parser.Opcode,
	operands ...int,
) {
	inst := runtime.MakeInstruction(opcode, operands...)
	require.Equal(t, expected, inst)
}
