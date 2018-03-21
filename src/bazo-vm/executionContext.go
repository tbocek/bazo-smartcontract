package bazo_vm

type StateData struct {
	data []byte
}

type Context struct {
	transactionSender       []byte
	transactioninputData    []byte
	maxGasAmount int
	smartContract SmartContract
	/*
	stateData StateData


	blockHeader []byte*/
}
