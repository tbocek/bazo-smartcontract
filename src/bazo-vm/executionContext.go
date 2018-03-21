package bazo_vm

type StateData struct {
	data []byte
}

type Context struct {
	sender       []byte
	maxGasPrice  int
	inputData    []byte
	maxGasAmount int
	/*
	owner []byte
	stateData StateData

	smartContract SmartContract
	blockHeader []byte*/
}
