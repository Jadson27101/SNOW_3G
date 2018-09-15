package Functions

var (
	sState [16] uint32
	R1     uint32
	R2     uint32
	R3     uint32
)

func Init(key [4] uint32, IV [4] uint32) {
	sState[15] = key[3] ^ IV[0]
	sState[14] = key[2]
	sState[13] = key[1]
	sState[12] = key[0] ^ IV[1]
	sState[11] = key[3] ^ 1
	sState[10] = key[2] ^ 1 ^ IV[2]
	sState[9] = key[1] ^ 1 ^ IV[3]
	sState[8] = key[0] ^ 1
	sState[7] = key[3]
	sState[6] = key[2]
	sState[5] = key[1]
	sState[4] = key[0]
	sState[3] = key[3] ^ 1
	sState[2] = key[2] ^ 1
	sState[1] = key[1] ^ 1
	sState[0] = key[0] ^ 1

	R1 = 0
	R2 = 0
	R3 = 0
	for i := 0; i < 32; i++ {
		InitNext()
	}
}
func InitNext() {
	s00 := (sState[0] >> (24)) & 0xFF
	s01 := (sState[0] >> (16)) & 0xFF
	s02 := (sState[0] >> 8) & 0xFF
	s03 := (sState[0] & 0xFF)
	s110 := (sState[11] >> (24)) & 0xFF
	s111 := (sState[11] >> (16)) & 0xFF
	s112 := (sState[11] >> 8) & 0xFF
	s113 := (sState[11] & 0xFF)
	v := (((s01 << 24) | (s02 << 16) | (s03 << 8) | 0x00) ^
		uint32(mul(uint8(s00))) ^ sState[2] ^ ((0x00 << 24) | (s110 << 16) | (s111 << 8) | s112) ^
		uint32(div(uint8(s113))) ^ FSM())

	for i := 0; i < 15; i++ {
		sState[i] = sState[i+1]
	}
	sState[15] = v
}
