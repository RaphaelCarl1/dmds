package phf

import (
    "fmt"
    "sort"
	"os"
)

type Data interface {
	Build()
	Get()
}

type myPHF struct {
	data map[uint64][10]byte
}

func Create() myPHF {
	return myPHF{}
}
func customHashKey(key uint64, hashes map[uint64]uint64) uint64 {
	var hash uint64 = 5381
    for i := 0; i < 64; i += 8 {
        b := byte((key >> i) & 0xff) // Extract a byte from the key
        hash = ((hash << 5) + hash + uint64(b)) % 10 // Limit the hash value to 10 slots
    }
	
	// Linear probing: if the slot is occupied, check the next one
	originalHash := hash
	for _, exists := hashes[hash]; exists; _, exists = hashes[hash] {
		hash = (hash + 1) % 10 // Wrap around to the start if we reach the end
	
		// If we've checked all the slots and they're all full, break the loop
		if hash == originalHash {
			break
		}
	}
	
	return hash
}

func (PHF myPHF) Build(keys []uint64) {
	hashes := make(map[uint64]uint64)
    collisions := make(map[uint64][]uint64)

    for _, key := range keys {
        hash := customHashKey(key, hashes)

        if _, exists := hashes[hash]; exists {
            collisions[hash] = append(collisions[hash], key)
        } else {
            hashes[hash] = key
        }
    }

    // Create a slice of the keys
    hashKeys := make([]int, 0, len(hashes))
    for k := range hashes {
        hashKeys = append(hashKeys, int(k))
    }

    // Sort the keys
    sort.Ints(hashKeys)
    // Print the keys and values in order
    for _, k := range hashKeys {
        fmt.Printf("Slot %d: Key %v\n", k, hashes[uint64(k)])
    }

    for hash, keys := range collisions {
        fmt.Printf("Hash %d is shared by keys %v\n", hash, keys)
    }
	file, err := os.Create("phf.txt")
    if err != nil {
        fmt.Println("Unable to create file:", err)
        return
    }
    defer file.Close()

    // Write the keys and values in order to the file
    for _, k := range hashKeys {
        fmt.Fprintf(file, "Slot %d: Key %v\n", k, hashes[uint64(k)])
    }

    for hash, keys := range collisions {
        fmt.Fprintf(file, "Hash %d is shared by keys %v\n", hash, keys)
    }
}

func (PHF myPHF) Get(key uint64) uint64 {
	//returns a key
	return key
}

/*I am aware that this looks very minimalistic, as the implementation I have in mind does
not require a lot of functions.*/
