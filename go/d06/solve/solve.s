
#include "textflag.h"

DATA half_ptr<>+0(SB)/8, $(0.5)
DATA half_ptr<>+8(SB)/8, $(0.5)
DATA half_ptr<>+16(SB)/8, $(0.5)
DATA half_ptr<>+24(SB)/8, $(0.5)
GLOBL half_ptr<>(SB), RODATA, $32

DATA one_ptr<>+0(SB)/8, $(1.0)
DATA one_ptr<>+8(SB)/8, $(1.0)
DATA one_ptr<>+16(SB)/8, $(1.0)
DATA one_ptr<>+24(SB)/8, $(1.0)
GLOBL one_ptr<>(SB), RODATA, $32

DATA neg_one_ptr<>+0(SB)/8, $(-1.0)
DATA neg_one_ptr<>+8(SB)/8, $(-1.0)
DATA neg_one_ptr<>+16(SB)/8, $(-1.0)
DATA neg_one_ptr<>+24(SB)/8, $(-1.0)
GLOBL neg_one_ptr<>(SB), RODATA, $32


DATA half_ptr_32<>+0(SB)/4, $(0.5)
DATA half_ptr_32<>+4(SB)/4, $(0.5)
DATA half_ptr_32<>+8(SB)/4, $(0.5)
DATA half_ptr_32<>+12(SB)/4, $(0.5)
GLOBL half_ptr_32<>(SB), RODATA, $16

DATA one_ptr_32<>+0(SB)/4, $(1.0)
DATA one_ptr_32<>+4(SB)/4, $(1.0)
DATA one_ptr_32<>+8(SB)/4, $(1.0)
DATA one_ptr_32<>+12(SB)/4, $(1.0)
GLOBL one_ptr_32<>(SB), RODATA, $16

DATA neg_one_ptr_32<>+0(SB)/4, $(-1.0)
DATA neg_one_ptr_32<>+4(SB)/4, $(-1.0)
DATA neg_one_ptr_32<>+8(SB)/4, $(-1.0)
DATA neg_one_ptr_32<>+12(SB)/4, $(-1.0)
GLOBL neg_one_ptr_32<>(SB), RODATA, $16


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
    MOVUPS times+0(FP), X0
    MOVUPS distances+16(FP), X1
    
    // timesHalf & -dist
    MULPS half_ptr_32<>+0(SB), X0
    MULPS neg_one_ptr_32<>+0(SB), X1

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
    ADDPS one_ptr_32<>+0(SB), X6 
    MOVUPS X6, ret+32(FP)

    RET

// func SolveAVX2(times [4]float64, distances [4]float64) [4]float64
// memory layout
// +0 times[0]
// +8 times[1] 
// +16 times[2] 
// +24 times[3] 
// +32 distances[0]
// +40 distances[1]
// +48 distances[2]
// +56 distances[3]
// +64 ret[0]
// +72 ret[1]
// +80 ret[2]
// +88 ret[3]
TEXT ·SolveAVX2(SB), NOSPLIT, $0-96
    //load to vector registers vector move precision single...
    VMOVUPS times+0(FP), Y0
    VMOVUPS distances+32(FP), Y1
    
    // timesHalf & -dist
    VMULPD half_ptr<>+0(SB), Y0, Y0
    VMULPD neg_one_ptr<>+0(SB), Y1, Y1 

     //fused multiply add => dist = x0*x0 + -dist 
    VFMADD231PD Y0, Y0, Y1

    //sqrt
    VSQRTPD Y1, Y2
    
    //ceil(sub) timeHalf - x2
    VSUBPD Y2, Y0, Y3
    VROUNDPD $2, Y3, Y4

    //floor(add)
    VADDPD Y2, Y0, Y5
    VROUNDPD $1, Y5, Y6

    // res = x2-x1
    VSUBPD Y4, Y6, Y6
    // res + 1
    VADDPD one_ptr<>+0(SB), Y6, Y6 
    VMOVUPD Y6, ret+64(FP)
    RET
