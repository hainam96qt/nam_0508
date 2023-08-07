package convert_type

import (
	"database/sql"
	"encoding/binary"
	"reflect"
	"time"
	"unsafe"
)

func NewNullString(arg string) sql.NullString {
	return sql.NullString{
		String: arg,
		Valid:  true,
	}
}

func NewNullInt32(arg int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: arg,
		Valid: true,
	}
}

func NewNullInt64(arg int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: arg,
		Valid: true,
	}
}
func NewNullFloat64(arg float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: float64(arg),
		Valid:   true,
	}
}

func NewNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func StringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Int64ToBytes(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func BytesToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}
