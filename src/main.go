package main

import (
	"fmt"
	"math/rand"
	"phf/phf"
)

func main() {
	// Generate 100 random keys
	keys := make([]uint64, 50)
	for i := range keys {
		keys[i] = rand.Uint64()
	}

	// Create a new HashAndDisplace
	hd := phf.NewHashAndDisplace(keys)
	hd.PrintToFile("phf.txt")

	// Query 5 random keys
	for i := 0; i < 100; i++ {
		key := keys[rand.Intn(len(keys))] // Pick a random key
		if hd.Query(key) {                // Check if the key exists in the hash table
			fmt.Printf("Key %d found\n", key)
		} else {
			fmt.Printf("Key %d not found\n", key)
		}
	}
}

/*
OLD VERSION
-----------------
func main() {
    for _, numKeys := range []int{1000, 100000, 1000000} {
        keys := make([]uint64, numKeys)
        for i := 0; i < numKeys; i++ {
            keys[i] = uint64(i)
        }
        bt := &phf.BinaryTree{}

        // Measure the time taken to insert keys
        startInsert := time.Now()
        for _, key := range keys {
            bt.Insert(key)
        }
        elapsedInsert := time.Since(startInsert)
        fmt.Printf("Time taken to insert %d keys: %s\n", numKeys, elapsedInsert)

        // Measure the time taken to search for keys
        startSearch := time.Now()
        for _, key := range keys {
            bt.Search(key)
        }
        elapsedSearch := time.Since(startSearch)
        fmt.Printf("Time taken to search for %d keys: %s\n", numKeys, elapsedSearch)

        // Measure the time taken for InOrder traversal
        startInOrder := time.Now()
        bt.InOrder()
        elapsedInOrder := time.Since(startInOrder)
        fmt.Printf("Time taken for InOrder traversal with %d keys: %s\n", numKeys, elapsedInOrder)

        // Existing code
        mp := phf.Create(numKeys*2)

        startBuild := time.Now()
        err := mp.Build(keys)
        elapsed := time.Since(startBuild)

        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("Time taken to build with %d keys: %s\n", numKeys, elapsed)
        }

        for _, numGets := range []int{1000, 100000, 1000000} {
            start := time.Now()

            for i := 0; i < numGets; i++ {
                index := rand.Intn(numKeys)
                key := keys[index]
                _, err := mp.Get(key)
                if err != nil {
                    fmt.Println("Error:", err)
                }
            }
            elapsed := time.Since(start)
            fmt.Printf("Time taken by %d Get() operations with %d keys: %s\n", numGets, numKeys, elapsed)
        }
    }
}
*/
