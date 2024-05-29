package phf

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prime = 2305843009213693951

type HashAndDisplace struct {
	keys       []uint64
	n          uint64
	table      []uint64
	g          []uint64
	primaryH   func(uint64) uint64
	secondaryH func(uint64) uint64
}

func (hd *HashAndDisplace) PrintToFile(filename string) {
	if hd == nil {
		fmt.Println("Error: HashAndDisplace is nil")
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	if hd.keys == nil {
		fmt.Println("Error: keys are nil")
		return
	}

	if hd.primaryH == nil {
		fmt.Println("Error: primary hash function is nil")
		return
	}

	if hd.table == nil {
		fmt.Println("Error: hash table is nil")
		return
	}

	for _, key := range hd.keys {
		hash := hd.primaryH(key) % uint64(len(hd.table))
		value := hd.table[hash]
		fmt.Fprintf(file, "Key: %d, Hash: %d, Value: %d\n", key, hash, value)
	}
}

func NewHashAndDisplace(keys []uint64) *HashAndDisplace {
	hd := &HashAndDisplace{
		keys:       keys,
		n:          uint64(len(keys)),
		table:      make([]uint64, 2*uint64(len(keys))), //Makes the hash table twice as large
		g:          make([]uint64, 2*uint64(len(keys))),
		secondaryH: generateSecondaryHashFunction(uint64(time.Now().UnixNano())), //Initialize secondaryH
	}
	hd.construct()
	return hd
}

func generateHashFunction(seed uint64) func(uint64) uint64 {
	s := rand.NewSource(int64(seed))
	r := rand.New(s)
	var a, b uint64
	for a == 0 {
		a = r.Uint64()
	}
	for b == 0 {
		b = r.Uint64()
	}
	return func(x uint64) uint64 {
		return (a*x + b) % prime
	}
}

func generateSecondaryHashFunction(seed uint64) func(uint64) uint64 {
	s := rand.NewSource(int64(seed))
	r := rand.New(s)
	var a, b uint64
	for a == 0 {
		a = r.Uint64()
	}
	for b == 0 {
		b = r.Uint64()
	}
	return func(x uint64) uint64 {
		return (a*x + b) % prime
	}
}

func (hd *HashAndDisplace) construct() {
    hd.primaryH = generateHashFunction(uint64(time.Now().UnixNano())) // Unique seed for each iteration
    hd.table = make([]uint64, 2*hd.n)
    hd.g = make([]uint64, 2*hd.n)
    buckets := make([][]uint64, 2*hd.n)
    for _, key := range hd.keys {
        slot := hd.primaryH(key) % uint64(len(buckets)) //Ensures slot is within bounds
        buckets[slot] = append(buckets[slot], key)
    }
    for i, bucket := range buckets {
        if len(bucket) > 0 {
            hg := generateHashFunction(uint64(time.Now().UnixNano())) // Unique seed for each bucket
            for _, key := range bucket {
                slot := hg(key) % uint64(len(hd.table))
                offset := hd.secondaryH(key) // Secondary hash function
                for hd.table[slot] != 0 {
                    slot = (slot + offset) % uint64(len(hd.table)) // Double hashing
                }
                hd.table[slot] = key
            }
            hd.g[i] = hg(bucket[0]) // Store the secondary hash function
        }
    }
}

func (hd *HashAndDisplace) Get(key uint64) (uint64, bool) {
	index := hd.primaryH(key) % uint64(len(hd.table))
	if hd.table[index] == key {
		//fmt.Println("Key found at primary position")
		return hd.table[index], true
	}

	secondaryIndex := (index + hd.g[index]) % uint64(len(hd.table))
	if hd.table[secondaryIndex] == key {
		//fmt.Println("Key found at secondary position")
		return hd.table[secondaryIndex], true
	}
	return 0, false
}

func GenerateUniqueKeys(n int) []uint64 {
	keys := make([]uint64, 0, n)
	seen := make(map[uint64]bool)

	for len(keys) < n {
		key := rand.Uint64() // Generate a random key
		if !seen[key] {      // If the key hasn't been seen before
			seen[key] = true         // Mark the key as seen
			keys = append(keys, key) // Add the key to the keys slice
		}
	}

	return keys
}

/* old version

func (hd *HashAndDisplace) construct() {
	constructStart:
		hd.primaryH = generateHashFunction(uint64(time.Now().UnixNano())) // Unique seed for each iteration
		fmt.Println("Primary hash function generated.")
		hd.table = make([]uint64, 2*hd.n)
		hd.g = make([]uint64, 2*hd.n)
		buckets := make([][]uint64, 2*hd.n)
		for _, key := range hd.keys {
			slot := hd.primaryH(key) % uint64(len(buckets)) //Ensures slot is within bounds
			for slot == 0 {
				hd.primaryH = generateHashFunction(uint64(time.Now().UnixNano())) // Unique seed for each iteration
				slot = hd.primaryH(key) % uint64(len(buckets))
			}
			buckets[slot] = append(buckets[slot], key)
		}
		for i, bucket := range buckets {
			if len(bucket) > 1 {
				collision := true
				for collision {
					collision = false
					hg := generateHashFunction(uint64(time.Now().UnixNano())) // Unique seed for each bucket
					secondaryTable := make([]uint64, len(bucket)*len(bucket)) // Secondary hash table
					for _, key := range bucket {
						secondarySlot := hg(key) % uint64(len(secondaryTable))
						if secondaryTable[secondarySlot] != 0 {
							collision = true
							break
						} else {
							secondaryTable[secondarySlot] = key
						}
					}
					if !collision {
						// Check if the slot in the main hash table is unoccupied
						secondarySlot := hg(bucket[0]) % uint64(len(hd.table))
						if hd.table[secondarySlot] == 0 {
							hd.g[i-1] = hg(bucket[0]) // Store the secondary hash function
							for _, key := range bucket {
								hd.table[hg(key)%uint64(len(hd.table))] = key
							}
						} else {
							collision = true
						}
					}
				}
				fmt.Printf("Collisions resolved for bucket %d.\n", i)
			} else if len(bucket) == 1 {
				pos := hd.primaryH(bucket[0]) % (2 * hd.n)
				if hd.table[pos] != 0 {
					goto constructStart
				}
				hd.table[pos] = bucket[0]
				hd.g[pos] = 0 // No secondary hash function needed
			}
		}
	}

*/
