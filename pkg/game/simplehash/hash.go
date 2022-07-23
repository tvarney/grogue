package simplehash

const (
	// Initial32 is the initial hash value to use when generating a hash.
	Initial32 Hash32 = 0x811c9dc5

	prime32 uint32 = 0x01000193
)

// Hash32 is an immutable hash value.
//
// This type implements the FNV-1a hash algorithm; to properly use it, all
// hash values must start as the Initial32 constant. Updating the hash value
// using the `Add...` functions returns a new Hash32 value. As this type is
// a wrapper around the uint32 type, this is very cheap. It also allows
// reusing an intermediate value.
type Hash32 uint32

// AddUint8 returns a new Hash32 value with the given value added to it.
func (h Hash32) AddUint8(i uint8) Hash32 {
	return Hash32((uint32(h) ^ uint32(i)) * prime32)
}

// AddUint8 returns a new Hash32 value with the given value added to it.
func (h Hash32) AddInt8(i int8) Hash32 {
	return h.AddUint8(uint8(i))
}

// AddUint16 returns a new Hash32 value with the given value added to it.
func (h Hash32) AddUint16(i uint16) Hash32 {
	return h.AddUint8(uint8(i & 0x00FF)).AddUint8(uint8((i & 0xFF00) >> 8))
}

// AddInt16 returns a new Hash32 value with the given value added to it.
func (h Hash32) AddInt16(i int16) Hash32 {
	return h.AddUint16(uint16(i))
}

// AddUint32 returns a new Hash32 value with the given value added to it.
func (h Hash32) AddUint32(i uint32) Hash32 {
	return h.
		AddUint8(uint8((i & 0x000000FF) >> 0)).
		AddUint8(uint8((i & 0x0000FF00) >> 8)).
		AddUint8(uint8((i & 0x00FF0000) >> 16)).
		AddUint8(uint8((i & 0xFF000000) >> 24))
}

// AddInt32 returns a new Hash32 value with the given value added to it.
func (h Hash32) AddInt32(i int32) Hash32 {
	return h.AddUint32(uint32(i))
}

// AddUint64 returns a new Hash32 value with the given value added to it.
func (h Hash32) AddUint64(i uint64) Hash32 {
	return h.
		AddUint32(uint32((i & 0x00000000FFFFFFFF) >> 0)).
		AddUint32(uint32((i & 0xFFFFFFFF00000000) >> 32))
}

// AddInt64 returns a new Hash32 value with the given value added to it.
func (h Hash32) AddInt64(i int64) Hash32 {
	return h.AddUint64(uint64(i))
}
