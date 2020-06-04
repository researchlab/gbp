TEXT main.main(SB) /Users/lihong/workbench/dev/src/github.com/researchlab/gbp/channel/0-1-channel-close.go
  0-1-channel-close.go:8	0x10a0170		65488b0c2530000000		MOVQ GS:0x30, CX					
  0-1-channel-close.go:8	0x10a0179		488d4424c8			LEAQ -0x38(SP), AX					
  0-1-channel-close.go:8	0x10a017e		483b4110			CMPQ 0x10(CX), AX					
  0-1-channel-close.go:8	0x10a0182		0f86d3010000			JBE 0x10a035b						
  0-1-channel-close.go:8	0x10a0188		4881ecb8000000			SUBQ $0xb8, SP						
  0-1-channel-close.go:8	0x10a018f		4889ac24b0000000		MOVQ BP, 0xb0(SP)					
  0-1-channel-close.go:8	0x10a0197		488dac24b0000000		LEAQ 0xb0(SP), BP					
  0-1-channel-close.go:9	0x10a019f		488d051ad80000			LEAQ type.*+54592(SB), AX				
  0-1-channel-close.go:9	0x10a01a6		48890424			MOVQ AX, 0(SP)						
  0-1-channel-close.go:9	0x10a01aa		48c744240802000000		MOVQ $0x2, 0x8(SP)					
  0-1-channel-close.go:9	0x10a01b3		e8e83ff6ff			CALL runtime.makechan(SB)				
  0-1-channel-close.go:9	0x10a01b8		488b442410			MOVQ 0x10(SP), AX					
  0-1-channel-close.go:9	0x10a01bd		4889442438			MOVQ AX, 0x38(SP)					
  0-1-channel-close.go:10	0x10a01c2		48890424			MOVQ AX, 0(SP)						
  0-1-channel-close.go:10	0x10a01c6		488d05c3be0400			LEAQ $f64.fffffffffffffffe+8(SB), AX			
  0-1-channel-close.go:10	0x10a01cd		4889442408			MOVQ AX, 0x8(SP)					
  0-1-channel-close.go:10	0x10a01d2		e8f941f6ff			CALL runtime.chansend1(SB)				
  0-1-channel-close.go:11	0x10a01d7		488b442438			MOVQ 0x38(SP), AX					
  0-1-channel-close.go:11	0x10a01dc		48890424			MOVQ AX, 0(SP)						
  0-1-channel-close.go:11	0x10a01e0		e88b4af6ff			CALL runtime.closechan(SB)				
  0-1-channel-close.go:12	0x10a01e5		488b442438			MOVQ 0x38(SP), AX					
  0-1-channel-close.go:12	0x10a01ea		48890424			MOVQ AX, 0(SP)						
  0-1-channel-close.go:12	0x10a01ee		48c744240800000000		MOVQ $0x0, 0x8(SP)					
  0-1-channel-close.go:12	0x10a01f7		e8344df6ff			CALL runtime.chanrecv2(SB)				
  0-1-channel-close.go:12	0x10a01fc		0fb6442410			MOVZX 0x10(SP), AX					
  0-1-channel-close.go:12	0x10a0201		88442437			MOVB AL, 0x37(SP)					
  0-1-channel-close.go:12	0x10a0205		88442436			MOVB AL, 0x36(SP)					
  0-1-channel-close.go:13	0x10a0209		0fb6442436			MOVZX 0x36(SP), AX					
  0-1-channel-close.go:13	0x10a020e		88442437			MOVB AL, 0x37(SP)					
  0-1-channel-close.go:13	0x10a0212		0f57c0				XORPS X0, X0						
  0-1-channel-close.go:13	0x10a0215		0f11442450			MOVUPS X0, 0x50(SP)					
  0-1-channel-close.go:13	0x10a021a		488d442450			LEAQ 0x50(SP), AX					
  0-1-channel-close.go:13	0x10a021f		4889442448			MOVQ AX, 0x48(SP)					
  0-1-channel-close.go:13	0x10a0224		8400				TESTB AL, 0(AX)						
  0-1-channel-close.go:13	0x10a0226		0fb64c2437			MOVZX 0x37(SP), CX					
  0-1-channel-close.go:13	0x10a022b		488d150ed70000			LEAQ type.*+54464(SB), DX				
  0-1-channel-close.go:13	0x10a0232		4889542450			MOVQ DX, 0x50(SP)					
  0-1-channel-close.go:13	0x10a0237		488d15e2ed0c00			LEAQ runtime.staticbytes(SB), DX			
  0-1-channel-close.go:13	0x10a023e		4801d1				ADDQ DX, CX						
  0-1-channel-close.go:13	0x10a0241		48894c2458			MOVQ CX, 0x58(SP)					
  0-1-channel-close.go:13	0x10a0246		8400				TESTB AL, 0(AX)						
  0-1-channel-close.go:13	0x10a0248		eb00				JMP 0x10a024a						
  0-1-channel-close.go:13	0x10a024a		4889442478			MOVQ AX, 0x78(SP)					
  0-1-channel-close.go:13	0x10a024f		48c784248000000001000000	MOVQ $0x1, 0x80(SP)					
  0-1-channel-close.go:13	0x10a025b		48c784248800000001000000	MOVQ $0x1, 0x88(SP)					
  0-1-channel-close.go:13	0x10a0267		488d050c020300			LEAQ go.string.*+794(SB), AX				
  0-1-channel-close.go:13	0x10a026e		48890424			MOVQ AX, 0(SP)						
  0-1-channel-close.go:13	0x10a0272		48c744240805000000		MOVQ $0x5, 0x8(SP)					
  0-1-channel-close.go:13	0x10a027b		488b442478			MOVQ 0x78(SP), AX					
  0-1-channel-close.go:13	0x10a0280		4889442410			MOVQ AX, 0x10(SP)					
  0-1-channel-close.go:13	0x10a0285		48c744241801000000		MOVQ $0x1, 0x18(SP)					
  0-1-channel-close.go:13	0x10a028e		48c744242001000000		MOVQ $0x1, 0x20(SP)					
  0-1-channel-close.go:13	0x10a0297		e8b4fdffff			CALL log.Printf(SB)					
  0-1-channel-close.go:14	0x10a029c		0fb6442436			MOVZX 0x36(SP), AX					
  0-1-channel-close.go:14	0x10a02a1		88442437			MOVB AL, 0x37(SP)					
  0-1-channel-close.go:14	0x10a02a5		0f57c0				XORPS X0, X0						
  0-1-channel-close.go:14	0x10a02a8		0f11842490000000		MOVUPS X0, 0x90(SP)					
  0-1-channel-close.go:14	0x10a02b0		0f57c0				XORPS X0, X0						
  0-1-channel-close.go:14	0x10a02b3		0f118424a0000000		MOVUPS X0, 0xa0(SP)					
  0-1-channel-close.go:14	0x10a02bb		488d842490000000		LEAQ 0x90(SP), AX					
  0-1-channel-close.go:14	0x10a02c3		4889442440			MOVQ AX, 0x40(SP)					
  0-1-channel-close.go:14	0x10a02c8		8400				TESTB AL, 0(AX)						
  0-1-channel-close.go:14	0x10a02ca		488d0d6fe40000			LEAQ type.*+58048(SB), CX				
  0-1-channel-close.go:14	0x10a02d1		48898c2490000000		MOVQ CX, 0x90(SP)					
  0-1-channel-close.go:14	0x10a02d9		488d0d50c20400			LEAQ runtime/internal/sys.DefaultGoroot.str+376(SB), CX	
  0-1-channel-close.go:14	0x10a02e0		48898c2498000000		MOVQ CX, 0x98(SP)					
  0-1-channel-close.go:14	0x10a02e8		8400				TESTB AL, 0(AX)						
  0-1-channel-close.go:14	0x10a02ea		0fb6442437			MOVZX 0x37(SP), AX					
  0-1-channel-close.go:14	0x10a02ef		488d0d4ad60000			LEAQ type.*+54464(SB), CX				
  0-1-channel-close.go:14	0x10a02f6		48898c24a0000000		MOVQ CX, 0xa0(SP)					
  0-1-channel-close.go:14	0x10a02fe		488d0d1bed0c00			LEAQ runtime.staticbytes(SB), CX			
  0-1-channel-close.go:14	0x10a0305		4801c8				ADDQ CX, AX						
  0-1-channel-close.go:14	0x10a0308		48898424a8000000		MOVQ AX, 0xa8(SP)					
  0-1-channel-close.go:14	0x10a0310		488b442440			MOVQ 0x40(SP), AX					
  0-1-channel-close.go:14	0x10a0315		8400				TESTB AL, 0(AX)						
  0-1-channel-close.go:14	0x10a0317		eb00				JMP 0x10a0319						
  0-1-channel-close.go:14	0x10a0319		4889442460			MOVQ AX, 0x60(SP)					
  0-1-channel-close.go:14	0x10a031e		48c744246802000000		MOVQ $0x2, 0x68(SP)					
  0-1-channel-close.go:14	0x10a0327		48c744247002000000		MOVQ $0x2, 0x70(SP)					
  0-1-channel-close.go:14	0x10a0330		48890424			MOVQ AX, 0(SP)						
  0-1-channel-close.go:14	0x10a0334		48c744240802000000		MOVQ $0x2, 0x8(SP)					
  0-1-channel-close.go:14	0x10a033d		48c744241002000000		MOVQ $0x2, 0x10(SP)					
  0-1-channel-close.go:14	0x10a0346		e8a56affff			CALL fmt.Println(SB)					
  0-1-channel-close.go:15	0x10a034b		488bac24b0000000		MOVQ 0xb0(SP), BP					
  0-1-channel-close.go:15	0x10a0353		4881c4b8000000			ADDQ $0xb8, SP						
  0-1-channel-close.go:15	0x10a035a		c3				RET							
  0-1-channel-close.go:8	0x10a035b		e86090fbff			CALL runtime.morestack_noctxt(SB)			
  0-1-channel-close.go:8	0x10a0360		e90bfeffff			JMP main.main(SB)					
  :-1				0x10a0365		cc				INT $0x3						
  :-1				0x10a0366		cc				INT $0x3						
  :-1				0x10a0367		cc				INT $0x3						
  :-1				0x10a0368		cc				INT $0x3						
  :-1				0x10a0369		cc				INT $0x3						
  :-1				0x10a036a		cc				INT $0x3						
  :-1				0x10a036b		cc				INT $0x3						
  :-1				0x10a036c		cc				INT $0x3						
  :-1				0x10a036d		cc				INT $0x3						
  :-1				0x10a036e		cc				INT $0x3						
  :-1				0x10a036f		cc				INT $0x3						
