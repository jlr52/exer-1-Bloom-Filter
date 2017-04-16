package bloom_filter

type BloomFilterConfig struct {
	mapSize uint32
	numHashes uint32
}

func NewBloomFilterConfig(mapSize uint32, numHashes uint32) *BloomFilterConfig {
	return &BloomFilterConfig{mapSize, numHashes}
}