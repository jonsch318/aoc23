
#include "textflag.h"

// func Solve(times []float32, distances []float32) float32
TEXT ·Solve(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), BX
    MOVQ y+8(FP), BP
    ADDQ BP,BX
    MOVQ BX, ret+16(FP)
    RET
    
