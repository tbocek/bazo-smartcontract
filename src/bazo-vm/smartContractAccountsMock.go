package bazo_vm

type ContractAccount struct {
	Address            []byte
	Balance            uint64
	TxCnt              uint64
	IsStaking          bool
	HashedSeed         []byte
	StakingBlockHeight uint64
	Code               []byte         // Additional to standard account
	ContractVariables  map[int][]byte // Additional to standard account
}

func NewContractAccount(address []byte, balance uint64, isStaking bool, hashedSeed []byte, code []byte) ContractAccount {
	newSC := ContractAccount{
		address,
		balance,
		0,
		isStaking,
		hashedSeed,
		0,
		code,
		map[int][]byte{},
	}
	return newSC
}

type ContractCallersTransaction struct {
	transactionSender    []byte
	transactioninputData []byte
	maxGasAmount         int
}
