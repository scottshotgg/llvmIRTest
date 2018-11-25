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
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

func main() {
	// Create a new LLVM IR module.
	m := ir.NewModule()

	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	mainFunc := m.NewFunc("main", types.I32)
	returnBlock := ir.NewBlock("")

	For(mainFunc, returnBlock, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 10))

	mainFunc.Blocks = append(mainFunc.Blocks, returnBlock)
	returnBlock.NewRet(constant.NewInt(types.I32, 0))

	// // Print the LLVM IR assembly of the module.
	fmt.Println(m)
}

func For(fromFunction *ir.Function, targetBlock *ir.BasicBlock, start, end *constant.Int /* array of token.Value here */) {
	allocBlock := fromFunction.NewBlock("")
	compareBlock := fromFunction.NewBlock("")
	incrementBlock := fromFunction.NewBlock("")

	startVar := allocBlock.NewAlloca(types.I32)
	indexVar := allocBlock.NewAlloca(types.I32)
	allocBlock.NewStore(constant.NewInt(types.I32, 0), startVar)
	allocBlock.NewStore(constant.NewInt(types.I32, 0), indexVar)
	allocBlock.NewBr(compareBlock)

	loadedIndexVar := compareBlock.NewLoad(indexVar)
	breakOut := compareBlock.NewICmp(enum.IPredSLT, loadedIndexVar, end)
	compareBlock.NewCondBr(breakOut, incrementBlock, targetBlock)

	added := incrementBlock.NewAdd(loadedIndexVar, constant.NewInt(types.I32, 1))
	incrementBlock.NewStore(added, indexVar)
	incrementBlock.NewBr(compareBlock)
}
