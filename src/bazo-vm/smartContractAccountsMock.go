package bazo_vm

//This is a struct because the contract can be of a
//different size.
type ContractCode struct {
	code []byte
}

type SmartContract struct {
	Address            []byte
	Balance            uint64
	TxCnt              uint64
	IsStaking          bool
	HashedSeed         []byte
	StakingBlockHeight uint64
	data               ContractCode
	contractVariables  map[int][]byte
}

func NewSmartContract(address []byte, balance uint64, isStaking bool, hashedSeed []byte, code []byte, data map[int][]byte) SmartContract {
	newSC := SmartContract{
		address,
		balance,
		0,
		isStaking,
		hashedSeed,
		0,
		ContractCode{code: code,},
		map[int][]byte{},
	}
	return newSC
}

type ContractCallersTransaction struct {
	transactionSender       []byte
	transactioninputData    []byte
	maxGasAmount int
}