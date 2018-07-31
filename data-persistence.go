package main

import (
	logger "log"
	"gopkg.in/mgo.v2"
)

type mongoHook struct {
	session                    *mgo.Session
	database                   *mgo.Database
	collectionOfEthBlock       *mgo.Collection
	collectionOfEthTransaction *mgo.Collection
	collectionOfBtcBlock       *mgo.Collection
	collectionOfBtcTransaction *mgo.Collection
}

func (ptr *mongoHook) Init() (err error) {
	logger.Println("Initializing mongoHook...")

	ptr.session, err = mgo.Dial(theConfigurations.Mongo.URL)
	if err != nil {
		return
	}

	ptr.database = ptr.session.DB(theConfigurations.Mongo.Database)
	ptr.collectionOfEthBlock = ptr.database.C(theConfigurations.Mongo.CollectionOfEthBlock)
	ptr.collectionOfEthTransaction = ptr.database.C(theConfigurations.Mongo.CollectionOfEthTransaction)
	ptr.collectionOfBtcBlock = ptr.database.C(theConfigurations.Mongo.CollectionOfBtcBlock)
	ptr.collectionOfBtcTransaction = ptr.database.C(theConfigurations.Mongo.CollectionOfBtcTransaction)

	return
}

func (ptr *mongoHook) OnEthBlock(block *EthBlock) (err error) {
	_, err = ptr.collectionOfEthBlock.UpsertId(block.Number, block)
	return
}
func (ptr *mongoHook) OnEthTransaction(transaction *EthTransaction) (err error) {
	_, err = ptr.collectionOfEthTransaction.UpsertId(transaction.Hash, transaction)
	return
}

func (ptr *mongoHook) Deinit() {
	logger.Println("Deinitializing mongoHook...")
	ptr.session.Close()
}

func (ptr *mongoHook) GetEthOffendBlockNumber() (result uint64, err error) {
	count, err := ptr.collectionOfEthBlock.Find(nil).Count()
	if err != nil {
		return
	}
	if count == 0 {
		return 0, nil
	}
	return uint64(count - 1), nil
}
