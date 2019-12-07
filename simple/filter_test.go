package simple

import (
	"math/rand"
	"testing"
)

const _seed int64 = 1903429042

func TestSingle(t *testing.T) {
	rnd := rand.New(rand.NewSource(_seed))
	for buckets := uint64(8); buckets < 64; buckets *= 2 {
		for hashes := uint64(2); hashes < 8; hashes++ {
			for size := 0; size < 64; size++ {
				for i := 0; i < 128; i++ {
					b := make([]byte, size)
					rnd.Read(b)
					filter := New(buckets, hashes)
					filter.Add(b)
					if !filter.Contains(b) {
						t.Errorf("Contains returned false for buckets %d hashes %d size %d index %d",
							buckets, hashes, size, i)
					}
				}
			}
		}
	}
}

func TestMultiple(t *testing.T) {
	rnd := rand.New(rand.NewSource(_seed))
	for buckets := uint64(8); buckets < 64; buckets *= 2 {
		for hashes := uint64(2); hashes < 8; hashes++ {
			for size := 0; size < 64; size++ {
				filter := New(buckets, hashes)
				bs := make([][]byte, 128)
				for i := range bs {
					b := make([]byte, size)
					rnd.Read(b)
					bs[i] = b
					filter.Add(bs[i])
				}
				for i, b := range bs {
					if !filter.Contains(b) {
						t.Errorf("Contains returned false for buckets %d hashes %d size %d index %d",
							buckets, hashes, size, i)
					}
				}
			}
		}
	}
}

func TestEmpty(t *testing.T) {
	rnd := rand.New(rand.NewSource(_seed))
	for buckets := uint64(8); buckets < 64; buckets *= 2 {
		for hashes := uint64(2); hashes < 8; hashes++ {
			for size := 0; size < 64; size++ {
				for i := 0; i < 128; i++ {
					b := make([]byte, size)
					rnd.Read(b)
					filter := New(buckets, hashes)
					if filter.Contains(b) {
						t.Errorf("Contains returned true for empty filter on buckets %d hashes %d size %d index %d",
							buckets, hashes, size, i)
					}
				}
			}
		}
	}
}
