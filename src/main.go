package main

import (
	"fmt"
	"math/rand"
	"phf/phf"
	"time"
)

func main() {
	//Generates random keys
	keys := phf.GenerateUniqueKeys(1000000)

	// Measure the time it takes to execute NewHashAndDisplace
	start1 := time.Now()
	hd := phf.NewHashAndDisplace(keys)
	elapsed1 := time.Since(start1)
	fmt.Printf("NewHashAndDisplace took %s\n", elapsed1)

	hd.PrintToFile("phf.txt")

	// Measure the time it takes to get 10 keys
	start2 := time.Now()
	for i := 0; i < 1000; i++ {
		key := keys[rand.Intn(len(keys))] // Pick a random key
		hd.Get(key)
	}
	elapsed2 := time.Since(start2)
	fmt.Printf("Getting 10 keys took %s\n", elapsed2)

	// Measure time for 100 gets
	start3 := time.Now()
	for i := 0; i < 10000; i++ {
		key := keys[rand.Intn(len(keys))] // Pick a random key
		hd.Get(key)
	}
	elapsed3 := time.Since(start3)
	fmt.Printf("Getting 100 keys took %s\n", elapsed3)

	// Measure time for 1000 gets
	start4 := time.Now()
	for i := 0; i < 100000; i++ {
		key := keys[rand.Intn(len(keys))] // Pick a random key
		hd.Get(key)
	}
	elapsed4 := time.Since(start4)
	fmt.Printf("Getting 1000 keys took %s\n", elapsed4)
}
