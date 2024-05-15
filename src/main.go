package main

/*
import (
    "fmt"
    "sort"
	"os"
)

/*

func customHashKey(key [10]byte, hashes map[uint64][10]byte) uint64 {
	var hash uint64 = 5381
	for _, b := range key {
		hash = ((hash << 5) + hash + uint64(b)) % 10 // Limit the hash value to 100 slots
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

func build(keys [][10]byte) {
    hashes := make(map[uint64][10]byte)
    collisions := make(map[uint64][][10]byte)

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
*/
import (
	phf "github.com/RaphaelCarl1/dmds/src/phf"
)

func main() {
	keys := make([][10]byte, 20) // Adjust this to the number of keys you want to test
	for i := 0; i < 20; i++ {
		keys[i] = [10]byte{byte(i & 0xff), byte((i >> 8) & 0xff), byte((i >> 16) & 0xff), 0, 0, 0, 0, 0, 0, 0}
	}
	phf.Build(keys) // Call the Build function from the phf package
}