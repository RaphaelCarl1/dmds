package phf

type Data interface {
	Build(keys []uint64)
	Get(key uint64)
}

type myPHF struct {
	data map[uint64][10]byte
}

func Create() myPHF {
	return myPHF{}
}

func (PHF myPHF) Build(keys []uint64) {
	//takes the input and builds the structure
	//also updates & deletes a key - maybe can be added as a recursive build later?
}

func (PHF myPHF) Get(key uint64) uint64 {
	//returns a key
	return key
}

/*I am aware that this looks very minimalistic, as the implementation I have in mind does
not require a lot of functions.*/
