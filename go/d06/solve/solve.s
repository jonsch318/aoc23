
#include "textflag.h"

DATA half_ptr<>+0(SB)/4, $(0.5)
DATA half_ptr<>+4(SB)/4, $(0.5)
DATA half_ptr<>+8(SB)/4, $(0.5)
DATA half_ptr<>+12(SB)/4, $(0.5)
GLOBL half_ptr<>(SB), RODATA|NOPTR, $16

DATA one_ptr<>+0(SB)/4, $(1.0)
DATA one_ptr<>+4(SB)/4, $(1.0)
DATA one_ptr<>+8(SB)/4, $(1.0)
DATA one_ptr<>+12(SB)/4, $(1.0)
GLOBL one_ptr<>(SB), RODATA|NOPTR, $16

DATA neg_one_ptr<>+0(SB)/4, $(-1.0)
DATA neg_one_ptr<>+4(SB)/4, $(-1.0)
DATA neg_one_ptr<>+8(SB)/4, $(-1.0)
DATA neg_one_ptr<>+12(SB)/4, $(-1.0)
GLOBL neg_one_ptr<>(SB), RODATA|NOPTR, $16

// func SolveAVX(times [4]float32, distances [4]float32) [4]float32
// memory layout
// +0 times[0] | times[1]
// +8 times[2] | times[3]
// +16 distances[0] | distances[1]
// +24 distances[2] | distances[3]
// +32 ret[0] | ret[1]
// +40 re[2] | ret[3]
TEXT ·SolveAVX(SB), NOSPLIT, $0-48
    //load to vector registers vector move precision single...
    MOVAPS times+0(FP), X0
    MOVAPS distances+16(FP), X1
    
    // timesHalf & -dist
    MULPS half_ptr<>+0(SB), X0
    MULPS neg_one_ptr<>+0(SB), X1

    //fused multiply add => dist = x0*x0 + -dist 
    VFMADD231PS X0, X0, X1

    //sqrt
    SQRTPS X1, X2
    
    //ceil(sub) timeHalf - x2
    VSUBPS X2, X0, X3
    ROUNDPS $2, X3, X4

    //floor(add)
    VADDPS X2, X0, X5
    ROUNDPS $1, X5, X6

    // res = x2-x1
    SUBPS X4, X6
    // res + 1
    ADDPS one_ptr<>+0(SB), X6 
    MOVAPS X6, ret+32(FP)
    RET

TEXT ·SolveAVX2(SB), NOSPLIT, $0-48
    RET
