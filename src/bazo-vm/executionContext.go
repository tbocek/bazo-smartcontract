package bazo_vm

type StateData struct {
	data []byte
}

type Context struct {
	transactionSender    []byte
	transactioninputData []byte
	maxGasAmount         int
	contractAccount      ContractAccount
	/*
		stateData StateData


		blockHeader []byte*/
}

func NewContext() *Context {
	//data := map[int][]byte{}

	return &Context{
		transactionSender:    []byte{},
		transactioninputData: []byte{},
		maxGasAmount:         100000,
		contractAccount:      ContractAccount{},
	}
}
