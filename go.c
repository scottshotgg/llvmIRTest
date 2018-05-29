	.section	__TEXT,__text,regular,pure_instructions
	.macosx_version_min 10, 12
	.globl	_main                   ## -- Begin function main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	imull	$22695477, _seed, %eax  ## imm = 0x15A4E35
	incl	%eax
	movl	%eax, _seed
	xorl	%eax, %eax
	retl
	.cfi_endproc
                                        ## -- End function
	.globl	_seed                   ## @seed
.zerofill __DATA,__common,_seed,4,2

.subsections_via_symbols
