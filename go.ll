define i32 @main() {
; <label>:0
	%1 = icmp eq i32 0, 10
	%2 = inttoptr i32 0 to i32*
	%3 = add i32 0, 1
	store i32 %3, i32* %2
	br i1 %1, label %4, label %0
; <label>:4
	ret i32 0
}

