// Code generated by Fiat Cryptography. DO NOT EDIT.
//
// Autogenerated: 'src/ExtractionOCaml/unsaturated_solinas' --lang Go --relax-primitive-carry-to-bitwidth 32,64 --cmovznz-by-mul --internal-static --package-case flatcase --public-function-case UpperCamelCase --private-function-case camelCase --public-type-case UpperCamelCase --private-type-case camelCase --no-prefix-fiat --doc-newline-in-typedef-bounds --doc-prepend-header 'Code generated by Fiat Cryptography. DO NOT EDIT.' --doc-text-before-function-name '' --doc-text-before-type-name '' --package-name curve25519 '' 32 '(auto)' '2^255 - 19' carry_mul carry_square carry add sub opp selectznz to_bytes from_bytes relax carry_scmul121666 carry_add carry_sub carry_opp
//
// curve description (via package name): curve25519
//
// machine_wordsize = 32 (from "32")
//
// requested operations: carry_mul, carry_square, carry, add, sub, opp, selectznz, to_bytes, from_bytes, relax, carry_scmul121666, carry_add, carry_sub, carry_opp
//
// n = 10 (from "(auto)")
//
// s-c = 2^255 - [(1, 19)] (from "2^255 - 19")
//
// tight_bounds_multiplier = 1 (from "")
//
//
//
// Computed values:
//
//   carry_chain = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1]
//
//   eval z = z[0] + (z[1] << 26) + (z[2] << 51) + (z[3] << 77) + (z[4] << 102) + (z[5] << 128) + (z[6] << 153) + (z[7] << 179) + (z[8] << 204) + (z[9] << 230)
//
//   bytes_eval z = z[0] + (z[1] << 8) + (z[2] << 16) + (z[3] << 24) + (z[4] << 32) + (z[5] << 40) + (z[6] << 48) + (z[7] << 56) + (z[8] << 64) + (z[9] << 72) + (z[10] << 80) + (z[11] << 88) + (z[12] << 96) + (z[13] << 104) + (z[14] << 112) + (z[15] << 120) + (z[16] << 128) + (z[17] << 136) + (z[18] << 144) + (z[19] << 152) + (z[20] << 160) + (z[21] << 168) + (z[22] << 176) + (z[23] << 184) + (z[24] << 192) + (z[25] << 200) + (z[26] << 208) + (z[27] << 216) + (z[28] << 224) + (z[29] << 232) + (z[30] << 240) + (z[31] << 248)
//
//   balance = [0x7ffffda, 0x3fffffe, 0x7fffffe, 0x3fffffe, 0x7fffffe, 0x3fffffe, 0x7fffffe, 0x3fffffe, 0x7fffffe, 0x3fffffe]
package curve25519

type uint1 uint64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/fiat-crypto/pull/1006#issuecomment-892625927
type int1 int64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/fiat-crypto/pull/1006#issuecomment-892625927

// LooseFieldElement is a field element with loose bounds.
//
// Bounds:
//
//   [[0x0 ~> 0xc000000], [0x0 ~> 0x6000000], [0x0 ~> 0xc000000], [0x0 ~> 0x6000000], [0x0 ~> 0xc000000], [0x0 ~> 0x6000000], [0x0 ~> 0xc000000], [0x0 ~> 0x6000000], [0x0 ~> 0xc000000], [0x0 ~> 0x6000000]]
type LooseFieldElement [10]uint32

// TightFieldElement is a field element with tight bounds.
//
// Bounds:
//
//   [[0x0 ~> 0x4000000], [0x0 ~> 0x2000000], [0x0 ~> 0x4000000], [0x0 ~> 0x2000000], [0x0 ~> 0x4000000], [0x0 ~> 0x2000000], [0x0 ~> 0x4000000], [0x0 ~> 0x2000000], [0x0 ~> 0x4000000], [0x0 ~> 0x2000000]]
type TightFieldElement [10]uint32

// addcarryxU26 is an addition with carry.
//
// Postconditions:
//   out1 = (arg1 + arg2 + arg3) mod 2^26
//   out2 = ⌊(arg1 + arg2 + arg3) / 2^26⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0x3ffffff]
//   arg3: [0x0 ~> 0x3ffffff]
// Output Bounds:
//   out1: [0x0 ~> 0x3ffffff]
//   out2: [0x0 ~> 0x1]
func addcarryxU26(out1 *uint32, out2 *uint1, arg1 uint1, arg2 uint32, arg3 uint32) {
	x1 := ((uint32(arg1) + arg2) + arg3)
	x2 := (x1 & 0x3ffffff)
	x3 := uint1((x1 >> 26))
	*out1 = x2
	*out2 = x3
}

// subborrowxU26 is a subtraction with borrow.
//
// Postconditions:
//   out1 = (-arg1 + arg2 + -arg3) mod 2^26
//   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^26⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0x3ffffff]
//   arg3: [0x0 ~> 0x3ffffff]
// Output Bounds:
//   out1: [0x0 ~> 0x3ffffff]
//   out2: [0x0 ~> 0x1]
func subborrowxU26(out1 *uint32, out2 *uint1, arg1 uint1, arg2 uint32, arg3 uint32) {
	x1 := ((int32(arg2) - int32(arg1)) - int32(arg3))
	x2 := int1((x1 >> 26))
	x3 := (uint32(x1) & 0x3ffffff)
	*out1 = x3
	*out2 = (0x0 - uint1(x2))
}

// addcarryxU25 is an addition with carry.
//
// Postconditions:
//   out1 = (arg1 + arg2 + arg3) mod 2^25
//   out2 = ⌊(arg1 + arg2 + arg3) / 2^25⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0x1ffffff]
//   arg3: [0x0 ~> 0x1ffffff]
// Output Bounds:
//   out1: [0x0 ~> 0x1ffffff]
//   out2: [0x0 ~> 0x1]
func addcarryxU25(out1 *uint32, out2 *uint1, arg1 uint1, arg2 uint32, arg3 uint32) {
	x1 := ((uint32(arg1) + arg2) + arg3)
	x2 := (x1 & 0x1ffffff)
	x3 := uint1((x1 >> 25))
	*out1 = x2
	*out2 = x3
}

// subborrowxU25 is a subtraction with borrow.
//
// Postconditions:
//   out1 = (-arg1 + arg2 + -arg3) mod 2^25
//   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^25⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0x1ffffff]
//   arg3: [0x0 ~> 0x1ffffff]
// Output Bounds:
//   out1: [0x0 ~> 0x1ffffff]
//   out2: [0x0 ~> 0x1]
func subborrowxU25(out1 *uint32, out2 *uint1, arg1 uint1, arg2 uint32, arg3 uint32) {
	x1 := ((int32(arg2) - int32(arg1)) - int32(arg3))
	x2 := int1((x1 >> 25))
	x3 := (uint32(x1) & 0x1ffffff)
	*out1 = x3
	*out2 = (0x0 - uint1(x2))
}

// cmovznzU32 is a single-word conditional move.
//
// Postconditions:
//   out1 = (if arg1 = 0 then arg2 else arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0xffffffff]
//   arg3: [0x0 ~> 0xffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0xffffffff]
func cmovznzU32(out1 *uint32, arg1 uint1, arg2 uint32, arg3 uint32) {
	x1 := (uint32(arg1) * 0xffffffff)
	x2 := ((x1 & arg3) | ((^x1) & arg2))
	*out1 = x2
}

// CarryMul multiplies two field elements and reduces the result.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 * eval arg2) mod m
//
func CarryMul(out1 *TightFieldElement, arg1 *LooseFieldElement, arg2 *LooseFieldElement) {
	x1 := (uint64(arg1[9]) * uint64((arg2[9] * 0x26)))
	x2 := (uint64(arg1[9]) * uint64((arg2[8] * 0x13)))
	x3 := (uint64(arg1[9]) * uint64((arg2[7] * 0x26)))
	x4 := (uint64(arg1[9]) * uint64((arg2[6] * 0x13)))
	x5 := (uint64(arg1[9]) * uint64((arg2[5] * 0x26)))
	x6 := (uint64(arg1[9]) * uint64((arg2[4] * 0x13)))
	x7 := (uint64(arg1[9]) * uint64((arg2[3] * 0x26)))
	x8 := (uint64(arg1[9]) * uint64((arg2[2] * 0x13)))
	x9 := (uint64(arg1[9]) * uint64((arg2[1] * 0x26)))
	x10 := (uint64(arg1[8]) * uint64((arg2[9] * 0x13)))
	x11 := (uint64(arg1[8]) * uint64((arg2[8] * 0x13)))
	x12 := (uint64(arg1[8]) * uint64((arg2[7] * 0x13)))
	x13 := (uint64(arg1[8]) * uint64((arg2[6] * 0x13)))
	x14 := (uint64(arg1[8]) * uint64((arg2[5] * 0x13)))
	x15 := (uint64(arg1[8]) * uint64((arg2[4] * 0x13)))
	x16 := (uint64(arg1[8]) * uint64((arg2[3] * 0x13)))
	x17 := (uint64(arg1[8]) * uint64((arg2[2] * 0x13)))
	x18 := (uint64(arg1[7]) * uint64((arg2[9] * 0x26)))
	x19 := (uint64(arg1[7]) * uint64((arg2[8] * 0x13)))
	x20 := (uint64(arg1[7]) * uint64((arg2[7] * 0x26)))
	x21 := (uint64(arg1[7]) * uint64((arg2[6] * 0x13)))
	x22 := (uint64(arg1[7]) * uint64((arg2[5] * 0x26)))
	x23 := (uint64(arg1[7]) * uint64((arg2[4] * 0x13)))
	x24 := (uint64(arg1[7]) * uint64((arg2[3] * 0x26)))
	x25 := (uint64(arg1[6]) * uint64((arg2[9] * 0x13)))
	x26 := (uint64(arg1[6]) * uint64((arg2[8] * 0x13)))
	x27 := (uint64(arg1[6]) * uint64((arg2[7] * 0x13)))
	x28 := (uint64(arg1[6]) * uint64((arg2[6] * 0x13)))
	x29 := (uint64(arg1[6]) * uint64((arg2[5] * 0x13)))
	x30 := (uint64(arg1[6]) * uint64((arg2[4] * 0x13)))
	x31 := (uint64(arg1[5]) * uint64((arg2[9] * 0x26)))
	x32 := (uint64(arg1[5]) * uint64((arg2[8] * 0x13)))
	x33 := (uint64(arg1[5]) * uint64((arg2[7] * 0x26)))
	x34 := (uint64(arg1[5]) * uint64((arg2[6] * 0x13)))
	x35 := (uint64(arg1[5]) * uint64((arg2[5] * 0x26)))
	x36 := (uint64(arg1[4]) * uint64((arg2[9] * 0x13)))
	x37 := (uint64(arg1[4]) * uint64((arg2[8] * 0x13)))
	x38 := (uint64(arg1[4]) * uint64((arg2[7] * 0x13)))
	x39 := (uint64(arg1[4]) * uint64((arg2[6] * 0x13)))
	x40 := (uint64(arg1[3]) * uint64((arg2[9] * 0x26)))
	x41 := (uint64(arg1[3]) * uint64((arg2[8] * 0x13)))
	x42 := (uint64(arg1[3]) * uint64((arg2[7] * 0x26)))
	x43 := (uint64(arg1[2]) * uint64((arg2[9] * 0x13)))
	x44 := (uint64(arg1[2]) * uint64((arg2[8] * 0x13)))
	x45 := (uint64(arg1[1]) * uint64((arg2[9] * 0x26)))
	x46 := (uint64(arg1[9]) * uint64(arg2[0]))
	x47 := (uint64(arg1[8]) * uint64(arg2[1]))
	x48 := (uint64(arg1[8]) * uint64(arg2[0]))
	x49 := (uint64(arg1[7]) * uint64(arg2[2]))
	x50 := (uint64(arg1[7]) * uint64((arg2[1] * 0x2)))
	x51 := (uint64(arg1[7]) * uint64(arg2[0]))
	x52 := (uint64(arg1[6]) * uint64(arg2[3]))
	x53 := (uint64(arg1[6]) * uint64(arg2[2]))
	x54 := (uint64(arg1[6]) * uint64(arg2[1]))
	x55 := (uint64(arg1[6]) * uint64(arg2[0]))
	x56 := (uint64(arg1[5]) * uint64(arg2[4]))
	x57 := (uint64(arg1[5]) * uint64((arg2[3] * 0x2)))
	x58 := (uint64(arg1[5]) * uint64(arg2[2]))
	x59 := (uint64(arg1[5]) * uint64((arg2[1] * 0x2)))
	x60 := (uint64(arg1[5]) * uint64(arg2[0]))
	x61 := (uint64(arg1[4]) * uint64(arg2[5]))
	x62 := (uint64(arg1[4]) * uint64(arg2[4]))
	x63 := (uint64(arg1[4]) * uint64(arg2[3]))
	x64 := (uint64(arg1[4]) * uint64(arg2[2]))
	x65 := (uint64(arg1[4]) * uint64(arg2[1]))
	x66 := (uint64(arg1[4]) * uint64(arg2[0]))
	x67 := (uint64(arg1[3]) * uint64(arg2[6]))
	x68 := (uint64(arg1[3]) * uint64((arg2[5] * 0x2)))
	x69 := (uint64(arg1[3]) * uint64(arg2[4]))
	x70 := (uint64(arg1[3]) * uint64((arg2[3] * 0x2)))
	x71 := (uint64(arg1[3]) * uint64(arg2[2]))
	x72 := (uint64(arg1[3]) * uint64((arg2[1] * 0x2)))
	x73 := (uint64(arg1[3]) * uint64(arg2[0]))
	x74 := (uint64(arg1[2]) * uint64(arg2[7]))
	x75 := (uint64(arg1[2]) * uint64(arg2[6]))
	x76 := (uint64(arg1[2]) * uint64(arg2[5]))
	x77 := (uint64(arg1[2]) * uint64(arg2[4]))
	x78 := (uint64(arg1[2]) * uint64(arg2[3]))
	x79 := (uint64(arg1[2]) * uint64(arg2[2]))
	x80 := (uint64(arg1[2]) * uint64(arg2[1]))
	x81 := (uint64(arg1[2]) * uint64(arg2[0]))
	x82 := (uint64(arg1[1]) * uint64(arg2[8]))
	x83 := (uint64(arg1[1]) * uint64((arg2[7] * 0x2)))
	x84 := (uint64(arg1[1]) * uint64(arg2[6]))
	x85 := (uint64(arg1[1]) * uint64((arg2[5] * 0x2)))
	x86 := (uint64(arg1[1]) * uint64(arg2[4]))
	x87 := (uint64(arg1[1]) * uint64((arg2[3] * 0x2)))
	x88 := (uint64(arg1[1]) * uint64(arg2[2]))
	x89 := (uint64(arg1[1]) * uint64((arg2[1] * 0x2)))
	x90 := (uint64(arg1[1]) * uint64(arg2[0]))
	x91 := (uint64(arg1[0]) * uint64(arg2[9]))
	x92 := (uint64(arg1[0]) * uint64(arg2[8]))
	x93 := (uint64(arg1[0]) * uint64(arg2[7]))
	x94 := (uint64(arg1[0]) * uint64(arg2[6]))
	x95 := (uint64(arg1[0]) * uint64(arg2[5]))
	x96 := (uint64(arg1[0]) * uint64(arg2[4]))
	x97 := (uint64(arg1[0]) * uint64(arg2[3]))
	x98 := (uint64(arg1[0]) * uint64(arg2[2]))
	x99 := (uint64(arg1[0]) * uint64(arg2[1]))
	x100 := (uint64(arg1[0]) * uint64(arg2[0]))
	x101 := (x100 + (x45 + (x44 + (x42 + (x39 + (x35 + (x30 + (x24 + (x17 + x9)))))))))
	x102 := (x101 >> 26)
	x103 := (uint32(x101) & 0x3ffffff)
	x104 := (x91 + (x82 + (x74 + (x67 + (x61 + (x56 + (x52 + (x49 + (x47 + x46)))))))))
	x105 := (x92 + (x83 + (x75 + (x68 + (x62 + (x57 + (x53 + (x50 + (x48 + x1)))))))))
	x106 := (x93 + (x84 + (x76 + (x69 + (x63 + (x58 + (x54 + (x51 + (x10 + x2)))))))))
	x107 := (x94 + (x85 + (x77 + (x70 + (x64 + (x59 + (x55 + (x18 + (x11 + x3)))))))))
	x108 := (x95 + (x86 + (x78 + (x71 + (x65 + (x60 + (x25 + (x19 + (x12 + x4)))))))))
	x109 := (x96 + (x87 + (x79 + (x72 + (x66 + (x31 + (x26 + (x20 + (x13 + x5)))))))))
	x110 := (x97 + (x88 + (x80 + (x73 + (x36 + (x32 + (x27 + (x21 + (x14 + x6)))))))))
	x111 := (x98 + (x89 + (x81 + (x40 + (x37 + (x33 + (x28 + (x22 + (x15 + x7)))))))))
	x112 := (x99 + (x90 + (x43 + (x41 + (x38 + (x34 + (x29 + (x23 + (x16 + x8)))))))))
	x113 := (x102 + x112)
	x114 := (x113 >> 25)
	x115 := (uint32(x113) & 0x1ffffff)
	x116 := (x114 + x111)
	x117 := (x116 >> 26)
	x118 := (uint32(x116) & 0x3ffffff)
	x119 := (x117 + x110)
	x120 := (x119 >> 25)
	x121 := (uint32(x119) & 0x1ffffff)
	x122 := (x120 + x109)
	x123 := (x122 >> 26)
	x124 := (uint32(x122) & 0x3ffffff)
	x125 := (x123 + x108)
	x126 := (x125 >> 25)
	x127 := (uint32(x125) & 0x1ffffff)
	x128 := (x126 + x107)
	x129 := (x128 >> 26)
	x130 := (uint32(x128) & 0x3ffffff)
	x131 := (x129 + x106)
	x132 := (x131 >> 25)
	x133 := (uint32(x131) & 0x1ffffff)
	x134 := (x132 + x105)
	x135 := (x134 >> 26)
	x136 := (uint32(x134) & 0x3ffffff)
	x137 := (x135 + x104)
	x138 := (x137 >> 25)
	x139 := (uint32(x137) & 0x1ffffff)
	x140 := (x138 * uint64(0x13))
	x141 := (uint64(x103) + x140)
	x142 := uint32((x141 >> 26))
	x143 := (uint32(x141) & 0x3ffffff)
	x144 := (x142 + x115)
	x145 := uint1((x144 >> 25))
	x146 := (x144 & 0x1ffffff)
	x147 := (uint32(x145) + x118)
	out1[0] = x143
	out1[1] = x146
	out1[2] = x147
	out1[3] = x121
	out1[4] = x124
	out1[5] = x127
	out1[6] = x130
	out1[7] = x133
	out1[8] = x136
	out1[9] = x139
}

// CarrySquare squares a field element and reduces the result.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 * eval arg1) mod m
//
func CarrySquare(out1 *TightFieldElement, arg1 *LooseFieldElement) {
	x1 := (arg1[9] * 0x13)
	x2 := (x1 * 0x2)
	x3 := (arg1[9] * 0x2)
	x4 := (arg1[8] * 0x13)
	x5 := (uint64(x4) * uint64(0x2))
	x6 := (arg1[8] * 0x2)
	x7 := (arg1[7] * 0x13)
	x8 := (x7 * 0x2)
	x9 := (arg1[7] * 0x2)
	x10 := (arg1[6] * 0x13)
	x11 := (uint64(x10) * uint64(0x2))
	x12 := (arg1[6] * 0x2)
	x13 := (arg1[5] * 0x13)
	x14 := (arg1[5] * 0x2)
	x15 := (arg1[4] * 0x2)
	x16 := (arg1[3] * 0x2)
	x17 := (arg1[2] * 0x2)
	x18 := (arg1[1] * 0x2)
	x19 := (uint64(arg1[9]) * uint64((x1 * 0x2)))
	x20 := (uint64(arg1[8]) * uint64(x2))
	x21 := (uint64(arg1[8]) * uint64(x4))
	x22 := (uint64(arg1[7]) * (uint64(x2) * uint64(0x2)))
	x23 := (uint64(arg1[7]) * x5)
	x24 := (uint64(arg1[7]) * uint64((x7 * 0x2)))
	x25 := (uint64(arg1[6]) * uint64(x2))
	x26 := (uint64(arg1[6]) * x5)
	x27 := (uint64(arg1[6]) * uint64(x8))
	x28 := (uint64(arg1[6]) * uint64(x10))
	x29 := (uint64(arg1[5]) * (uint64(x2) * uint64(0x2)))
	x30 := (uint64(arg1[5]) * x5)
	x31 := (uint64(arg1[5]) * (uint64(x8) * uint64(0x2)))
	x32 := (uint64(arg1[5]) * x11)
	x33 := (uint64(arg1[5]) * uint64((x13 * 0x2)))
	x34 := (uint64(arg1[4]) * uint64(x2))
	x35 := (uint64(arg1[4]) * x5)
	x36 := (uint64(arg1[4]) * uint64(x8))
	x37 := (uint64(arg1[4]) * x11)
	x38 := (uint64(arg1[4]) * uint64(x14))
	x39 := (uint64(arg1[4]) * uint64(arg1[4]))
	x40 := (uint64(arg1[3]) * (uint64(x2) * uint64(0x2)))
	x41 := (uint64(arg1[3]) * x5)
	x42 := (uint64(arg1[3]) * (uint64(x8) * uint64(0x2)))
	x43 := (uint64(arg1[3]) * uint64(x12))
	x44 := (uint64(arg1[3]) * uint64((x14 * 0x2)))
	x45 := (uint64(arg1[3]) * uint64(x15))
	x46 := (uint64(arg1[3]) * uint64((arg1[3] * 0x2)))
	x47 := (uint64(arg1[2]) * uint64(x2))
	x48 := (uint64(arg1[2]) * x5)
	x49 := (uint64(arg1[2]) * uint64(x9))
	x50 := (uint64(arg1[2]) * uint64(x12))
	x51 := (uint64(arg1[2]) * uint64(x14))
	x52 := (uint64(arg1[2]) * uint64(x15))
	x53 := (uint64(arg1[2]) * uint64(x16))
	x54 := (uint64(arg1[2]) * uint64(arg1[2]))
	x55 := (uint64(arg1[1]) * (uint64(x2) * uint64(0x2)))
	x56 := (uint64(arg1[1]) * uint64(x6))
	x57 := (uint64(arg1[1]) * uint64((x9 * 0x2)))
	x58 := (uint64(arg1[1]) * uint64(x12))
	x59 := (uint64(arg1[1]) * uint64((x14 * 0x2)))
	x60 := (uint64(arg1[1]) * uint64(x15))
	x61 := (uint64(arg1[1]) * uint64((x16 * 0x2)))
	x62 := (uint64(arg1[1]) * uint64(x17))
	x63 := (uint64(arg1[1]) * uint64((arg1[1] * 0x2)))
	x64 := (uint64(arg1[0]) * uint64(x3))
	x65 := (uint64(arg1[0]) * uint64(x6))
	x66 := (uint64(arg1[0]) * uint64(x9))
	x67 := (uint64(arg1[0]) * uint64(x12))
	x68 := (uint64(arg1[0]) * uint64(x14))
	x69 := (uint64(arg1[0]) * uint64(x15))
	x70 := (uint64(arg1[0]) * uint64(x16))
	x71 := (uint64(arg1[0]) * uint64(x17))
	x72 := (uint64(arg1[0]) * uint64(x18))
	x73 := (uint64(arg1[0]) * uint64(arg1[0]))
	x74 := (x73 + (x55 + (x48 + (x42 + (x37 + x33)))))
	x75 := (x74 >> 26)
	x76 := (uint32(x74) & 0x3ffffff)
	x77 := (x64 + (x56 + (x49 + (x43 + x38))))
	x78 := (x65 + (x57 + (x50 + (x44 + (x39 + x19)))))
	x79 := (x66 + (x58 + (x51 + (x45 + x20))))
	x80 := (x67 + (x59 + (x52 + (x46 + (x22 + x21)))))
	x81 := (x68 + (x60 + (x53 + (x25 + x23))))
	x82 := (x69 + (x61 + (x54 + (x29 + (x26 + x24)))))
	x83 := (x70 + (x62 + (x34 + (x30 + x27))))
	x84 := (x71 + (x63 + (x40 + (x35 + (x31 + x28)))))
	x85 := (x72 + (x47 + (x41 + (x36 + x32))))
	x86 := (x75 + x85)
	x87 := (x86 >> 25)
	x88 := (uint32(x86) & 0x1ffffff)
	x89 := (x87 + x84)
	x90 := (x89 >> 26)
	x91 := (uint32(x89) & 0x3ffffff)
	x92 := (x90 + x83)
	x93 := (x92 >> 25)
	x94 := (uint32(x92) & 0x1ffffff)
	x95 := (x93 + x82)
	x96 := (x95 >> 26)
	x97 := (uint32(x95) & 0x3ffffff)
	x98 := (x96 + x81)
	x99 := (x98 >> 25)
	x100 := (uint32(x98) & 0x1ffffff)
	x101 := (x99 + x80)
	x102 := (x101 >> 26)
	x103 := (uint32(x101) & 0x3ffffff)
	x104 := (x102 + x79)
	x105 := (x104 >> 25)
	x106 := (uint32(x104) & 0x1ffffff)
	x107 := (x105 + x78)
	x108 := (x107 >> 26)
	x109 := (uint32(x107) & 0x3ffffff)
	x110 := (x108 + x77)
	x111 := (x110 >> 25)
	x112 := (uint32(x110) & 0x1ffffff)
	x113 := (x111 * uint64(0x13))
	x114 := (uint64(x76) + x113)
	x115 := uint32((x114 >> 26))
	x116 := (uint32(x114) & 0x3ffffff)
	x117 := (x115 + x88)
	x118 := uint1((x117 >> 25))
	x119 := (x117 & 0x1ffffff)
	x120 := (uint32(x118) + x91)
	out1[0] = x116
	out1[1] = x119
	out1[2] = x120
	out1[3] = x94
	out1[4] = x97
	out1[5] = x100
	out1[6] = x103
	out1[7] = x106
	out1[8] = x109
	out1[9] = x112
}

// Carry reduces a field element.
//
// Postconditions:
//   eval out1 mod m = eval arg1 mod m
//
func Carry(out1 *TightFieldElement, arg1 *LooseFieldElement) {
	x1 := arg1[0]
	x2 := ((x1 >> 26) + arg1[1])
	x3 := ((x2 >> 25) + arg1[2])
	x4 := ((x3 >> 26) + arg1[3])
	x5 := ((x4 >> 25) + arg1[4])
	x6 := ((x5 >> 26) + arg1[5])
	x7 := ((x6 >> 25) + arg1[6])
	x8 := ((x7 >> 26) + arg1[7])
	x9 := ((x8 >> 25) + arg1[8])
	x10 := ((x9 >> 26) + arg1[9])
	x11 := ((x1 & 0x3ffffff) + ((x10 >> 25) * 0x13))
	x12 := (uint32(uint1((x11 >> 26))) + (x2 & 0x1ffffff))
	x13 := (x11 & 0x3ffffff)
	x14 := (x12 & 0x1ffffff)
	x15 := (uint32(uint1((x12 >> 25))) + (x3 & 0x3ffffff))
	x16 := (x4 & 0x1ffffff)
	x17 := (x5 & 0x3ffffff)
	x18 := (x6 & 0x1ffffff)
	x19 := (x7 & 0x3ffffff)
	x20 := (x8 & 0x1ffffff)
	x21 := (x9 & 0x3ffffff)
	x22 := (x10 & 0x1ffffff)
	out1[0] = x13
	out1[1] = x14
	out1[2] = x15
	out1[3] = x16
	out1[4] = x17
	out1[5] = x18
	out1[6] = x19
	out1[7] = x20
	out1[8] = x21
	out1[9] = x22
}

// Add adds two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 + eval arg2) mod m
//
func Add(out1 *LooseFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := (arg1[0] + arg2[0])
	x2 := (arg1[1] + arg2[1])
	x3 := (arg1[2] + arg2[2])
	x4 := (arg1[3] + arg2[3])
	x5 := (arg1[4] + arg2[4])
	x6 := (arg1[5] + arg2[5])
	x7 := (arg1[6] + arg2[6])
	x8 := (arg1[7] + arg2[7])
	x9 := (arg1[8] + arg2[8])
	x10 := (arg1[9] + arg2[9])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
	out1[5] = x6
	out1[6] = x7
	out1[7] = x8
	out1[8] = x9
	out1[9] = x10
}

// Sub subtracts two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 - eval arg2) mod m
//
func Sub(out1 *LooseFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := ((0x7ffffda + arg1[0]) - arg2[0])
	x2 := ((0x3fffffe + arg1[1]) - arg2[1])
	x3 := ((0x7fffffe + arg1[2]) - arg2[2])
	x4 := ((0x3fffffe + arg1[3]) - arg2[3])
	x5 := ((0x7fffffe + arg1[4]) - arg2[4])
	x6 := ((0x3fffffe + arg1[5]) - arg2[5])
	x7 := ((0x7fffffe + arg1[6]) - arg2[6])
	x8 := ((0x3fffffe + arg1[7]) - arg2[7])
	x9 := ((0x7fffffe + arg1[8]) - arg2[8])
	x10 := ((0x3fffffe + arg1[9]) - arg2[9])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
	out1[5] = x6
	out1[6] = x7
	out1[7] = x8
	out1[8] = x9
	out1[9] = x10
}

// Opp negates a field element.
//
// Postconditions:
//   eval out1 mod m = -eval arg1 mod m
//
func Opp(out1 *LooseFieldElement, arg1 *TightFieldElement) {
	x1 := (0x7ffffda - arg1[0])
	x2 := (0x3fffffe - arg1[1])
	x3 := (0x7fffffe - arg1[2])
	x4 := (0x3fffffe - arg1[3])
	x5 := (0x7fffffe - arg1[4])
	x6 := (0x3fffffe - arg1[5])
	x7 := (0x7fffffe - arg1[6])
	x8 := (0x3fffffe - arg1[7])
	x9 := (0x7fffffe - arg1[8])
	x10 := (0x3fffffe - arg1[9])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
	out1[5] = x6
	out1[6] = x7
	out1[7] = x8
	out1[8] = x9
	out1[9] = x10
}

// Selectznz is a multi-limb conditional select.
//
// Postconditions:
//   out1 = (if arg1 = 0 then arg2 else arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
//   arg3: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
func Selectznz(out1 *[10]uint32, arg1 uint1, arg2 *[10]uint32, arg3 *[10]uint32) {
	var x1 uint32
	cmovznzU32(&x1, arg1, arg2[0], arg3[0])
	var x2 uint32
	cmovznzU32(&x2, arg1, arg2[1], arg3[1])
	var x3 uint32
	cmovznzU32(&x3, arg1, arg2[2], arg3[2])
	var x4 uint32
	cmovznzU32(&x4, arg1, arg2[3], arg3[3])
	var x5 uint32
	cmovznzU32(&x5, arg1, arg2[4], arg3[4])
	var x6 uint32
	cmovznzU32(&x6, arg1, arg2[5], arg3[5])
	var x7 uint32
	cmovznzU32(&x7, arg1, arg2[6], arg3[6])
	var x8 uint32
	cmovznzU32(&x8, arg1, arg2[7], arg3[7])
	var x9 uint32
	cmovznzU32(&x9, arg1, arg2[8], arg3[8])
	var x10 uint32
	cmovznzU32(&x10, arg1, arg2[9], arg3[9])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
	out1[5] = x6
	out1[6] = x7
	out1[7] = x8
	out1[8] = x9
	out1[9] = x10
}

// ToBytes serializes a field element to bytes in little-endian order.
//
// Postconditions:
//   out1 = map (λ x, ⌊((eval arg1 mod m) mod 2^(8 * (x + 1))) / 2^(8 * x)⌋) [0..31]
//
// Output Bounds:
//   out1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x7f]]
func ToBytes(out1 *[32]uint8, arg1 *TightFieldElement) {
	var x1 uint32
	var x2 uint1
	subborrowxU26(&x1, &x2, 0x0, arg1[0], 0x3ffffed)
	var x3 uint32
	var x4 uint1
	subborrowxU25(&x3, &x4, x2, arg1[1], 0x1ffffff)
	var x5 uint32
	var x6 uint1
	subborrowxU26(&x5, &x6, x4, arg1[2], 0x3ffffff)
	var x7 uint32
	var x8 uint1
	subborrowxU25(&x7, &x8, x6, arg1[3], 0x1ffffff)
	var x9 uint32
	var x10 uint1
	subborrowxU26(&x9, &x10, x8, arg1[4], 0x3ffffff)
	var x11 uint32
	var x12 uint1
	subborrowxU25(&x11, &x12, x10, arg1[5], 0x1ffffff)
	var x13 uint32
	var x14 uint1
	subborrowxU26(&x13, &x14, x12, arg1[6], 0x3ffffff)
	var x15 uint32
	var x16 uint1
	subborrowxU25(&x15, &x16, x14, arg1[7], 0x1ffffff)
	var x17 uint32
	var x18 uint1
	subborrowxU26(&x17, &x18, x16, arg1[8], 0x3ffffff)
	var x19 uint32
	var x20 uint1
	subborrowxU25(&x19, &x20, x18, arg1[9], 0x1ffffff)
	var x21 uint32
	cmovznzU32(&x21, x20, uint32(0x0), 0xffffffff)
	var x22 uint32
	var x23 uint1
	addcarryxU26(&x22, &x23, 0x0, x1, (x21 & 0x3ffffed))
	var x24 uint32
	var x25 uint1
	addcarryxU25(&x24, &x25, x23, x3, (x21 & 0x1ffffff))
	var x26 uint32
	var x27 uint1
	addcarryxU26(&x26, &x27, x25, x5, (x21 & 0x3ffffff))
	var x28 uint32
	var x29 uint1
	addcarryxU25(&x28, &x29, x27, x7, (x21 & 0x1ffffff))
	var x30 uint32
	var x31 uint1
	addcarryxU26(&x30, &x31, x29, x9, (x21 & 0x3ffffff))
	var x32 uint32
	var x33 uint1
	addcarryxU25(&x32, &x33, x31, x11, (x21 & 0x1ffffff))
	var x34 uint32
	var x35 uint1
	addcarryxU26(&x34, &x35, x33, x13, (x21 & 0x3ffffff))
	var x36 uint32
	var x37 uint1
	addcarryxU25(&x36, &x37, x35, x15, (x21 & 0x1ffffff))
	var x38 uint32
	var x39 uint1
	addcarryxU26(&x38, &x39, x37, x17, (x21 & 0x3ffffff))
	var x40 uint32
	var x41 uint1
	addcarryxU25(&x40, &x41, x39, x19, (x21 & 0x1ffffff))
	x42 := (x40 << 6)
	x43 := (x38 << 4)
	x44 := (x36 << 3)
	x45 := (x34 * uint32(0x2))
	x46 := (x30 << 6)
	x47 := (x28 << 5)
	x48 := (x26 << 3)
	x49 := (x24 << 2)
	x50 := (uint8(x22) & 0xff)
	x51 := (x22 >> 8)
	x52 := (uint8(x51) & 0xff)
	x53 := (x51 >> 8)
	x54 := (uint8(x53) & 0xff)
	x55 := uint8((x53 >> 8))
	x56 := (x49 + uint32(x55))
	x57 := (uint8(x56) & 0xff)
	x58 := (x56 >> 8)
	x59 := (uint8(x58) & 0xff)
	x60 := (x58 >> 8)
	x61 := (uint8(x60) & 0xff)
	x62 := uint8((x60 >> 8))
	x63 := (x48 + uint32(x62))
	x64 := (uint8(x63) & 0xff)
	x65 := (x63 >> 8)
	x66 := (uint8(x65) & 0xff)
	x67 := (x65 >> 8)
	x68 := (uint8(x67) & 0xff)
	x69 := uint8((x67 >> 8))
	x70 := (x47 + uint32(x69))
	x71 := (uint8(x70) & 0xff)
	x72 := (x70 >> 8)
	x73 := (uint8(x72) & 0xff)
	x74 := (x72 >> 8)
	x75 := (uint8(x74) & 0xff)
	x76 := uint8((x74 >> 8))
	x77 := (x46 + uint32(x76))
	x78 := (uint8(x77) & 0xff)
	x79 := (x77 >> 8)
	x80 := (uint8(x79) & 0xff)
	x81 := (x79 >> 8)
	x82 := (uint8(x81) & 0xff)
	x83 := uint8((x81 >> 8))
	x84 := (uint8(x32) & 0xff)
	x85 := (x32 >> 8)
	x86 := (uint8(x85) & 0xff)
	x87 := (x85 >> 8)
	x88 := (uint8(x87) & 0xff)
	x89 := uint1((x87 >> 8))
	x90 := (x45 + uint32(x89))
	x91 := (uint8(x90) & 0xff)
	x92 := (x90 >> 8)
	x93 := (uint8(x92) & 0xff)
	x94 := (x92 >> 8)
	x95 := (uint8(x94) & 0xff)
	x96 := uint8((x94 >> 8))
	x97 := (x44 + uint32(x96))
	x98 := (uint8(x97) & 0xff)
	x99 := (x97 >> 8)
	x100 := (uint8(x99) & 0xff)
	x101 := (x99 >> 8)
	x102 := (uint8(x101) & 0xff)
	x103 := uint8((x101 >> 8))
	x104 := (x43 + uint32(x103))
	x105 := (uint8(x104) & 0xff)
	x106 := (x104 >> 8)
	x107 := (uint8(x106) & 0xff)
	x108 := (x106 >> 8)
	x109 := (uint8(x108) & 0xff)
	x110 := uint8((x108 >> 8))
	x111 := (x42 + uint32(x110))
	x112 := (uint8(x111) & 0xff)
	x113 := (x111 >> 8)
	x114 := (uint8(x113) & 0xff)
	x115 := (x113 >> 8)
	x116 := (uint8(x115) & 0xff)
	x117 := uint8((x115 >> 8))
	out1[0] = x50
	out1[1] = x52
	out1[2] = x54
	out1[3] = x57
	out1[4] = x59
	out1[5] = x61
	out1[6] = x64
	out1[7] = x66
	out1[8] = x68
	out1[9] = x71
	out1[10] = x73
	out1[11] = x75
	out1[12] = x78
	out1[13] = x80
	out1[14] = x82
	out1[15] = x83
	out1[16] = x84
	out1[17] = x86
	out1[18] = x88
	out1[19] = x91
	out1[20] = x93
	out1[21] = x95
	out1[22] = x98
	out1[23] = x100
	out1[24] = x102
	out1[25] = x105
	out1[26] = x107
	out1[27] = x109
	out1[28] = x112
	out1[29] = x114
	out1[30] = x116
	out1[31] = x117
}

// FromBytes deserializes a field element from bytes in little-endian order.
//
// Postconditions:
//   eval out1 mod m = bytes_eval arg1 mod m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x7f]]
func FromBytes(out1 *TightFieldElement, arg1 *[32]uint8) {
	x1 := (uint32(arg1[31]) << 18)
	x2 := (uint32(arg1[30]) << 10)
	x3 := (uint32(arg1[29]) << 2)
	x4 := (uint32(arg1[28]) << 20)
	x5 := (uint32(arg1[27]) << 12)
	x6 := (uint32(arg1[26]) << 4)
	x7 := (uint32(arg1[25]) << 21)
	x8 := (uint32(arg1[24]) << 13)
	x9 := (uint32(arg1[23]) << 5)
	x10 := (uint32(arg1[22]) << 23)
	x11 := (uint32(arg1[21]) << 15)
	x12 := (uint32(arg1[20]) << 7)
	x13 := (uint32(arg1[19]) << 24)
	x14 := (uint32(arg1[18]) << 16)
	x15 := (uint32(arg1[17]) << 8)
	x16 := arg1[16]
	x17 := (uint32(arg1[15]) << 18)
	x18 := (uint32(arg1[14]) << 10)
	x19 := (uint32(arg1[13]) << 2)
	x20 := (uint32(arg1[12]) << 19)
	x21 := (uint32(arg1[11]) << 11)
	x22 := (uint32(arg1[10]) << 3)
	x23 := (uint32(arg1[9]) << 21)
	x24 := (uint32(arg1[8]) << 13)
	x25 := (uint32(arg1[7]) << 5)
	x26 := (uint32(arg1[6]) << 22)
	x27 := (uint32(arg1[5]) << 14)
	x28 := (uint32(arg1[4]) << 6)
	x29 := (uint32(arg1[3]) << 24)
	x30 := (uint32(arg1[2]) << 16)
	x31 := (uint32(arg1[1]) << 8)
	x32 := arg1[0]
	x33 := (x31 + uint32(x32))
	x34 := (x30 + x33)
	x35 := (x29 + x34)
	x36 := (x35 & 0x3ffffff)
	x37 := uint8((x35 >> 26))
	x38 := (x28 + uint32(x37))
	x39 := (x27 + x38)
	x40 := (x26 + x39)
	x41 := (x40 & 0x1ffffff)
	x42 := uint8((x40 >> 25))
	x43 := (x25 + uint32(x42))
	x44 := (x24 + x43)
	x45 := (x23 + x44)
	x46 := (x45 & 0x3ffffff)
	x47 := uint8((x45 >> 26))
	x48 := (x22 + uint32(x47))
	x49 := (x21 + x48)
	x50 := (x20 + x49)
	x51 := (x50 & 0x1ffffff)
	x52 := uint8((x50 >> 25))
	x53 := (x19 + uint32(x52))
	x54 := (x18 + x53)
	x55 := (x17 + x54)
	x56 := (x15 + uint32(x16))
	x57 := (x14 + x56)
	x58 := (x13 + x57)
	x59 := (x58 & 0x1ffffff)
	x60 := uint8((x58 >> 25))
	x61 := (x12 + uint32(x60))
	x62 := (x11 + x61)
	x63 := (x10 + x62)
	x64 := (x63 & 0x3ffffff)
	x65 := uint8((x63 >> 26))
	x66 := (x9 + uint32(x65))
	x67 := (x8 + x66)
	x68 := (x7 + x67)
	x69 := (x68 & 0x1ffffff)
	x70 := uint8((x68 >> 25))
	x71 := (x6 + uint32(x70))
	x72 := (x5 + x71)
	x73 := (x4 + x72)
	x74 := (x73 & 0x3ffffff)
	x75 := uint8((x73 >> 26))
	x76 := (x3 + uint32(x75))
	x77 := (x2 + x76)
	x78 := (x1 + x77)
	out1[0] = x36
	out1[1] = x41
	out1[2] = x46
	out1[3] = x51
	out1[4] = x55
	out1[5] = x59
	out1[6] = x64
	out1[7] = x69
	out1[8] = x74
	out1[9] = x78
}

// Relax is the identity function converting from tight field elements to loose field elements.
//
// Postconditions:
//   out1 = arg1
//
func Relax(out1 *LooseFieldElement, arg1 *TightFieldElement) {
	x1 := arg1[0]
	x2 := arg1[1]
	x3 := arg1[2]
	x4 := arg1[3]
	x5 := arg1[4]
	x6 := arg1[5]
	x7 := arg1[6]
	x8 := arg1[7]
	x9 := arg1[8]
	x10 := arg1[9]
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
	out1[5] = x6
	out1[6] = x7
	out1[7] = x8
	out1[8] = x9
	out1[9] = x10
}

// CarryScmul121666 multiplies a field element by 121666 and reduces the result.
//
// Postconditions:
//   eval out1 mod m = (121666 * eval arg1) mod m
//
func CarryScmul121666(out1 *TightFieldElement, arg1 *LooseFieldElement) {
	x1 := (uint64(0x1db42) * uint64(arg1[9]))
	x2 := (uint64(0x1db42) * uint64(arg1[8]))
	x3 := (uint64(0x1db42) * uint64(arg1[7]))
	x4 := (uint64(0x1db42) * uint64(arg1[6]))
	x5 := (uint64(0x1db42) * uint64(arg1[5]))
	x6 := (uint64(0x1db42) * uint64(arg1[4]))
	x7 := (uint64(0x1db42) * uint64(arg1[3]))
	x8 := (uint64(0x1db42) * uint64(arg1[2]))
	x9 := (uint64(0x1db42) * uint64(arg1[1]))
	x10 := (uint64(0x1db42) * uint64(arg1[0]))
	x11 := uint32((x10 >> 26))
	x12 := (uint32(x10) & 0x3ffffff)
	x13 := (uint64(x11) + x9)
	x14 := uint32((x13 >> 25))
	x15 := (uint32(x13) & 0x1ffffff)
	x16 := (uint64(x14) + x8)
	x17 := uint32((x16 >> 26))
	x18 := (uint32(x16) & 0x3ffffff)
	x19 := (uint64(x17) + x7)
	x20 := uint32((x19 >> 25))
	x21 := (uint32(x19) & 0x1ffffff)
	x22 := (uint64(x20) + x6)
	x23 := uint32((x22 >> 26))
	x24 := (uint32(x22) & 0x3ffffff)
	x25 := (uint64(x23) + x5)
	x26 := uint32((x25 >> 25))
	x27 := (uint32(x25) & 0x1ffffff)
	x28 := (uint64(x26) + x4)
	x29 := uint32((x28 >> 26))
	x30 := (uint32(x28) & 0x3ffffff)
	x31 := (uint64(x29) + x3)
	x32 := uint32((x31 >> 25))
	x33 := (uint32(x31) & 0x1ffffff)
	x34 := (uint64(x32) + x2)
	x35 := uint32((x34 >> 26))
	x36 := (uint32(x34) & 0x3ffffff)
	x37 := (uint64(x35) + x1)
	x38 := uint32((x37 >> 25))
	x39 := (uint32(x37) & 0x1ffffff)
	x40 := (x38 * 0x13)
	x41 := (x12 + x40)
	x42 := uint1((x41 >> 26))
	x43 := (x41 & 0x3ffffff)
	x44 := (uint32(x42) + x15)
	x45 := uint1((x44 >> 25))
	x46 := (x44 & 0x1ffffff)
	x47 := (uint32(x45) + x18)
	out1[0] = x43
	out1[1] = x46
	out1[2] = x47
	out1[3] = x21
	out1[4] = x24
	out1[5] = x27
	out1[6] = x30
	out1[7] = x33
	out1[8] = x36
	out1[9] = x39
}

// CarryAdd adds two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 + eval arg2) mod m
//
func CarryAdd(out1 *TightFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := (arg1[0] + arg2[0])
	x2 := ((x1 >> 26) + (arg1[1] + arg2[1]))
	x3 := ((x2 >> 25) + (arg1[2] + arg2[2]))
	x4 := ((x3 >> 26) + (arg1[3] + arg2[3]))
	x5 := ((x4 >> 25) + (arg1[4] + arg2[4]))
	x6 := ((x5 >> 26) + (arg1[5] + arg2[5]))
	x7 := ((x6 >> 25) + (arg1[6] + arg2[6]))
	x8 := ((x7 >> 26) + (arg1[7] + arg2[7]))
	x9 := ((x8 >> 25) + (arg1[8] + arg2[8]))
	x10 := ((x9 >> 26) + (arg1[9] + arg2[9]))
	x11 := ((x1 & 0x3ffffff) + ((x10 >> 25) * 0x13))
	x12 := (uint32(uint1((x11 >> 26))) + (x2 & 0x1ffffff))
	x13 := (x11 & 0x3ffffff)
	x14 := (x12 & 0x1ffffff)
	x15 := (uint32(uint1((x12 >> 25))) + (x3 & 0x3ffffff))
	x16 := (x4 & 0x1ffffff)
	x17 := (x5 & 0x3ffffff)
	x18 := (x6 & 0x1ffffff)
	x19 := (x7 & 0x3ffffff)
	x20 := (x8 & 0x1ffffff)
	x21 := (x9 & 0x3ffffff)
	x22 := (x10 & 0x1ffffff)
	out1[0] = x13
	out1[1] = x14
	out1[2] = x15
	out1[3] = x16
	out1[4] = x17
	out1[5] = x18
	out1[6] = x19
	out1[7] = x20
	out1[8] = x21
	out1[9] = x22
}

// CarrySub subtracts two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 - eval arg2) mod m
//
func CarrySub(out1 *TightFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := ((0x7ffffda + arg1[0]) - arg2[0])
	x2 := ((x1 >> 26) + ((0x3fffffe + arg1[1]) - arg2[1]))
	x3 := ((x2 >> 25) + ((0x7fffffe + arg1[2]) - arg2[2]))
	x4 := ((x3 >> 26) + ((0x3fffffe + arg1[3]) - arg2[3]))
	x5 := ((x4 >> 25) + ((0x7fffffe + arg1[4]) - arg2[4]))
	x6 := ((x5 >> 26) + ((0x3fffffe + arg1[5]) - arg2[5]))
	x7 := ((x6 >> 25) + ((0x7fffffe + arg1[6]) - arg2[6]))
	x8 := ((x7 >> 26) + ((0x3fffffe + arg1[7]) - arg2[7]))
	x9 := ((x8 >> 25) + ((0x7fffffe + arg1[8]) - arg2[8]))
	x10 := ((x9 >> 26) + ((0x3fffffe + arg1[9]) - arg2[9]))
	x11 := ((x1 & 0x3ffffff) + ((x10 >> 25) * 0x13))
	x12 := (uint32(uint1((x11 >> 26))) + (x2 & 0x1ffffff))
	x13 := (x11 & 0x3ffffff)
	x14 := (x12 & 0x1ffffff)
	x15 := (uint32(uint1((x12 >> 25))) + (x3 & 0x3ffffff))
	x16 := (x4 & 0x1ffffff)
	x17 := (x5 & 0x3ffffff)
	x18 := (x6 & 0x1ffffff)
	x19 := (x7 & 0x3ffffff)
	x20 := (x8 & 0x1ffffff)
	x21 := (x9 & 0x3ffffff)
	x22 := (x10 & 0x1ffffff)
	out1[0] = x13
	out1[1] = x14
	out1[2] = x15
	out1[3] = x16
	out1[4] = x17
	out1[5] = x18
	out1[6] = x19
	out1[7] = x20
	out1[8] = x21
	out1[9] = x22
}

// CarryOpp negates a field element.
//
// Postconditions:
//   eval out1 mod m = -eval arg1 mod m
//
func CarryOpp(out1 *TightFieldElement, arg1 *TightFieldElement) {
	x1 := (0x7ffffda - arg1[0])
	x2 := (uint32(uint1((x1 >> 26))) + (0x3fffffe - arg1[1]))
	x3 := (uint32(uint1((x2 >> 25))) + (0x7fffffe - arg1[2]))
	x4 := (uint32(uint1((x3 >> 26))) + (0x3fffffe - arg1[3]))
	x5 := (uint32(uint1((x4 >> 25))) + (0x7fffffe - arg1[4]))
	x6 := (uint32(uint1((x5 >> 26))) + (0x3fffffe - arg1[5]))
	x7 := (uint32(uint1((x6 >> 25))) + (0x7fffffe - arg1[6]))
	x8 := (uint32(uint1((x7 >> 26))) + (0x3fffffe - arg1[7]))
	x9 := (uint32(uint1((x8 >> 25))) + (0x7fffffe - arg1[8]))
	x10 := (uint32(uint1((x9 >> 26))) + (0x3fffffe - arg1[9]))
	x11 := ((x1 & 0x3ffffff) + (uint32(uint1((x10 >> 25))) * 0x13))
	x12 := (uint32(uint1((x11 >> 26))) + (x2 & 0x1ffffff))
	x13 := (x11 & 0x3ffffff)
	x14 := (x12 & 0x1ffffff)
	x15 := (uint32(uint1((x12 >> 25))) + (x3 & 0x3ffffff))
	x16 := (x4 & 0x1ffffff)
	x17 := (x5 & 0x3ffffff)
	x18 := (x6 & 0x1ffffff)
	x19 := (x7 & 0x3ffffff)
	x20 := (x8 & 0x1ffffff)
	x21 := (x9 & 0x3ffffff)
	x22 := (x10 & 0x1ffffff)
	out1[0] = x13
	out1[1] = x14
	out1[2] = x15
	out1[3] = x16
	out1[4] = x17
	out1[5] = x18
	out1[6] = x19
	out1[7] = x20
	out1[8] = x21
	out1[9] = x22
}
