package keyValueStore

//import _ "github.com/go-sql-driver/mysql" - for db-stuff

type Data interface {
	Put(uint64, [10]byte)
	Get(uint64) uint64
}

type Control interface {
	Create(uint64) //Create(string, uint64) for directory & memorySize
	Open(string, string)
	Close()
	Delete(string)
}

//creates a new kvStrore with a given directoryName and memorySize
type myKeyValueStore struct {
	data map[uint64][10]byte
	//directory string
	//memorySize int
}

//creates a new kvStore
func Create(memorySize uint64) myKeyValueStore {
	return myKeyValueStore{
		data: make(map[uint64][10]byte),
	}
}

//puts in a new key
func (kvStore myKeyValueStore) Put(key uint64, value [10]byte) error {
	//tbd check if key already exists, then return error
	//tbd check if the kvStore exists (probably not necessary?)
	//puts a key in the kvStore
	kvStore.data[key] = value
	//return nil if everything went well
	return nil
}

//returns an exisiting key
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

//accesses a datbase & establishes connection
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

//closes the database and prevent new queries form starting
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

//deletes a kvStore
func (kvStore myKeyValueStore) Delete(directory string) {
	//checking that the connection to the db is closed
	kvStore.Close()
	//tbd check that the kvStore exists (probably not necessary?)
	//tbd implementation of deleting the kvStore
	//does it make a difference if in memory or disk?
}
