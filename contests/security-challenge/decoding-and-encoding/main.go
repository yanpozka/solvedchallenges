package main

import (
	"fmt"
	// "strings"
)

const (
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	enc32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
)

var decMap [256]byte
var N int

func main() {

	for i := 0; i < len(enc32); i++ {
		decMap[enc32[i]] = byte(i)
	}

	fmt.Scanf("%d", &N)

	for ; N > 0; N-- {
		var line string
		fmt.Scanf("%s", &line)

		s := []byte(line)
		size_r32 := len(line) / 8 * 5
		r32 := make([]byte, size_r32)
		decbase32(r32, s)

		var decoded32 string = string(r32)
		fmt.Println(decoded32)
		/*var lenmulti4 int = (len(decoded32)) / 3 * 4 // size_r32 + (size_r32 % 4) //

		result := make([]byte, lenmulti4)
		encodeBase64(result, []byte(decoded32)) // !!

		fmt.Println(strings.Replace(string(result), "...=", "", -1))*/
	}
}

func decbase32(buffer, source []byte) {
	var end bool = false

	for len(source) > 0 && !end {
		var dbuf [8]byte
		var dlen int = 8
		for j := 0; j < 8; {
			if len(source) == 0 {
				return
			}
			in := source[0]
			source = source[1:]
			if in == '=' && j >= 2 && len(source) < 8 {
				if len(source)+j < 8-1 {
					return
				}
				for k := 0; k < 8-1-j; k++ {
					if len(source) > k && source[k] != '=' {
						return
					}
				}
				dlen, end = j, true
				if dlen == 1 || dlen == 3 || dlen == 6 {
					return
				}
				break
			}
			dbuf[j] = decMap[in]
			if dbuf[j] == 0xFF {
				return
			}
			j++ // !!
		}
		switch dlen {
		case 8:
			buffer[4] = dbuf[6]<<5 | dbuf[7]
			fallthrough
		case 7:
			buffer[3] = dbuf[4]<<7 | dbuf[5]<<2 | dbuf[6]>>3
			fallthrough
		case 5:
			buffer[2] = dbuf[3]<<4 | dbuf[4]>>1
			fallthrough
		case 4:
			buffer[1] = dbuf[1]<<6 | dbuf[2]<<1 | dbuf[3]>>4
			fallthrough
		case 2:
			buffer[0] = dbuf[0]<<3 | dbuf[1]>>2
		}
		buffer = buffer[5:]
	}
}

func encodeBase64(buffer, source []byte) {
	for len(source) > 0 {
		var a_byte, b_byte, c_byte, d_byte byte

		var lensource = len(source)

		switch lensource {
		default:
			d_byte = source[2] & 0x3F
			c_byte = source[2] >> 6
			fallthrough
		case 2:
			c_byte |= (source[1] << 2) & 0x3F
			b_byte = source[1] >> 4
			fallthrough
		case 1:
			b_byte |= (source[0] << 4) & 0x3F
			a_byte = source[0] >> 2
		}
		// fmt.Println(a_byte, b_byte, c_byte, d_byte)
		len_buffer := len(buffer)

		if len_buffer > 0 {
			buffer[0] = chars[a_byte]
			if len_buffer > 1 {
				buffer[1] = chars[b_byte]
				if len_buffer > 2 {
					buffer[2] = chars[c_byte]
				}
				if len_buffer > 3 {
					buffer[3] = chars[d_byte]
				}
			}
		}

		if lensource < 3 {
			if lensource < 2 && len_buffer >= 3 {
				buffer[2] = '='
			}
			if len_buffer >= 4 {
				buffer[3] = '='
			}
			break
		}
		source, buffer = source[3:], buffer[4:]
	}
}
