package bloomgo

import (
	"math"
	"sync"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	NumBits       uint64
	HashCount     uint64
	FalsePositive float64
	NumItems      uint64
	bits          []bool
	mutex         sync.RWMutex
}

func New(numItems uint64, falsePos float64) *BloomFilter {
	numBits := bfGetNumBits(numItems, falsePos)
	hashCount := bfGetHashCount(numBits, numItems)

	return &BloomFilter{
		NumBits:       numBits,
		HashCount:     hashCount,
		FalsePositive: falsePos,
		NumItems:      numItems,
		bits:          make([]bool, numBits),
	}
}

func bfGetNumBits(numItems uint64, falsePos float64) uint64 {
	numBits := -1 * (float64(numItems) * math.Log(falsePos)) / (math.Pow(math.Log(2), 2))
	return uint64(numBits)
}

func bfGetHashCount(numBits uint64, numItems uint64) uint64 {
	numHash := float64(numBits) / float64(numItems) * math.Log(2)
	return uint64(numHash)
}

func (bf *BloomFilter) getIndexes(item []byte) []uint64 {
	var indexes []uint64
	for i := 0; i < int(bf.HashCount); i++ {
		hash := murmur3.Sum64WithSeed(item, uint32(i) /*seed*/)
		hash = hash % bf.NumBits
		indexes = append(indexes, hash)
	}
	return indexes
}

func (bf *BloomFilter) Add(item []byte) {
	indexes := bf.getIndexes(item)
	bf.mutex.Lock()
	defer bf.mutex.Unlock()

	for _, index := range indexes {
		bf.bits[index] = true
	}
}

func (bf *BloomFilter) Exists(item []byte) bool {
	indexes := bf.getIndexes(item)
	bf.mutex.RLock()
	defer bf.mutex.RUnlock()

	for _, index := range indexes {
		if !bf.bits[index] {
			return false
		}
	}
	return true
}
