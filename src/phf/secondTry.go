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
	//secondaryH func(uint64) uint64
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
		keys:  keys,
		n:     uint64(len(keys)),
		table: make([]uint64, 2*uint64(len(keys))), // make the hash table twice as large
		g:     make([]uint64, 2*uint64(len(keys))),
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

func (hd *HashAndDisplace) construct() {
	collision := true
	for collision {
		fmt.Println("Generating new primary hash function...")
		h := generateHashFunction(uint64(time.Now().UnixNano())) // Unique seed for each iteration
		hd.primaryH = h
		collision = false
		hd.table = make([]uint64, 2*hd.n)
		hd.g = make([]uint64, 2*hd.n)
		buckets := make([][]uint64, 2*hd.n)
		for _, key := range hd.keys {
			slot := h(key) % uint64(len(buckets)) // Ensure slot is within bounds
			buckets[slot] = append(buckets[slot], key)
		}
		for i, bucket := range buckets {
			if len(bucket) > 1 {
				fmt.Printf("Collision detected at slot %d, generating secondary hash function...\n", i)
				hg := generateHashFunction(uint64(time.Now().UnixNano())) // Unique seed for each bucket
				for _, key := range bucket {
					secondarySlot := (h(key) + hg(key)) % (2 * hd.n)
					if hd.table[secondarySlot] != 0 {
						collision = true
						fmt.Printf("Collision detected with secondary hash function at slot %d, regenerating primary hash function...\n", i)
						break
					} else {
						hd.table[secondarySlot] = key
						hd.g[secondarySlot] = hg(key)
					}
				}
				if collision {
					break
				}
			} else if len(bucket) == 1 {
				pos := h(bucket[0]) % (2 * hd.n)
				if hd.table[pos] != 0 {
					collision = true
					fmt.Printf("Unexpected collision at slot %d, regenerating primary hash function...\n", i)
					break
				}
				hd.table[pos] = bucket[0]
				hd.g[pos] = h(bucket[0])
			}
		}
	}
	fmt.Println("Hash functions generated successfully, no collisions.")
}

// checks if the key exists in the hash table
func (hd *HashAndDisplace) Query(key uint64) bool {
	if hd == nil || hd.primaryH == nil {
		fmt.Println("Error: HashAndDisplace or its hash functions are nil")
		return false
	}
	tableSize := uint64(len(hd.table))
	index := hd.primaryH(key) % tableSize
	if hd.table[index] == key {
		return true
	}
	displacement := hd.g[index]
	pos := (index + displacement) % tableSize
	return hd.table[pos] == key
}
