package simple

import (
	"encoding/binary"
	"hash/fnv"

	"github.com/joberly/bloom"
)

// Filter is a simple Bloom filter with customizable bucket size.
// It uses the FNV-1a hash with the a*ib % m trick to create i hashes
// from a single FNV-1a computation.
type Filter struct {
	hashes  uint64
	buckets uint64
	bits    bloom.BitArray
}

// New creates a new Filter with the given number of buckets and hashes.
func New(buckets, hashes uint64) *Filter {
	if buckets < 1 {
		panic("must have at least 1 bucket")
	}
	if hashes < 1 {
		panic("must have at least 1 hash")
	}

	f := Filter{
		hashes:  hashes,
		buckets: buckets,
		bits:    bloom.NewBitArray(buckets),
	}
	return &f
}

// calcHashes calculates the hash values using FNV-1a.
// Each hash value returned is modulo the number of buckets.
func (f *Filter) calcHashes(b []byte) []uint64 {
	// Use the single hash trick.
	hash := fnv.New128a().Sum(b)
	ha := binary.BigEndian.Uint64(hash[0:8])
	hb := binary.BigEndian.Uint64(hash[8:])
	hashes := make([]uint64, f.hashes)
	for i := uint64(0); i < uint64(f.hashes); i++ {
		hashes[i] = (ha + (i * hb)) % uint64(f.buckets)
	}
	return hashes
}

// Add adds value b to the filter.
func (f *Filter) Add(b []byte) {
	hashes := f.calcHashes(b)
	// Set bit in bit array for each hash.
	for _, hash := range hashes {
		f.bits.Set(hash)
	}
}

// Contains checks to see if the filter contains value b.
func (f *Filter) Contains(b []byte) bool {
	hashes := f.calcHashes(b)
	for _, hash := range hashes {
		if f.bits.Get(hash) == 0 {
			return false
		}
	}
	return true
}
