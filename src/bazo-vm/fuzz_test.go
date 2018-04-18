package bazo_vm
/*
import (
	"log"
	"testing"
)


func Fuzz() {
	vm := NewVM()
	code := RandomBytes()
	vm.context.MaxGasAmount = 10000
	vm.context.ContractAccount.Contract = code

	defer func() {
		if err := recover(); err != nil {
			log.Println("Execution failed", err, code)
		}
	}()

	vm.Exec(false)
}

func TestFuzz(t *testing.T) {
	for i := 0; i <= 1000000; i++ {
		Fuzz()
	}
}
*/