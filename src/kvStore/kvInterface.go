package keyValueStore

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

//import _ "github.com/go-sql-driver/mysql" - for db-stuff

type Data interface {
	Put(uint64, [10]byte)
	Get(uint64) uint64
	Print(myKeyValueStore)
}

type Control interface {
	Create(uint64) //Create(string, uint64) for directory & memorySize
	Open(string, string)
	Close()
	Delete(string)
}

// creates a new kvStrore with a given directoryName and memorySize
type myKeyValueStore struct {
	data map[uint64][10]byte
}

type SSTable struct {
	minKey       uint64
	maxKey       uint64
	shallowIndex []uint64
	skipList     map[uint64][10]byte
}

// creates a new kvStore
func Create(memorySize int) *myKeyValueStore {
	return &myKeyValueStore{
		data: make(map[uint64][10]byte, memorySize),
	}
}

// puts in a new key
func (kvStore myKeyValueStore) Put(key uint64, value [10]byte) error {
	//puts a key in the kvStore & creates a new map if there is none
	if kvStore.data == nil {
		kvStore.data = make(map[uint64][10]byte)
	}
	kvStore.data[key] = value

	if len(kvStore.data) == 5 {
		fmt.Println("SkipList is full, creating SSTable...")
		kvStore.CreateSSTable()
	}
	//return nil if everything went well
	return nil
}

// returns an exisiting key
func (kvStore myKeyValueStore) Get(key uint64) [10]byte {
	value, keyExists := kvStore.data[key]
	//returns a key only if it exists
	if keyExists {
		return value
	} else {
		var keyNotFound [10]byte
		return keyNotFound
	}
}

func (kvStore myKeyValueStore) CreateSSTable() {
	keys := make([]uint64, 0, len(kvStore.data))
	for key := range kvStore.data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	fmt.Println("Sorted Key Value Store:")
	builder := strings.Builder{}
	for _, key := range keys {
		fmt.Fprintf(&builder, "Key: %d, Value: %s\n", key, kvStore.data[key])
	}
	var shallowIndex []uint64
	for i := 1; i < len(keys); i += 3 {
		shallowIndex = append(shallowIndex, keys[i])
	}
	ssTable := SSTable{
		minKey:       keys[0],
		maxKey:       keys[len(keys)-1],
		shallowIndex: shallowIndex,
		skipList:     kvStore.data,
	}
	fmt.Println("SSTable:")
	ssTable.Format(&builder)

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Write a string to the file
	_, err = file.WriteString(builder.String())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Data written to file successfully.")
}

func (s *SSTable) Format(builder *strings.Builder) {
	builder.WriteString("Min Key: ")
	builder.WriteString(fmt.Sprintf("%d\n", s.minKey))
	builder.WriteString("Max Key: ")
	builder.WriteString(fmt.Sprintf("%d\n", s.maxKey))
	builder.WriteString("Shallow Index: ")
	builder.WriteString(fmt.Sprintf("%v\n", s.shallowIndex))
	builder.WriteString("Sorted Map:\n")
	for key, value := range s.skipList {
		builder.WriteString(fmt.Sprintf("Key: %d, Value: %s\n", key, value))
	}
}

// accesses a datbase & establishes connection
func (kvStore myKeyValueStore) Open(driverName string, databaseName string) {
	//tbd check that the driverName & datbase exist
	//what happens if already open?
	/* documenation for implementation: https://pkg.go.dev/database/sql#Open
	db, error = sql.Open(driverName, databaseName)
	if, error := nil{
		log.Fatal(error)
	}
	*/
}

// closes the database and prevent new queries form starting
func (kvStore myKeyValueStore) Close() error {
	/* documenation for implementation: https://pkg.go.dev/database/sql#DB.Close
	error := db.Close()
	if error != nil{
		return error
	} else{
		return
	}
	*/
	//replace once implemented
	return nil
}

// deletes a kvStore
func (kvStore myKeyValueStore) Delete(directory string) {
	//checking that the connection to the db is closed
	kvStore.Close()
	//tbd check that the kvStore exists (probably not necessary?)
	//tbd implementation of deleting the kvStore
	//does it make a difference if in memory or disk?
}
