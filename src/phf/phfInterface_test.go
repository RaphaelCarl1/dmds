package phf

import "testing"

var (
	phfTable MyPHF
	keys     = make([]uint64, 5)
)

func TestCustomHashKey(t *testing.T) {
    hashes := make(map[uint64]uint64)

    // Test case 1: Check that the hash of a key is within the expected range
    key1 := uint64(1234567890)
    hash1 := customHashKey(key1, hashes)
    if hash1 < 1 || hash1 >= 10 {
        t.Errorf("Hash of key1 is out of range: got %v, want a value between 0 and 9", hash1)
    }

    // Test case 2: Check that the hash of a different key is not the same
    key2 := uint64(9876543210)	
    hash2 := customHashKey(key2, hashes)
    if hash2 == hash1 {
        t.Errorf("Hash of key2 is the same as hash of key1: got %v, want a different value", hash2)
    }

    // Test case 3: Check that the hash of the same key is the same
    hash3 := customHashKey(key1, hashes)
    if hash3 != hash1 {
        t.Errorf("Hash of the same key is not the same: got %v, want %v", hash3, hash1)
    }
}

func TestBuildNull(t *testing.T) {
	keys := []uint64{}
	phfTable.Build(keys)

	if len(phfTable.data) != 0 {
		t.Errorf("PHF should be empty, and not %v", phfTable)
	}

}

func TestBuildValidKeys(t *testing.T) {
	phfTable.Build(keys)
	key := uint64(2)
	result := phfTable.Get(key)

	if result != key {
		t.Errorf("Expected result should be %d, and not %d", key, result)
	}
	// Assert that the resulting phTable is constructed correctly
}

func TestBuildMultipleTimes(t *testing.T) {
	phfTable.Build(keys)
	phfTable.Build(keys)
	key := uint64(2)
	result := phfTable.Get(key)

	if result != key {
		t.Errorf("Expected result should be %d, and not %d", key, result)
	}
	// Assert that the resulting phTable is constructed correctly
}

func TestGetExistingKey(t *testing.T) {
	phfTable.Build(keys)
	key := uint64(2)
	result := phfTable.Get(key)

	if result != key {
		t.Errorf("Expected result should be %d, and not %d", key, result)
	}
}

func TestGetNullKey(t *testing.T) {
	phfTable.Build(keys)
	key := uint64(0)
	result := phfTable.Get(key)

	if result != key {
		t.Errorf("Expected result should be %d, and not %d", key, result)
	}
}

func TestGetNonExistentKey(t *testing.T) {
	phfTable.Build(keys)
	key := uint64(6)
	result := phfTable.Get(key)

	if result != 0 {
		t.Errorf("Expected result should be error, but got %d", result)
	}
}
