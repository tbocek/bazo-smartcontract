package main

type virtualMachine struct {
	program         []byte
	executionEngine executionEngine
}

func newVM(program []byte) virtualMachine {
	return virtualMachine{
		program:         program,
		executionEngine: newExecutionEngine(),
	}
}

func (vm virtualMachine) getProgram() []byte {
	return vm.program
}

func (vm virtualMachine) run() int {
	for _, opCode := range vm.program {
		vm.executionEngine.executeOperation(opCode)
	}

	return 0
}
