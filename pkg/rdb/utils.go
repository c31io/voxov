package rdb

import "encoding/binary"

func Int64ToByteSlice(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func ByteSliceToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}
