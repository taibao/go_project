package main

import (
	"fmt"
	"unsafe"
)

const (
	c1_32 int = 0xcc9e2d51
	c2_32 int = 0x1b873593
)

//获取订阅表号
func GetHash(data []byte) int {
	var h1 int = 0
	var h1b int = 0
	nblocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
		p = uintptr(unsafe.Pointer(&data[0]))
	}
	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
		k1 := *(*int)(unsafe.Pointer(p))
		k1 = ((k1 & 0xffff) * c1_32 + (((k1 >> 16) * 0xcc9e2d51) & 0xffff) << 16) & 0xffffffff
		k1 = (k1 << 15) | (k1 >> 17)
		k1 = ((k1 & 0xffff) * c2_32 + ((((k1 >> 16) * 0x1b873593) & 0xffff) << 16)) & 0xffffffff
		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19)
		h1b = (((h1 & 0xffff) * 5) + ((((h1 >> 16) * 5) & 0xffff) << 16)) & 0xffffffff
		h1  = ((h1b & 0xffff) + 0x6b64) + ((((h1b >> 16) + 0xe654) & 0xffff) << 16)
	}
	tail := data[nblocks*4:]

	var k1 int
	switch len(tail) & 3 {
	case 3:
		k1 ^= int(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= int(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= int(tail[0])
		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17)
		k1 *= c2_32
		h1 ^= k1
	}

	h1 ^= len(data)

	if h1 > 0 {
		h1 ^= h1 >> 16
	}else{
		h1 ^=  (h1 & 0x7fffffff) >> 16 | 0x8000
	}
	h1 = ((h1 & 0xffff) * 0x85ebca6b + ((((h1 >> 16) * 0x85ebca6b) & 0xffff) << 16)) & 0xffffffff

	h1 ^= h1 >> 13
	h1  = (((h1 & 0xffff) * 0xc2b2ae35) + ((((h1 >> 16) * 0xc2b2ae35) & 0xffff) << 16)) & 0xffffffff
	h1 ^= h1 >> 16
	return h1 % 100
}


func main(){
	//var t *testing.T
	h := GetHash([]byte("oTHW5v0M2zh9NiFVSXXGdWICecKQ"))
	fmt.Println(h%100)
	//if h != 4008393376 {
	//	t.Errorf("Hash %d is not equal to %d", h, 1815237614)
	//}
}

