
TEXT ·hammingdistance(SB), 4, $0
	mov x+0(FP), rbx // x
	mov y+8(FP), rbp // y

	xor rcx, rcx // set counter to 0
	
checkLowerBit:
	test rbx, rbx
	jne checkLastBits

	test rbp, rbp
	je ending

checkLastBits:
	xor rdx, rdx
	mov rdx, rbx
	and rdx, 1

	xor rdi, rdi
	mov rdi, rbp
	and rdi, 1

	or rdx, rdi
	jz continueLoop
	
	inc rcx // counter++

continueLoop:
	shr rbx, 1 // shift (SH) bits to the right (R)
	shr rbp, 1 // shift (SH) bits to the right (R)

	jmp checkLowerBit

ending:
	mov rcx, ret+16(FP)
	RET
