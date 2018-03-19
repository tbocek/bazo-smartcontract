package bazo_vm

//This is a struct because the contract can be of a
//different size.
type ContractCode struct {
	code []byte
}

type SmartContract struct {
	Address            [64]byte
	Balance            uint64
	TxCnt              uint32
	IsStaking          bool
	HashedSeed         [32]byte
	StakingBlockHeight uint32
	data               ContractCode
}

func NewSmartContract(address [64]byte, balance uint64, isStaking bool, hashedSeed [32]byte, code []byte) SmartContract {
	newSC := SmartContract{
		address,
		balance,
		0,
		isStaking,
		hashedSeed,
		0,
		ContractCode{code: code,},
	}
	return newSC
}
