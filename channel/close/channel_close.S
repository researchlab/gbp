  channel_close.go:5	0x1093e60		65488b0c2530000000	MOVQ GS:0x30, CX			
  channel_close.go:5	0x1093e69		488d4424f8		LEAQ -0x8(SP), AX			
  channel_close.go:5	0x1093e6e		483b4110		CMPQ 0x10(CX), AX			
  channel_close.go:5	0x1093e72		0f8665010000		JBE 0x1093fdd				
  channel_close.go:5	0x1093e78		4881ec88000000		SUBQ $0x88, SP				
  channel_close.go:5	0x1093e7f		4889ac2480000000	MOVQ BP, 0x80(SP)			
  channel_close.go:5	0x1093e87		488dac2480000000	LEAQ 0x80(SP), BP			
  channel_close.go:6	0x1093e8f		488d05caf50000		LEAQ runtime.types+62464(SB), AX	
  channel_close.go:6	0x1093e96		48890424		MOVQ AX, 0(SP)				
  channel_close.go:6	0x1093e9a		48c744240802000000	MOVQ $0x2, 0x8(SP)			
  channel_close.go:6	0x1093ea3		e8b800f7ff		CALL runtime.makechan(SB)		
  channel_close.go:6	0x1093ea8		488b442410		MOVQ 0x10(SP), AX			
  channel_close.go:6	0x1093ead		4889442440		MOVQ AX, 0x40(SP)			
  channel_close.go:7	0x1093eb2		48890424		MOVQ AX, 0(SP)				
  channel_close.go:7	0x1093eb6		488d05535a0400		LEAQ main.statictmp_0(SB), AX		
  channel_close.go:7	0x1093ebd		4889442408		MOVQ AX, 0x8(SP)			
  channel_close.go:7	0x1093ec2		e80903f7ff		CALL runtime.chansend1(SB)		
  channel_close.go:8	0x1093ec7		488b442440		MOVQ 0x40(SP), AX			
  channel_close.go:8	0x1093ecc		48890424		MOVQ AX, 0(SP)				
  channel_close.go:8	0x1093ed0		e86b0bf7ff		CALL runtime.closechan(SB)		
  channel_close.go:9	0x1093ed5		488b442440		MOVQ 0x40(SP), AX			
  channel_close.go:9	0x1093eda		4889442448		MOVQ AX, 0x48(SP)			
  channel_close.go:9	0x1093edf		eb00			JMP 0x1093ee1				
  channel_close.go:9	0x1093ee1		488b442448		MOVQ 0x48(SP), AX			
  channel_close.go:9	0x1093ee6		48890424		MOVQ AX, 0(SP)				
  channel_close.go:9	0x1093eea		488d442430		LEAQ 0x30(SP), AX			
  channel_close.go:9	0x1093eef		4889442408		MOVQ AX, 0x8(SP)			
  channel_close.go:9	0x1093ef4		e8170ef7ff		CALL runtime.chanrecv2(SB)		
  channel_close.go:9	0x1093ef9		0fb6442410		MOVZX 0x10(SP), AX			
  channel_close.go:9	0x1093efe		8844242d		MOVB AL, 0x2d(SP)			
  channel_close.go:9	0x1093f02		84c0			TESTL AL, AL				
  channel_close.go:9	0x1093f04		7502			JNE 0x1093f08				
  channel_close.go:9	0x1093f06		eb0d			JMP 0x1093f15				
  channel_close.go:9	0x1093f08		48c744243000000000	MOVQ $0x0, 0x30(SP)			
  channel_close.go:9	0x1093f11		eb00			JMP 0x1093f13				
  channel_close.go:9	0x1093f13		ebcc			JMP 0x1093ee1				
  channel_close.go:13	0x1093f15		eb00			JMP 0x1093f17				
  channel_close.go:13	0x1093f17		eb00			JMP 0x1093f19				
  channel_close.go:13	0x1093f19		eb00			JMP 0x1093f1b				
  channel_close.go:16	0x1093f1b		488b442440		MOVQ 0x40(SP), AX			
  channel_close.go:16	0x1093f20		48890424		MOVQ AX, 0(SP)				
  channel_close.go:16	0x1093f24		488d442438		LEAQ 0x38(SP), AX			
  channel_close.go:16	0x1093f29		4889442408		MOVQ AX, 0x8(SP)			
  channel_close.go:16	0x1093f2e		e8dd0df7ff		CALL runtime.chanrecv2(SB)		
  channel_close.go:16	0x1093f33		0fb6442410		MOVZX 0x10(SP), AX			
  channel_close.go:16	0x1093f38		8844242f		MOVB AL, 0x2f(SP)			
  channel_close.go:16	0x1093f3c		8844242c		MOVB AL, 0x2c(SP)			
  channel_close.go:17	0x1093f40		0fb644242c		MOVZX 0x2c(SP), AX			
  channel_close.go:17	0x1093f45		8844242e		MOVB AL, 0x2e(SP)			
  channel_close.go:17	0x1093f49		0f57c0			XORPS X0, X0				
  channel_close.go:17	0x1093f4c		0f11442458		MOVUPS X0, 0x58(SP)			
  channel_close.go:17	0x1093f51		488d442458		LEAQ 0x58(SP), AX			
  channel_close.go:17	0x1093f56		4889442450		MOVQ AX, 0x50(SP)			
  channel_close.go:17	0x1093f5b		8400			TESTB AL, 0(AX)				
  channel_close.go:17	0x1093f5d		0fb64c242e		MOVZX 0x2e(SP), CX			
  channel_close.go:17	0x1093f62		488d1577f40000		LEAQ runtime.types+62336(SB), DX	
  channel_close.go:17	0x1093f69		4889542458		MOVQ DX, 0x58(SP)			
  channel_close.go:17	0x1093f6e		488d150b390c00		LEAQ runtime.staticbytes(SB), DX	
  channel_close.go:17	0x1093f75		4801d1			ADDQ DX, CX				
  channel_close.go:17	0x1093f78		48894c2460		MOVQ CX, 0x60(SP)			
  channel_close.go:17	0x1093f7d		8400			TESTB AL, 0(AX)				
  channel_close.go:17	0x1093f7f		eb00			JMP 0x1093f81				
  channel_close.go:17	0x1093f81		4889442468		MOVQ AX, 0x68(SP)			
  channel_close.go:17	0x1093f86		48c744247001000000	MOVQ $0x1, 0x70(SP)			
  channel_close.go:17	0x1093f8f		48c744247801000000	MOVQ $0x1, 0x78(SP)			
  channel_close.go:17	0x1093f98		488d0582140300		LEAQ go.string.*+833(SB), AX		
  channel_close.go:17	0x1093f9f		48890424		MOVQ AX, 0(SP)				
  channel_close.go:17	0x1093fa3		48c744240805000000	MOVQ $0x5, 0x8(SP)			
  channel_close.go:17	0x1093fac		488b442468		MOVQ 0x68(SP), AX			
  channel_close.go:17	0x1093fb1		4889442410		MOVQ AX, 0x10(SP)			
  channel_close.go:17	0x1093fb6		48c744241801000000	MOVQ $0x1, 0x18(SP)			
  channel_close.go:17	0x1093fbf		48c744242001000000	MOVQ $0x1, 0x20(SP)			
  channel_close.go:17	0x1093fc8		e823fdffff		CALL log.Printf(SB)			
  channel_close.go:18	0x1093fcd		488bac2480000000	MOVQ 0x80(SP), BP			
  channel_close.go:18	0x1093fd5		4881c488000000		ADDQ $0x88, SP				
  channel_close.go:18	0x1093fdc		c3			RET					
  channel_close.go:5	0x1093fdd		e81eacfbff		CALL runtime.morestack_noctxt(SB)	
  channel_close.go:5	0x1093fe2		e979feffff		JMP main.main(SB)			
  :-1			0x1093fe7		cc			INT $0x3				
  :-1			0x1093fe8		cc			INT $0x3				
  :-1			0x1093fe9		cc			INT $0x3				
  :-1			0x1093fea		cc			INT $0x3				
  :-1			0x1093feb		cc			INT $0x3				
  :-1			0x1093fec		cc			INT $0x3				
  :-1			0x1093fed		cc			INT $0x3				
  :-1			0x1093fee		cc			INT $0x3				
  :-1			0x1093fef		cc			INT $0x3				
