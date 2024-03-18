package main

import (
    "fmt"
    "os"
	"dmds/kvStore/keyValueStore"
)

func main() {

    // var kvStore = Create()
    //kvStore.put(10, 10.0)
    // Open a file for writing. Create it if it doesn't exist, truncate it if it does.
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Write a string to the file
    _, err = file.WriteString("Hello, world!\n")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Data written to file successfully.")
}

func Create() {
	panic("unimplemented")
}
