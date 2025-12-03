package runtime_test

import (
	"testing"
	"time"

	"github.com/malivvan/cui/runtime"
	"github.com/malivvan/cui/runtime/parser"
	"github.com/malivvan/cui/runtime/require"
)

type srcfile struct {
	name string
	size int
}

func TestBytecode(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))

	testBytecodeSerialization(t, bytecode(
		concatInsts(), objectsArray(
			&runtime.Char{Value: 'y'},
			&runtime.Float{Value: 93.11},
			compiledFunction(1, 0,
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpSetLocal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetFree, 0)),
			&runtime.Float{Value: 39.2},
			&runtime.Int{Value: 192},
			&runtime.String{Value: "bar"})))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			runtime.MakeInstruction(parser.OpConstant, 0),
			runtime.MakeInstruction(parser.OpSetGlobal, 0),
			runtime.MakeInstruction(parser.OpConstant, 6),
			runtime.MakeInstruction(parser.OpPop)),
		objectsArray(
			&runtime.Int{Value: 55},
			&runtime.Int{Value: 66},
			&runtime.Int{Value: 77},
			&runtime.Int{Value: 88},
			&runtime.ImmutableMap{
				Value: map[string]runtime.Object{
					"array": &runtime.ImmutableArray{
						Value: []runtime.Object{
							&runtime.Int{Value: 1},
							&runtime.Int{Value: 2},
							&runtime.Int{Value: 3},
							runtime.TrueValue,
							runtime.FalseValue,
							runtime.UndefinedValue,
						},
					},
					"true":  runtime.TrueValue,
					"false": runtime.FalseValue,
					"bytes": &runtime.Bytes{Value: make([]byte, 16)},
					"char":  &runtime.Char{Value: 'Y'},
					"error": &runtime.Error{Value: &runtime.String{
						Value: "some error",
					}},
					"float": &runtime.Float{Value: -19.84},
					"immutable_array": &runtime.ImmutableArray{
						Value: []runtime.Object{
							&runtime.Int{Value: 1},
							&runtime.Int{Value: 2},
							&runtime.Int{Value: 3},
							runtime.TrueValue,
							runtime.FalseValue,
							runtime.UndefinedValue,
						},
					},
					"immutable_map": &runtime.ImmutableMap{
						Value: map[string]runtime.Object{
							"a": &runtime.Int{Value: 1},
							"b": &runtime.Int{Value: 2},
							"c": &runtime.Int{Value: 3},
							"d": runtime.TrueValue,
							"e": runtime.FalseValue,
							"f": runtime.UndefinedValue,
						},
					},
					"int": &runtime.Int{Value: 91},
					"map": &runtime.Map{
						Value: map[string]runtime.Object{
							"a": &runtime.Int{Value: 1},
							"b": &runtime.Int{Value: 2},
							"c": &runtime.Int{Value: 3},
							"d": runtime.TrueValue,
							"e": runtime.FalseValue,
							"f": runtime.UndefinedValue,
						},
					},
					"string":    &runtime.String{Value: "foo bar"},
					"time":      &runtime.Time{Value: time.Now()},
					"undefined": runtime.UndefinedValue,
				},
			},
			compiledFunction(1, 0,
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpSetLocal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetFree, 0),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpGetFree, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpGetLocal, 0),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpSetLocal, 0),
				runtime.MakeInstruction(parser.OpGetFree, 0),
				runtime.MakeInstruction(parser.OpGetLocal, 0),
				runtime.MakeInstruction(parser.OpClosure, 4, 2),
				runtime.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpSetLocal, 0),
				runtime.MakeInstruction(parser.OpGetLocal, 0),
				runtime.MakeInstruction(parser.OpClosure, 5, 1),
				runtime.MakeInstruction(parser.OpReturn, 1))),
		fileSet(srcfile{name: "file1", size: 100},
			srcfile{name: "file2", size: 200})))
}

func TestBytecode_RemoveDuplicates(t *testing.T) {
	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(), objectsArray(
				&runtime.Char{Value: 'y'},
				&runtime.Float{Value: 93.11},
				compiledFunction(1, 0,
					runtime.MakeInstruction(parser.OpConstant, 3),
					runtime.MakeInstruction(parser.OpSetLocal, 0),
					runtime.MakeInstruction(parser.OpGetGlobal, 0),
					runtime.MakeInstruction(parser.OpGetFree, 0)),
				&runtime.Float{Value: 39.2},
				&runtime.Int{Value: 192},
				&runtime.String{Value: "bar"})),
		bytecode(
			concatInsts(), objectsArray(
				&runtime.Char{Value: 'y'},
				&runtime.Float{Value: 93.11},
				compiledFunction(1, 0,
					runtime.MakeInstruction(parser.OpConstant, 3),
					runtime.MakeInstruction(parser.OpSetLocal, 0),
					runtime.MakeInstruction(parser.OpGetGlobal, 0),
					runtime.MakeInstruction(parser.OpGetFree, 0)),
				&runtime.Float{Value: 39.2},
				&runtime.Int{Value: 192},
				&runtime.String{Value: "bar"})))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpConstant, 4),
				runtime.MakeInstruction(parser.OpConstant, 5),
				runtime.MakeInstruction(parser.OpConstant, 6),
				runtime.MakeInstruction(parser.OpConstant, 7),
				runtime.MakeInstruction(parser.OpConstant, 8),
				runtime.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&runtime.Int{Value: 1},
				&runtime.Float{Value: 2.0},
				&runtime.Char{Value: '3'},
				&runtime.String{Value: "four"},
				compiledFunction(1, 0,
					runtime.MakeInstruction(parser.OpConstant, 3),
					runtime.MakeInstruction(parser.OpConstant, 7),
					runtime.MakeInstruction(parser.OpSetLocal, 0),
					runtime.MakeInstruction(parser.OpGetGlobal, 0),
					runtime.MakeInstruction(parser.OpGetFree, 0)),
				&runtime.Int{Value: 1},
				&runtime.Float{Value: 2.0},
				&runtime.Char{Value: '3'},
				&runtime.String{Value: "four"})),
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpConstant, 4),
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&runtime.Int{Value: 1},
				&runtime.Float{Value: 2.0},
				&runtime.Char{Value: '3'},
				&runtime.String{Value: "four"},
				compiledFunction(1, 0,
					runtime.MakeInstruction(parser.OpConstant, 3),
					runtime.MakeInstruction(parser.OpConstant, 2),
					runtime.MakeInstruction(parser.OpSetLocal, 0),
					runtime.MakeInstruction(parser.OpGetGlobal, 0),
					runtime.MakeInstruction(parser.OpGetFree, 0)))))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpConstant, 4)),
			objectsArray(
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2},
				&runtime.Int{Value: 3},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 3})),
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2},
				&runtime.Int{Value: 3})))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			&runtime.Int{Value: 55},
			&runtime.Int{Value: 66},
			&runtime.Int{Value: 77},
			&runtime.Int{Value: 88},
			compiledFunction(1, 0,
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpReturn, 1))))
	require.Equal(t, 7, b.CountObjects())
}

func fileSet(files ...srcfile) *parser.SourceFileSet {
	fileSet := parser.NewFileSet()
	for _, f := range files {
		fileSet.AddFile(f.name, -1, f.size)
	}
	return fileSet
}

func bytecodeFileSet(
	instructions []byte,
	constants []runtime.Object,
	fileSet *parser.SourceFileSet,
) *runtime.Bytecode {
	return &runtime.Bytecode{
		FileSet:      fileSet,
		MainFunction: &runtime.CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func testBytecodeRemoveDuplicates(
	t *testing.T,
	input, expected *runtime.Bytecode,
) {
	input.RemoveDuplicates()

	require.Equal(t, expected.FileSet, input.FileSet)
	require.Equal(t, expected.MainFunction, input.MainFunction)
	require.Equal(t, expected.Constants, input.Constants)
}

func testBytecodeSerialization(t *testing.T, b *runtime.Bytecode) {
	bc, err := b.Marshal()
	require.NoError(t, err)

	r := &runtime.Bytecode{}
	err = r.Unmarshal(bc, nil)
	require.NoError(t, err)

	require.Equal(t, b.FileSet, r.FileSet)
	require.Equal(t, b.MainFunction, r.MainFunction)
	require.Equal(t, b.Constants, r.Constants)
}
