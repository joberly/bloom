package bloom

// BitArray is an array of bits represented by a slice of bytes.
type BitArray []byte

// NewBitArray creates a new array of bits of size n bits.
func NewBitArray(n uint64) BitArray {
	if n <= 0 {
		panic("bit array size must be greater than zero")
	}
	return make([]byte, (n+7)/8)
}

// Get returns the value of bit i, where i is from 0 to
// the size of the BitArray minus 1.
func (a BitArray) Get(i uint64) uint {
	return uint((a[i/8] >> (i % 8)) & 1)
}

// Set sets bit i to 1, where i is from 0 to the size of
// the BitArray minus 1.
func (a BitArray) Set(i uint64) {
	temp := a[i/8]
	temp |= (1 << (i % 8))
	a[i/8] = temp
}
