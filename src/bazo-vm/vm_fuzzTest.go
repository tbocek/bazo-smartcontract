package bazo_vm

func Fuzz(data []byte) int {

	vm := NewVM()
	vm.context.contractAccount.Code = data
	vm.Exec(true)

	_, err := vm.evaluationStack.Peek()

	if err != nil {
		return 0
	}

	return 1
}
