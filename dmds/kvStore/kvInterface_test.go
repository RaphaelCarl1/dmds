package keyValueStore

import "testing"

//do tests work with zero values?
var (
	kvStore      = Create()
	key          uint64 //key cannot be 123, see TestGetNonexistent
	value        [10]byte
	driverName   string
	databaseName string
	directory    string
)

//puts and gets a key - happy path
func TestPutGet(t *testing.T) {
	kvStore.Put(key, value)
	var recievedValue [10]byte = kvStore.Get(key)
	t.Errorf("recievedKey %d is not givenKey %d", value, recievedValue)
}

//tests that put does not accept doublettes
func TestPutExisting(t *testing.T) {
	kvStore.Put(key, value)
	kvStore.Put(key, value)
	//expected error
	t.Errorf("The key %d already exists in that kvStore", key)
	//assert that the correct error is handed back as soon as implementation is done
}

//tests behaviour of put() when a kvStore does not exist (can kvStore.Put() work if kvStore does not exist?)
func TestPutDeleted(t *testing.T) {
	kvStore.Delete(directory)
	kvStore.Put(key, value)
	//expected error
	t.Errorf("The kvStore %s does not exist", directory)
	//assert that the correct error is handed back as soon as implementation is done
}

//behaviour when a get() has a nonexisting key
func TestGetNonesistent(t *testing.T) {
	var wrongKey uint64 = 123
	kvStore.Put(key, value) //(need to put a key there first?)
	value = kvStore.Get(wrongKey)
	//expected error
	t.Errorf("Key %d has not been found", wrongKey)
	//tbd test passes if value = keyNotFound [10]byte
}

//confirms the connection to db is open
func TestOpenDatabse(t *testing.T) {
	kvStore.Open(driverName, databaseName)
	/* update when implementation is done
		documentation of implementation: https://go.dev/doc/database/open-handle
	if err ;: db.Ping(); err !=nil{
		log.Fatal(err)
	}
	*/
	t.Errorf("kvStore %s was not able to access the database", directory)
}

//opens an open db
func TestOpenDatabseAgain(t *testing.T) {
	kvStore.Open(driverName, databaseName)
	kvStore.Open(driverName, databaseName)
	//should this be an error or is this unnecessary?
	t.Errorf("kvStore %s has already access to the database", directory)
}

//tries to open a nonexisting db
func TestOpenNonexisting(t *testing.T) {
	var wrongDriverName string = "WrongName"
	kvStore.Open(wrongDriverName, databaseName)
	//tbd update when implemented (what is expected behaviour?)
	//expected error
	t.Errorf("kvStore %s was not able to access the database", directory)
}

//confirms the connection to the db is closed
func TestCloseDatabase(t *testing.T) {
	kvStore.Open(driverName, databaseName)
	kvStore.Close()
	//tbd update when implementation is done
	//defer db.Closse() (accoring to documentation)
	t.Errorf("kvStore %s was not able to close the connection to the database", directory)
}

//deleting an existing kvStore
func TestDeleteKVStore(t *testing.T) {
	kvStore.Delete(directory)
	//tbd assert that the kvStore does not exist
	t.Errorf("kvStore %s was not deleted", directory)
}

//deleting a nonexisting kvStore
func TestDeleteNonexisting(t *testing.T) {
	//delete the kvStore successfully
	kvStore.Delete(directory)
	//delete the kvStore again
	kvStore.Delete(directory)
	//expected error
	t.Errorf("kvStore %s does not exist", directory)
}
