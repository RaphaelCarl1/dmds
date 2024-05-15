package phf

import "testing"

var (
	phfTable myPHF
	keys     = []uint64{1, 2, 3, 4}
)

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
	key := uint64(5)
	result := phfTable.Get(key)

	if result != 0 {
		t.Errorf("Expected result should be error, but got %d", result)
	}
}
