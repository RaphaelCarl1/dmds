package main

import (
    "fmt"
    "os"
    "github.com/RaphaelCarl1/dmds/src/kvStore"
	//"github.com/RaphaelCarl1/dmds/src/kvStore/keyValueStore"
)

func main() {
    
    kvStore := keyValueStore.Create(120)
    var givenValue [10]byte
    givenValue[0] = 10
    kvStore.Put(10, givenValue)
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
