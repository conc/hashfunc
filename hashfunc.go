package hashfunc

import ()

func SDBMHash(str string) uint32 {
	var hash uint32 = 0
	for i := 0; i < len(str); i++ {
		hash = uint32(str[i]) + (hash << 6) + (hash << 16) - hash
	}

	return (hash & 0x7FFFFFFF)
}

func RSHash(str string) uint32 {
	var (
		b    uint32 = 378551
		a    uint32 = 63689
		hash uint32 = 0
	)

	for i := 0; i < len(str); i++ {
		hash = hash*a + uint32(str[i])
		a *= b
	}

	return (hash & 0x7FFFFFFF)
}

func JSHash(str string) uint32 {
	var hash uint32 = 1315423911

	for i := 0; i < len(str); i++ {
		hash ^= ((hash << 5) + uint32(str[i]) + (hash >> 2))
	}

	return (hash & 0x7FFFFFFF)
}

func PJWHash(str string) uint32 {
	var (
		BitsInUnignedInt uint32 = 32 //(sizeof(uint32) * 8);
		ThreeQuarters    uint32 = (uint32)((BitsInUnignedInt * 3) / 4)
		OneEighth        uint32 = (uint32)(BitsInUnignedInt / 8)
		HighBits         uint32 = (uint32)(0xFFFFFFFF) << (BitsInUnignedInt - OneEighth)
		hash             uint32 = 0
	)

	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint32(str[i])
		if test := hash & HighBits; test != 0 {
			hash = ((hash ^ (test >> ThreeQuarters)) & (^HighBits))
		}
	}

	return (hash & 0x7FFFFFFF)
}

func ELFHash(str string) uint32 {
	var hash uint32 = 0

	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint32(str[i])
		if x := (hash & 0xF0000000); x != 0 {
			hash ^= (x >> 24)
			hash &= ^x
		}
	}

	return (hash & 0x7FFFFFFF)
}

func BKDRHash(str string) uint32 {
	var (
		seed uint32 = 131 // 31 131 1313 13131 131313 etc..
		hash uint32 = 0
	)

	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint32(str[i])
	}

	return (hash & 0x7FFFFFFF)
}

func DJBHash(str string) uint32 {
	var hash uint32 = 5381

	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + uint32(str[i])
	}

	return (hash & 0x7FFFFFFF)
}

func APHash(str string) uint32 {
	var hash uint32 = 0

	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= ((hash << 7) ^ uint32(str[i]) ^ (hash >> 3))
		} else {
			hash ^= (^((hash << 11) ^ uint32(str[i]) ^ (hash >> 5)))
		}
	}

	return (hash & 0x7FFFFFFF)
}
