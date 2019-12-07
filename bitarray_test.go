package bloom

import (
	"testing"
)

func TestBitArraySetOne(t *testing.T) {
	for size := uint64(1); size < 64; size++ {
		for i := uint64(0); i < size; i++ {
			bits := NewBitArray(size)
			bits.Set(i)
			for j := uint64(0); j < size; j++ {
				if i == j {
					if bits.Get(j) != 1 {
						t.Errorf("size %d bit %d not set", size, j)
					}
				} else {
					if bits.Get(j) != 0 {
						t.Errorf("size %d bit %d set", size, j)
					}
				}
			}
		}
	}
}

func TestBitArraySetMany(t *testing.T) {
	for size := uint64(1); size < 64; size++ {
		bits := NewBitArray(size)
		for i := uint64(0); i < size; i++ {
			bits.Set(i)
			for j := uint64(0); j < size; j++ {
				if j <= i {
					if bits.Get(j) != 1 {
						t.Errorf("size %d bit %d not set", size, j)
					}
				} else {
					if bits.Get(j) != 0 {
						t.Errorf("size %d bit %d set", size, j)
					}
				}
			}
		}
	}
}
