package main

// This example produces LLVM IR code equivalent to the following C code,
// which implements a pseudo-random number generator.
//
//    int abs(int x);
//
//    int seed = 0;
//
//    // ref: https://en.wikipedia.org/wiki/Linear_congruential_generator
//    //    a = 0x15A4E35
//    //    c = 1
//    int rand(void) {
//       seed = seed*0x15A4E35 + 1;
//       return abs(seed);
//    }

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	// Create convenience types and constants.
	i32 := types.I32
	zero := constant.NewInt(0, i32)
	one := constant.NewInt(1, i32)

	// Create a new LLVM IR module.
	m := ir.NewModule()

	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	mainFunc := m.NewFunction("main", i32)

	// Create instructions and append them to the entry basic block.

	until := constant.NewInt(10, i32)
	i := zero

	forLoop := mainFunc.NewBlock("")
	returnBlock := mainFunc.NewBlock("")
	forLoop.NewCondBr(forLoop.NewICmp(ir.IntEQ, i, until), returnBlock, forLoop)
	ptr := forLoop.NewIntToPtr(i, types.NewPointer(i32))
	forLoop.NewStore(forLoop.NewAdd(i, one), ptr)

	returnBlock.NewRet(i)

	// Print the LLVM IR assembly of the module.
	fmt.Println(m)
}
