package bazo_vm

type StateData struct {
	data []byte
}

type Context struct {
	sender []byte
	maxGasPrice int
	inputData []byte

	amount int

	owner []byte
	stateData StateData

	smartContract SmartContract
	blockHeader []byte
}



//create new Context Obj before executing run.
func run(context Context){
	vm := NewVM(0)
	//vm.Exec(context.smartContract.data, inputData, maxGasPrice)





}