#include "textflag.h"

// func rdrandU64() uint64
TEXT ·rdrandU64(SB), NOSPLIT, $0
REDO:
    BYTE $0x48; BYTE $0x0F; BYTE $0xC7; BYTE $0xF0;
    JNC REDO
OK:
    MOVQ AX, ret+0(FP)
    RET

// func rdseedU64() uint64
TEXT ·rdseedU64(SB), NOSPLIT, $0
REDO:
    BYTE $0x48; BYTE $0x0F; BYTE $0xC7; BYTE $0xF8;
    JNC REDO
OK:
    MOVQ AX, ret+0(FP)
    RET
