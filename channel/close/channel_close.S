TEXT main.main(SB) /Users/lihong/workbench/dev/src/github.com/researchlab/experiments/channel/close/channel_close.go
  channel_close.go:3	0x104e060		65488b0c2530000000	MOVQ GS:0x30, CX			
  channel_close.go:3	0x104e069		483b6110		CMPQ 0x10(CX), SP			
  channel_close.go:3	0x104e06d		0f86a4000000		JBE 0x104e117				
  channel_close.go:3	0x104e073		4883ec40		SUBQ $0x40, SP				
  channel_close.go:3	0x104e077		48896c2438		MOVQ BP, 0x38(SP)			
  channel_close.go:3	0x104e07c		488d6c2438		LEAQ 0x38(SP), BP			
  channel_close.go:4	0x104e081		488d05b89a0000		LEAQ runtime.types+39328(SB), AX	
  channel_close.go:4	0x104e088		48890424		MOVQ AX, 0(SP)				
  channel_close.go:4	0x104e08c		48c744240802000000	MOVQ $0x2, 0x8(SP)			
  channel_close.go:4	0x104e095		e8165cfbff		CALL runtime.makechan(SB)		
  channel_close.go:4	0x104e09a		488b442410		MOVQ 0x10(SP), AX			
  channel_close.go:4	0x104e09f		4889442428		MOVQ AX, 0x28(SP)			
  channel_close.go:5	0x104e0a4		48890424		MOVQ AX, 0(SP)				
  channel_close.go:5	0x104e0a8		488d05c9850200		LEAQ main.statictmp_0(SB), AX		
  channel_close.go:5	0x104e0af		4889442408		MOVQ AX, 0x8(SP)			
  channel_close.go:5	0x104e0b4		e8675efbff		CALL runtime.chansend1(SB)		
  channel_close.go:6	0x104e0b9		488b442428		MOVQ 0x28(SP), AX			
  channel_close.go:6	0x104e0be		48890424		MOVQ AX, 0(SP)				
  channel_close.go:6	0x104e0c2		e8c966fbff		CALL runtime.closechan(SB)		
  channel_close.go:7	0x104e0c7		488b442428		MOVQ 0x28(SP), AX			
  channel_close.go:7	0x104e0cc		4889442430		MOVQ AX, 0x30(SP)			
  channel_close.go:7	0x104e0d1		eb00			JMP 0x104e0d3				
  channel_close.go:7	0x104e0d3		488b442430		MOVQ 0x30(SP), AX			
  channel_close.go:7	0x104e0d8		48890424		MOVQ AX, 0(SP)				
  channel_close.go:7	0x104e0dc		488d442420		LEAQ 0x20(SP), AX			
  channel_close.go:7	0x104e0e1		4889442408		MOVQ AX, 0x8(SP)			
  channel_close.go:7	0x104e0e6		e87569fbff		CALL runtime.chanrecv2(SB)		
  channel_close.go:7	0x104e0eb		0fb6442410		MOVZX 0x10(SP), AX			
  channel_close.go:7	0x104e0f0		8844241f		MOVB AL, 0x1f(SP)			
  channel_close.go:7	0x104e0f4		84c0			TESTL AL, AL				
  channel_close.go:7	0x104e0f6		7502			JNE 0x104e0fa				
  channel_close.go:7	0x104e0f8		eb0d			JMP 0x104e107				
  channel_close.go:7	0x104e0fa		48c744242000000000	MOVQ $0x0, 0x20(SP)			
  channel_close.go:7	0x104e103		eb00			JMP 0x104e105				
  channel_close.go:7	0x104e105		ebcc			JMP 0x104e0d3				
  channel_close.go:11	0x104e107		eb00			JMP 0x104e109				
  channel_close.go:11	0x104e109		eb00			JMP 0x104e10b				
  channel_close.go:11	0x104e10b		eb00			JMP 0x104e10d				
  channel_close.go:11	0x104e10d		488b6c2438		MOVQ 0x38(SP), BP			
  channel_close.go:11	0x104e112		4883c440		ADDQ $0x40, SP				
  channel_close.go:11	0x104e116		c3			RET					
  channel_close.go:3	0x104e117		e8e488ffff		CALL runtime.morestack_noctxt(SB)	
  channel_close.go:3	0x104e11c		e93fffffff		JMP main.main(SB)			
  :-1			0x104e121		cc			INT $0x3				
  :-1			0x104e122		cc			INT $0x3				
  :-1			0x104e123		cc			INT $0x3				
  :-1			0x104e124		cc			INT $0x3				
  :-1			0x104e125		cc			INT $0x3				
  :-1			0x104e126		cc			INT $0x3				
  :-1			0x104e127		cc			INT $0x3				
  :-1			0x104e128		cc			INT $0x3				
  :-1			0x104e129		cc			INT $0x3				
  :-1			0x104e12a		cc			INT $0x3				
  :-1			0x104e12b		cc			INT $0x3				
  :-1			0x104e12c		cc			INT $0x3				
  :-1			0x104e12d		cc			INT $0x3				
  :-1			0x104e12e		cc			INT $0x3				
  :-1			0x104e12f		cc			INT $0x3				
