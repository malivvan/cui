package runtime_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/malivvan/cui/runtime"
	"github.com/malivvan/cui/runtime/parser"
	"github.com/malivvan/cui/runtime/require"
)

func TestCompiler_Compile(t *testing.T) {
	expectCompile(t, `1 + 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1; 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 - 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 12),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 * 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 13),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `2 / 1`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 14),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(2),
				intObject(1))))

	expectCompile(t, `true`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpTrue),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `false`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpFalse),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `1 > 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 39),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 < 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 39),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(2),
				intObject(1))))

	expectCompile(t, `1 >= 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 44),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 <= 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 44),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(2),
				intObject(1))))

	expectCompile(t, `1 == 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpEqual),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `1 != 2`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpNotEqual),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `true == false`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpTrue),
				runtime.MakeInstruction(parser.OpFalse),
				runtime.MakeInstruction(parser.OpEqual),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `true != false`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpTrue),
				runtime.MakeInstruction(parser.OpFalse),
				runtime.MakeInstruction(parser.OpNotEqual),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `-1`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpMinus),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1))))

	expectCompile(t, `!true`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpTrue),
				runtime.MakeInstruction(parser.OpLNot),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `if true { 10 }; 3333`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpTrue),         // 0000
				runtime.MakeInstruction(parser.OpJumpFalsy, 8), // 0001
				runtime.MakeInstruction(parser.OpConstant, 0),  // 0004
				runtime.MakeInstruction(parser.OpPop),          // 0007
				runtime.MakeInstruction(parser.OpConstant, 1),  // 0008
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)), // 0011
			objectsArray(
				intObject(10),
				intObject(3333))))

	expectCompile(t, `if (true) { 10 } else { 20 }; 3333;`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpTrue),          // 0000
				runtime.MakeInstruction(parser.OpJumpFalsy, 11), // 0001
				runtime.MakeInstruction(parser.OpConstant, 0),   // 0004
				runtime.MakeInstruction(parser.OpPop),           // 0007
				runtime.MakeInstruction(parser.OpJump, 15),      // 0008
				runtime.MakeInstruction(parser.OpConstant, 1),   // 0011
				runtime.MakeInstruction(parser.OpPop),           // 0014
				runtime.MakeInstruction(parser.OpConstant, 2),   // 0015
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)), // 0018
			objectsArray(
				intObject(10),
				intObject(20),
				intObject(3333))))

	expectCompile(t, `"kami"`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("kami"))))

	expectCompile(t, `"ka" + "mi"`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("ka"),
				stringObject("mi"))))

	expectCompile(t, `a := 1; b := 2; a += b`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpSetGlobal, 1),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `a := 1; b := 2; a /= b`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpSetGlobal, 1),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 14),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2))))

	expectCompile(t, `[]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpArray, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `[1, 2, 3]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpArray, 3),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3))))

	expectCompile(t, `[1 + 2, 3 - 4, 5 * 6]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpBinaryOp, 12),
				runtime.MakeInstruction(parser.OpConstant, 4),
				runtime.MakeInstruction(parser.OpConstant, 5),
				runtime.MakeInstruction(parser.OpBinaryOp, 13),
				runtime.MakeInstruction(parser.OpArray, 3),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3),
				intObject(4),
				intObject(5),
				intObject(6))))

	expectCompile(t, `{}`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpMap, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `{a: 2, b: 4, c: 6}`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpConstant, 4),
				runtime.MakeInstruction(parser.OpConstant, 5),
				runtime.MakeInstruction(parser.OpMap, 6),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("a"),
				intObject(2),
				stringObject("b"),
				intObject(4),
				stringObject("c"),
				intObject(6))))

	expectCompile(t, `{a: 2 + 3, b: 5 * 6}`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpConstant, 4),
				runtime.MakeInstruction(parser.OpConstant, 5),
				runtime.MakeInstruction(parser.OpBinaryOp, 13),
				runtime.MakeInstruction(parser.OpMap, 4),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("a"),
				intObject(2),
				intObject(3),
				stringObject("b"),
				intObject(5),
				intObject(6))))

	expectCompile(t, `[1, 2, 3][1 + 1]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpArray, 3),
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpIndex),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3))))

	expectCompile(t, `{a: 2}[2 - 1]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpMap, 2),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpBinaryOp, 12),
				runtime.MakeInstruction(parser.OpIndex),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				stringObject("a"),
				intObject(2),
				intObject(1))))

	expectCompile(t, `[1, 2, 3][:]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpArray, 3),
				runtime.MakeInstruction(parser.OpNull),
				runtime.MakeInstruction(parser.OpNull),
				runtime.MakeInstruction(parser.OpSliceIndex),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3))))

	expectCompile(t, `[1, 2, 3][0 : 2]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpArray, 3),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpSliceIndex),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3),
				intObject(0))))

	expectCompile(t, `[1, 2, 3][:2]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpArray, 3),
				runtime.MakeInstruction(parser.OpNull),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpSliceIndex),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3))))

	expectCompile(t, `[1, 2, 3][0:]`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpArray, 3),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpNull),
				runtime.MakeInstruction(parser.OpSliceIndex),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3),
				intObject(0))))

	expectCompile(t, `f1 := func(a) { return a }; f1([1, 2]...);`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpArray, 2),
				runtime.MakeInstruction(parser.OpCall, 1, 1),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)),
				intObject(1),
				intObject(2))))

	expectCompile(t, `func() { return 5 + 10 }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(5),
				intObject(10),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpConstant, 1),
					runtime.MakeInstruction(parser.OpBinaryOp, 11),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `func() { 5 + 10 }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(5),
				intObject(10),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpConstant, 1),
					runtime.MakeInstruction(parser.OpBinaryOp, 11),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { 1; 2 }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpConstant, 1),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { 1; return 2 }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpConstant, 1),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `func() { if(true) { return 1 } else { return 2 } }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpTrue),         // 0000
					runtime.MakeInstruction(parser.OpJumpFalsy, 9), // 0001
					runtime.MakeInstruction(parser.OpConstant, 0),  // 0004
					runtime.MakeInstruction(parser.OpReturn, 1),    // 0007
					runtime.MakeInstruction(parser.OpConstant, 1),  // 0009
					runtime.MakeInstruction(parser.OpReturn, 1))))) // 0012

	expectCompile(t, `func() { 1; if(true) { 2 } else { 3 }; 4 }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 4),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(1),
				intObject(2),
				intObject(3),
				intObject(4),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),   // 0000
					runtime.MakeInstruction(parser.OpPop),           // 0003
					runtime.MakeInstruction(parser.OpTrue),          // 0004
					runtime.MakeInstruction(parser.OpJumpFalsy, 15), // 0005
					runtime.MakeInstruction(parser.OpConstant, 1),   // 0008
					runtime.MakeInstruction(parser.OpPop),           // 0011
					runtime.MakeInstruction(parser.OpJump, 19),      // 0012
					runtime.MakeInstruction(parser.OpConstant, 2),   // 0015
					runtime.MakeInstruction(parser.OpPop),           // 0018
					runtime.MakeInstruction(parser.OpConstant, 3),   // 0019
					runtime.MakeInstruction(parser.OpPop),           // 0022
					runtime.MakeInstruction(parser.OpReturn, 0)))))  // 0023

	expectCompile(t, `func() { }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { 24 }()`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpCall, 0, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(24),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { return 24 }()`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpCall, 0, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(24),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `noArg := func() { 24 }; noArg();`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpCall, 0, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(24),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `noArg := func() { return 24 }; noArg();`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpCall, 0, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(24),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `n := 55; func() { n };`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpGetGlobal, 0),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `func() { n := 55; return n }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				compiledFunction(1, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpDefineLocal, 0),
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `func() { a := 55; b := 77; return a + b }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				intObject(77),
				compiledFunction(2, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpDefineLocal, 0),
					runtime.MakeInstruction(parser.OpConstant, 1),
					runtime.MakeInstruction(parser.OpDefineLocal, 1),
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpGetLocal, 1),
					runtime.MakeInstruction(parser.OpBinaryOp, 11),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `f1 := func(a) { return a }; f1(24);`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpCall, 1, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)),
				intObject(24))))

	expectCompile(t, `varTest := func(...a) { return a }; varTest(1,2,3);`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpCall, 3, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)),
				intObject(1), intObject(2), intObject(3))))

	expectCompile(t, `f1 := func(a, b, c) { a; b; return c; }; f1(24, 25, 26);`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpCall, 3, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(3, 3,
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpGetLocal, 1),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpGetLocal, 2),
					runtime.MakeInstruction(parser.OpReturn, 1)),
				intObject(24),
				intObject(25),
				intObject(26))))

	expectCompile(t, `func() { n := 55; n = 23; return n }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				intObject(23),
				compiledFunction(1, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpDefineLocal, 0),
					runtime.MakeInstruction(parser.OpConstant, 1),
					runtime.MakeInstruction(parser.OpSetLocal, 0),
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)))))
	expectCompile(t, `len([]);`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpGetBuiltin, 0),
				runtime.MakeInstruction(parser.OpArray, 0),
				runtime.MakeInstruction(parser.OpCall, 1, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `func() { return len([]) }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpGetBuiltin, 0),
					runtime.MakeInstruction(parser.OpArray, 0),
					runtime.MakeInstruction(parser.OpCall, 1, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `func(a) { func(b) { return a + b } }`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					runtime.MakeInstruction(parser.OpGetFree, 0),
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpBinaryOp, 11),
					runtime.MakeInstruction(parser.OpReturn, 1)),
				compiledFunction(1, 1,
					runtime.MakeInstruction(parser.OpGetLocalPtr, 0),
					runtime.MakeInstruction(parser.OpClosure, 0, 1),
					runtime.MakeInstruction(parser.OpPop),
					runtime.MakeInstruction(parser.OpReturn, 0)))))

	expectCompile(t, `
func(a) {
	return func(b) {
		return func(c) {
			return a + b + c
		}
	}
}`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				compiledFunction(1, 1,
					runtime.MakeInstruction(parser.OpGetFree, 0),
					runtime.MakeInstruction(parser.OpGetFree, 1),
					runtime.MakeInstruction(parser.OpBinaryOp, 11),
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpBinaryOp, 11),
					runtime.MakeInstruction(parser.OpReturn, 1)),
				compiledFunction(1, 1,
					runtime.MakeInstruction(parser.OpGetFreePtr, 0),
					runtime.MakeInstruction(parser.OpGetLocalPtr, 0),
					runtime.MakeInstruction(parser.OpClosure, 0, 2),
					runtime.MakeInstruction(parser.OpReturn, 1)),
				compiledFunction(1, 1,
					runtime.MakeInstruction(parser.OpGetLocalPtr, 0),
					runtime.MakeInstruction(parser.OpClosure, 1, 1),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `
g := 55;

func() {
	a := 66;

	return func() {
		b := 77;

		return func() {
			c := 88;

			return g + a + b + c;
		}
	}
}`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 6),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(55),
				intObject(66),
				intObject(77),
				intObject(88),
				compiledFunction(1, 0,
					runtime.MakeInstruction(parser.OpConstant, 3),
					runtime.MakeInstruction(parser.OpDefineLocal, 0),
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
					runtime.MakeInstruction(parser.OpDefineLocal, 0),
					runtime.MakeInstruction(parser.OpGetFreePtr, 0),
					runtime.MakeInstruction(parser.OpGetLocalPtr, 0),
					runtime.MakeInstruction(parser.OpClosure, 4, 2),
					runtime.MakeInstruction(parser.OpReturn, 1)),
				compiledFunction(1, 0,
					runtime.MakeInstruction(parser.OpConstant, 1),
					runtime.MakeInstruction(parser.OpDefineLocal, 0),
					runtime.MakeInstruction(parser.OpGetLocalPtr, 0),
					runtime.MakeInstruction(parser.OpClosure, 5, 1),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `for i:=0; i<10; i++ {}`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpBinaryOp, 39),
				runtime.MakeInstruction(parser.OpJumpFalsy, 31),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpJump, 6),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(0),
				intObject(10),
				intObject(1))))

	expectCompile(t, `m := {}; for k, v in m {}`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpMap, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpIteratorInit),
				runtime.MakeInstruction(parser.OpSetGlobal, 1),
				runtime.MakeInstruction(parser.OpGetGlobal, 1),
				runtime.MakeInstruction(parser.OpIteratorNext),
				runtime.MakeInstruction(parser.OpJumpFalsy, 37),
				runtime.MakeInstruction(parser.OpGetGlobal, 1),
				runtime.MakeInstruction(parser.OpIteratorKey),
				runtime.MakeInstruction(parser.OpSetGlobal, 2),
				runtime.MakeInstruction(parser.OpGetGlobal, 1),
				runtime.MakeInstruction(parser.OpIteratorValue),
				runtime.MakeInstruction(parser.OpSetGlobal, 3),
				runtime.MakeInstruction(parser.OpJump, 13),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray()))

	expectCompile(t, `a := 0; a == 0 && a != 1 || a < 1`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpSetGlobal, 0),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpEqual),
				runtime.MakeInstruction(parser.OpAndJump, 23),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpNotEqual),
				runtime.MakeInstruction(parser.OpOrJump, 34),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpGetGlobal, 0),
				runtime.MakeInstruction(parser.OpBinaryOp, 39),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(0),
				intObject(1))))

	// unknown module name
	expectCompileError(t, `import("user1")`, "module 'user1' not found")

	// too many errors
	expectCompileError(t, `
r["x"] = {
    @a:1,
    @b:1,
    @c:1,
    @d:1,
    @e:1,
    @f:1,
    @g:1,
    @h:1,
    @i:1,
    @j:1,
    @k:1
}
`, "Parse Error: illegal character U+0040 '@'\n\tat test:3:5 (and 10 more errors)")

	expectCompileError(t, `import("")`, "empty module name")

	// https://github.com/malivvan/vv/issues/314
	expectCompileError(t, `
(func() {
	fn := fn()
})()	
`, "unresolved reference 'fn")
}

func TestCompilerErrorReport(t *testing.T) {
	expectCompileError(t, `import("user1")`,
		"Compile Error: module 'user1' not found\n\tat test:1:1")

	expectCompileError(t, `a = 1`,
		"Compile Error: unresolved reference 'a'\n\tat test:1:1")
	expectCompileError(t, `a, b := 1, 2`,
		"Compile Error: tuple assignment not allowed\n\tat test:1:1")
	expectCompileError(t, `a.b := 1`,
		"not allowed with selector")
	expectCompileError(t, `a:=1; a:=3`,
		"Compile Error: 'a' redeclared in this block\n\tat test:1:7")

	expectCompileError(t, `return 5`,
		"Compile Error: return not allowed outside function\n\tat test:1:1")
	expectCompileError(t, `func() { break }`,
		"Compile Error: break not allowed outside loop\n\tat test:1:10")
	expectCompileError(t, `func() { continue }`,
		"Compile Error: continue not allowed outside loop\n\tat test:1:10")
	expectCompileError(t, `func() { export 5 }`,
		"Compile Error: export not allowed inside function\n\tat test:1:10")
}

func TestCompilerDeadCode(t *testing.T) {
	expectCompile(t, `
func() {
	a := 4
	return a

	b := 5 // dead code from here
	c := a
	return b
}`,
		bytecode(
			concatInsts(
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpSuspend)),
			objectsArray(
				intObject(4),
				intObject(5),
				compiledFunction(0, 0,
					runtime.MakeInstruction(parser.OpConstant, 0),
					runtime.MakeInstruction(parser.OpDefineLocal, 0),
					runtime.MakeInstruction(parser.OpGetLocal, 0),
					runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `
func() {
	if true {
		return 5
		a := 4  // dead code from here
		b := a
		return b
	} else {
		return 4
		c := 5  // dead code from here
		d := c
		return d
	}
}`, bytecode(
		concatInsts(
			runtime.MakeInstruction(parser.OpConstant, 2),
			runtime.MakeInstruction(parser.OpPop),
			runtime.MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(5),
			intObject(4),
			compiledFunction(0, 0,
				runtime.MakeInstruction(parser.OpTrue),
				runtime.MakeInstruction(parser.OpJumpFalsy, 9),
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpReturn, 1),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `
func() {
	a := 1
	for {
		if a == 5 {
			return 10
		}
		5 + 5
		return 20
		b := a
		return b
	}
}`, bytecode(
		concatInsts(
			runtime.MakeInstruction(parser.OpConstant, 4),
			runtime.MakeInstruction(parser.OpPop),
			runtime.MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(1),
			intObject(5),
			intObject(10),
			intObject(20),
			compiledFunction(0, 0,
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpDefineLocal, 0),
				runtime.MakeInstruction(parser.OpGetLocal, 0),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpEqual),
				runtime.MakeInstruction(parser.OpJumpFalsy, 19),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpReturn, 1),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpBinaryOp, 11),
				runtime.MakeInstruction(parser.OpPop),
				runtime.MakeInstruction(parser.OpConstant, 3),
				runtime.MakeInstruction(parser.OpReturn, 1)))))

	expectCompile(t, `
func() {
	if true {
		return 5
		a := 4  // dead code from here
		b := a
		return b
	} else {
		return 4
		c := 5  // dead code from here
		d := c
		return d
	}
}`, bytecode(
		concatInsts(
			runtime.MakeInstruction(parser.OpConstant, 2),
			runtime.MakeInstruction(parser.OpPop),
			runtime.MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(5),
			intObject(4),
			compiledFunction(0, 0,
				runtime.MakeInstruction(parser.OpTrue),
				runtime.MakeInstruction(parser.OpJumpFalsy, 9),
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpReturn, 1),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpReturn, 1)))))
}

func TestCompilerScopes(t *testing.T) {
	expectCompile(t, `
if a := 1; a {
    a = 2
	b := a
} else {
    a = 3
	b := a
}`, bytecode(
		concatInsts(
			runtime.MakeInstruction(parser.OpConstant, 0),
			runtime.MakeInstruction(parser.OpSetGlobal, 0),
			runtime.MakeInstruction(parser.OpGetGlobal, 0),
			runtime.MakeInstruction(parser.OpJumpFalsy, 27),
			runtime.MakeInstruction(parser.OpConstant, 1),
			runtime.MakeInstruction(parser.OpSetGlobal, 0),
			runtime.MakeInstruction(parser.OpGetGlobal, 0),
			runtime.MakeInstruction(parser.OpSetGlobal, 1),
			runtime.MakeInstruction(parser.OpJump, 39),
			runtime.MakeInstruction(parser.OpConstant, 2),
			runtime.MakeInstruction(parser.OpSetGlobal, 0),
			runtime.MakeInstruction(parser.OpGetGlobal, 0),
			runtime.MakeInstruction(parser.OpSetGlobal, 2),
			runtime.MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(1),
			intObject(2),
			intObject(3))))

	expectCompile(t, `
func() {
	if a := 1; a {
    	a = 2
		b := a
	} else {
    	a = 3
		b := a
	}
}`, bytecode(
		concatInsts(
			runtime.MakeInstruction(parser.OpConstant, 3),
			runtime.MakeInstruction(parser.OpPop),
			runtime.MakeInstruction(parser.OpSuspend)),
		objectsArray(
			intObject(1),
			intObject(2),
			intObject(3),
			compiledFunction(0, 0,
				runtime.MakeInstruction(parser.OpConstant, 0),
				runtime.MakeInstruction(parser.OpDefineLocal, 0),
				runtime.MakeInstruction(parser.OpGetLocal, 0),
				runtime.MakeInstruction(parser.OpJumpFalsy, 22),
				runtime.MakeInstruction(parser.OpConstant, 1),
				runtime.MakeInstruction(parser.OpSetLocal, 0),
				runtime.MakeInstruction(parser.OpGetLocal, 0),
				runtime.MakeInstruction(parser.OpDefineLocal, 1),
				runtime.MakeInstruction(parser.OpJump, 31),
				runtime.MakeInstruction(parser.OpConstant, 2),
				runtime.MakeInstruction(parser.OpSetLocal, 0),
				runtime.MakeInstruction(parser.OpGetLocal, 0),
				runtime.MakeInstruction(parser.OpDefineLocal, 1),
				runtime.MakeInstruction(parser.OpReturn, 0)))))
}

func concatInsts(instructions ...[]byte) []byte {
	var concat []byte
	for _, i := range instructions {
		concat = append(concat, i...)
	}
	return concat
}

func bytecode(
	instructions []byte,
	constants []runtime.Object,
) *runtime.Bytecode {
	return &runtime.Bytecode{
		FileSet:      parser.NewFileSet(),
		MainFunction: &runtime.CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func expectCompile(
	t *testing.T,
	input string,
	expected *runtime.Bytecode,
) {
	actual, trace, err := traceCompile(input, nil)

	var ok bool
	defer func() {
		if !ok {
			for _, tr := range trace {
				t.Log(tr)
			}
		}
	}()

	require.NoError(t, err)
	equalBytecode(t, expected, actual)
	ok = true
}

func expectCompileError(t *testing.T, input, expected string) {
	_, trace, err := traceCompile(input, nil)

	var ok bool
	defer func() {
		if !ok {
			for _, tr := range trace {
				t.Log(tr)
			}
		}
	}()

	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), expected),
		"expected error string: %s, got: %s", expected, err.Error())
	ok = true
}

func equalBytecode(t *testing.T, expected, actual *runtime.Bytecode) {
	require.Equal(t, expected.MainFunction, actual.MainFunction)
	equalConstants(t, expected.Constants, actual.Constants)
}

func equalConstants(t *testing.T, expected, actual []runtime.Object) {
	require.Equal(t, len(expected), len(actual))
	for i := 0; i < len(expected); i++ {
		require.Equal(t, expected[i], actual[i])
	}
}

type compileTracer struct {
	Out []string
}

func (o *compileTracer) Write(p []byte) (n int, err error) {
	o.Out = append(o.Out, string(p))
	return len(p), nil
}

func traceCompile(
	input string,
	symbols map[string]runtime.Object,
) (res *runtime.Bytecode, trace []string, err error) {
	fileSet := parser.NewFileSet()
	file := fileSet.AddFile("test", -1, len(input))

	p := parser.NewParser(file, []byte(input), nil)

	symTable := runtime.NewSymbolTable()
	for name := range symbols {
		symTable.Define(name)
	}
	for idx, fn := range runtime.GetAllBuiltinFunctions() {
		symTable.DefineBuiltin(idx, fn.Name)
	}

	tr := &compileTracer{}
	c := runtime.NewCompiler(file, symTable, nil, nil, tr)
	parsed, err := p.ParseFile()
	if err != nil {
		return
	}

	err = c.Compile(parsed)
	res = c.Bytecode()
	res.RemoveDuplicates()
	{
		trace = append(trace, fmt.Sprintf("Compiler Trace:\n%s",
			strings.Join(tr.Out, "")))
		trace = append(trace, fmt.Sprintf("Compiled Constants:\n%s",
			strings.Join(res.FormatConstants(), "\n")))
		trace = append(trace, fmt.Sprintf("Compiled Instructions:\n%s\n",
			strings.Join(res.FormatInstructions(), "\n")))
	}
	if err != nil {
		return
	}
	return
}

func objectsArray(o ...runtime.Object) []runtime.Object {
	return o
}

func intObject(v int64) *runtime.Int {
	return &runtime.Int{Value: v}
}

func stringObject(v string) *runtime.String {
	return &runtime.String{Value: v}
}

func compiledFunction(
	numLocals, numParams int,
	insts ...[]byte,
) *runtime.CompiledFunction {
	return &runtime.CompiledFunction{
		Instructions:  concatInsts(insts...),
		NumLocals:     numLocals,
		NumParameters: numParams,
	}
}
