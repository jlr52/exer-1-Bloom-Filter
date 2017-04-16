package bloom_filter

import (
	"hash/fnv"
)

type BloomFilter struct {
	filterMap []bool
	numHashes uint32
}

const (
    NEGATIVE = 1
    INCONCLUSIVE  = 2
)

type IsMemberReponse struct {
	negative bool
	inconclusive bool
}


// Constructor
func NewBloomFilter(config *BloomFilterConfig) *BloomFilter {
	filterMap := make([]bool, config.mapSize, config.mapSize)
	return &BloomFilter{filterMap, config.numHashes}
}

// Public Functions

func (bloom *BloomFilter) IsMember(elem string) int{
	for i := uint32(1); i <= bloom.numHashes; i++ {
	    if indicatorIndex := bloom.getNthHash(elem, i); !bloom.filterMap[indicatorIndex] {
            return NEGATIVE;
	    }
	}
	return INCONCLUSIVE
}

func (bloom *BloomFilter) Insert(elem string) {
	for i := uint32(1); i <= bloom.numHashes; i++ {
        indicatorIndex := bloom.getNthHash(elem, i)
        bloom.filterMap[indicatorIndex] = true
	}
}

// Private Functions

func (bloom *BloomFilter) getNthHash(s string, n uint32) uint32 {
    return (getFNV1Hash(s) + n * getFNV1aHash(s)) % uint32(cap(bloom.filterMap))
}

func getFNV1Hash(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

func getFNV1aHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}



