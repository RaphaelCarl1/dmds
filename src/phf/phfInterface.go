package phf

import (
    "fmt"
    "os"
)

type Data interface {
    Build(keys []uint64)
    Get(key uint64)
}

type MyPHF struct {
    data    map[uint64]uint64
    hashes  map[uint64]uint64
    numKeys int
}

func Create(numKeys int) MyPHF {
    return MyPHF{
        data:    make(map[uint64]uint64),
        hashes:  make(map[uint64]uint64),
        numKeys: numKeys * 2,
    }
}

func secondaryHashKey(key uint64, numKeys int) uint64 {
    // Simple secondary hash function, you can replace this with something more complex if needed
    return (key % uint64(numKeys-1)) + 1
}

func primaryHashKey(key uint64, numKeys int) uint64 {
    // Constants for the multiply-shift method
    const multiplier = 11400714819323198485
    const shift = 32

    return (key * multiplier) >> (64 - shift) % uint64(numKeys)
}

func (mp *MyPHF) Build(keys []uint64) error {
    mp.data = make(map[uint64]uint64)
    mp.hashes = make(map[uint64]uint64)

    for _, key := range keys {
        hash := primaryHashKey(key, mp.numKeys)

        if hash >= uint64(mp.numKeys) {
            return fmt.Errorf("hash value %d is out of bounds", hash)
        }

        // If the slot is already occupied, it means there's a collision
        if _, exists := mp.hashes[hash]; exists {
            // Use double hashing to find a new slot for the key
            step := secondaryHashKey(key, mp.numKeys)
            originalHash := hash
            for _, exists = mp.hashes[hash]; exists; _, exists = mp.hashes[hash] {
                hash = (hash + step) % uint64(mp.numKeys)

                if hash == originalHash {
                    return fmt.Errorf("all slots are full")
                }
            }
        }

        // Insert the key into the found slot
        mp.hashes[hash] = key
        mp.data[key] = key
    }

    file, err := os.Create("phf.txt")
    if err != nil {
        fmt.Println("Unable to create file:", err)
        return nil
    }
    defer file.Close()

    // writes the keys and values to the file - for testing
    for k, v := range mp.data {
        fmt.Fprintf(file, "Slot %d: Key %v\n", k, v)
    }

    return nil
}

func (mp *MyPHF) Get(key uint64) (uint64, error) {
    if mp.data == nil {
        return 0, fmt.Errorf("data map is not initialized")
    }

    if value, exists := mp.data[key]; exists {
        // Found the key
        return value, nil
    } else {
        // The key does not exist
        return 0, fmt.Errorf("key %d does not exist", key)
    }
}