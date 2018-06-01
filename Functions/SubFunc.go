package Functions

func FSM() uint32 {
	var F = (sState[15] + R1) ^ R2
	var r = R2+(R3^sState[5])
	R3 = S2(R2)
	R2 = S1(R1)
	R1 = r
	return F
}

func S1(w uint32) uint32  {
	var W = splitS(w)
	var r0 = xMUL(Sr[W[0]], 0x1b) ^ Sr[W[1]] ^ Sr[W[2]] ^ xMUL(Sr[W[3]], 0x1b) ^ Sr[W[3]]
	var r1 = xMUL(Sr[W[0]], 0x1b) ^ Sr[W[0]] ^ xMUL(Sr[W[1]], 0x1b) ^ Sr[W[2]] ^ Sr[W[3]]
	var r2 = Sr[W[0]] ^ xMUL(Sr[W[1]], 0x1b) ^ Sr[W[1]]^xMUL(Sr[W[2]], 0x1b) ^ Sr[W[3]]
	var r3 = Sr[W[0]] ^ Sr[W[1]] ^ xMUL(Sr[W[2]], 0x1b) ^ xMUL(Sr[W[3]], 0x1b)^ Sr[W[2]]
	return uint32((r0 << 24) | (r1 << 16) | (r2 << 8) | r3)
}
func S2(w uint32) uint32  {
	var W = splitS(w)
	var r0 = xMUL(Sq[W[0]], 0x1b) ^ Sq[W[1]] ^ Sq[W[2]] ^ xMUL(Sq[W[3]], 0x1b) ^ Sq[W[3]]
	var r1 = xMUL(Sq[W[0]], 0x1b) ^ Sq[W[0]] ^ xMUL(Sq[W[1]], 0x1b) ^ Sq[W[2]] ^ Sq[W[3]]
	var r2 = Sq[W[0]] ^ xMUL(Sq[W[1]], 0x1b) ^ Sq[W[1]]^xMUL(Sq[W[2]], 0x1b) ^ Sq[W[3]]
	var r3 = Sq[W[0]] ^ Sq[W[1]] ^ xMUL(Sq[W[2]], 0x1b) ^ xMUL(Sq[W[3]], 0x1b)^ Sq[W[2]]
	return uint32((r0 << 24) | (r1 << 16) | (r2 << 8) | r3)
}
func splitS(w uint32) [4]uint8{
	var result[4] uint8
	result[3] = uint8(w & 0xff)
	result[2] = uint8((w>>8) & 0xff)
	result[1] = uint8((w>>16) & 0xff)
	result[0] = uint8((w>>24) & 0xff)
	return result
}
func xMUL(V uint8, c uint8) uint8{
	var leftmost = (V & 0xFF) >> 7
	if leftmost == 1 {
		return ((V << 1)&0xFF) ^ c
	}else{
		return (V << 1)&0xFF
	}
}

func mul(c uint8) uint8{
return ((MULxPOW(c, 23, 0xA9) << 24) | (MULxPOW(c, 245, 0xA9) << 16) |
	(MULxPOW(c, 48, 0xA9) << 8) | MULxPOW(c, 239, 0xA9) )
}

func div(c uint8) uint8{
return ((MULxPOW(c, 16, 0xA9) << 24) | (MULxPOW(c, 39, 0xA9) << 16) |
	(MULxPOW(c, 6, 0xA9) << 8) | MULxPOW(c, 64, 0xA9))
}

func MULxPOW(V uint8, i uint8, c uint8) uint8 {
	var result uint8
	if(i == 0){
		result = V
	}else{
		xMUL(MULxPOW(V,i-1,c),c)
	}
	return result
}
