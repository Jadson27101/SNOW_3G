package main

import ("SNOW-3G/Functions")

func main(){
	key := [4]uint32{0x00000000, 0x00000000, 0x00000000, 0x80000000}
	IV := [4]uint32{0x00000001, 0x00000002, 0x00000003, 0x00000004}
	Functions.Init(key,IV)
	for i:=0; i<(1024*1024*1024)/4; i++{
		Functions.Next()
	}
}
