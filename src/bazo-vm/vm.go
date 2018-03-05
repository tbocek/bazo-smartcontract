package main

type virtualMachine struct {
	program []byte
	executionEngine executionEngine
}

func newVM(program []byte) virtualMachine{
	return virtualMachine{
		program: program,
		executionEngine: newExecutionEngine(),
	}
}

func (vm virtualMachine) getProgram() []byte {
	return vm.program
}
