package Functions

func Next(){
	s00 := (sState[0] >> (24)) & 0xFF
	s01 := (sState[0] >> (16)) & 0xFF
	s02 := (sState[0] >> 8) & 0xFF
	s03 :=  (sState[0] & 0xFF)
	s110 := (sState[11] >> (24)) & 0xFF
	s111 := (sState[11] >> (16)) & 0xFF
	s112 := (sState[11] >> 8) & 0xFF
	s113 :=  (sState[11] & 0xFF)
	v :=  (((s01 << 24) | (s02 << 16) | (s03 << 8) | 0x00) ^
		uint32(mul(uint8(s00))) ^ sState[2] ^ ((0x00 << 24) | (s110 << 16) | (s111 << 8) | s112) ^
		uint32(div(uint8(s113))));
	for i:=0; i<15; i++{
		sState[i] = sState[i+1]
	}
	sState[15] = v
}