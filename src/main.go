package main

import (
    "phf/phf"
    "fmt"
    "math/rand"
)
func main() {
	keys := make([]uint64, phf.NumberOfSlots)
    for i := 0; i < phf.NumberOfSlots; i++ {
        keys[i] = uint64(i)
    }
    mp := phf.Create()
    err := mp.Build(keys)
    if err != nil {
        fmt.Println("Error:", err)
    }

    for i := 0; i < 5; i++ {
        key := uint64(rand.Intn(phf.NumberOfSlots))
        slot, err := mp.Get(key)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("The slot for key %d is %d\n", key, slot)
        }
    }
}