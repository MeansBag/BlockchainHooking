package main

type BlockchainHook interface {
	Init() error
	OnEthBlock(*EthBlock) error
	OnEthTransaction(*EthTransaction) error
	Deinit()
}

type DataPersistenceHook interface {
	BlockchainHook
	GetEthOffendBlockNumber() (uint64, error)
}
