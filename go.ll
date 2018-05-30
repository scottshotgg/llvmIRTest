define i32 @main() {
; <label>:0
	%1 = alloca i32
	%2 = alloca i32
	store i32 0, i32* %1
	store i32 0, i32* %2
	br label %3
; <label>:3
	%4 = load i32, i32* %2
	%5 = icmp slt i32 %4, 10
	br i1 %5, label %6, label %8
; <label>:6
	%7 = add i32 %4, 1
	store i32 %7, i32* %2
	br label %3
; <label>:8
	ret i32 0
}

