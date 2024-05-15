package phf

import (
	"fmt"
	"os"
	"sort"
)

type Data interface {
	Build(keys []uint64)
	Get(key uint64)
}

type MyPHF struct {
	data map[uint64]uint64
}

const NumberOfSlots = 25

func Create() MyPHF {
	return MyPHF{}
}
func customHashKey(key uint64, hashes map[uint64]uint64) uint64 {
	var hash uint64 = 5381
	for i := 0; i < 64; i += 8 {
		b := byte((key >> i) & 0xff)
		hash = ((hash << 5) + hash + uint64(b)) % NumberOfSlots
	}

	// Linear probing: if the slot is occupied, check the next one
	originalHash := hash
	for _, exists := hashes[hash]; exists; _, exists = hashes[hash] {
		hash = (hash + 1) % NumberOfSlots

		if hash == originalHash {
			break
		}
	}

	return hash
}

func (mp *MyPHF) Build(keys []uint64) error {
	mp.data = make(map[uint64]uint64)
	collisions := make(map[uint64][]uint64)

	for _, key := range keys {
		hash := customHashKey(key, mp.data)

		if hash >= NumberOfSlots {
			return fmt.Errorf("hash value %d is out of bounds", hash)
		}

		if _, exists := mp.data[hash]; exists {
			collisions[hash] = append(collisions[hash], key)
		} else {
			mp.data[hash] = key
		}
	}

	hashKeys := make([]int, 0, len(mp.data))
	for k := range mp.data {
		hashKeys = append(hashKeys, int(k))
	}

	sort.Ints(hashKeys)

	for hash, keys := range collisions {
		fmt.Printf("Hash %d is shared by keys %v\n", hash, keys)
	}
	file, err := os.Create("phf.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return nil
	}
	defer file.Close()

	// writes the keys and values to the file - for testing
	for _, k := range hashKeys {
		fmt.Fprintf(file, "Slot %d: Key %v\n", k, mp.data[uint64(k)])
	}

	for hash, keys := range collisions {
		fmt.Fprintf(file, "Hash %d is shared by keys %v\n", hash, keys)
	}
	return nil
}

func (mp MyPHF) Get(key uint64) (uint64, error) {
	if mp.data == nil {
		return 0, fmt.Errorf("data map is not initialized")
	}

	hash := customHashKey(key, mp.data)

	if _, exists := mp.data[hash]; !exists {
		return 0, fmt.Errorf("key %d does not exist", key)
	}
	return hash, nil
}
