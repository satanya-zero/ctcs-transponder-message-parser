package main

import (
	_ "TransponderMsgParse/packets"
	_ "TransponderMsgParse/packets/ETCSPacks"
	_ "TransponderMsgParse/packets/ETCSPacks/CTCSPacks"
)

var str = "1001000000010010011111111100110001010101100011101100101100010001000111101000000001010001000100110010000000110011010000000110000010101000011010000000000100000101010000100110001100000101010000100000000100000101001101000000001100000101000101000110000100000101001001000000001100000111101111000110000100000111101111100000001100000111110011000110000100000111110011000000001100000111100111000110000100000111100111000000001100000111101100000110000100000111101100000000001100000101000010100000000100000101000011000110001100000101000011000000000100000111110000000110001100000111110000000000000100000111101101000110001100000111101101101001111100000001010001010101011000111010000000000000001011100010100100001010000011111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111100"

func main() {
	msg := NewBinMessage(str)
	msg.ParseBody()
}
