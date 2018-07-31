package main

type configurations struct {
	Eth struct {
		URL      string `json:"url"`
		Interval int    `json:"interval"`
	} `json:"eth"`

	Btc struct {
		URL      string `json:"url"`
		Interval int    `json:"interval"`
	} `json:"btc"`

	DataPersistence string `json:"dataPersistence"`

	Mongo struct {
		URL                        string `json:"url"`
		Database                   string `json:"database"`
		CollectionOfEthBlock       string `json:"collectionOfEthBlock"`
		CollectionOfEthTransaction string `json:"collectionOfEthTransaction"`
		CollectionOfBtcBlock       string `json:"collectionOfBtcBlock"`
		CollectionOfBtcTransaction string `json:"CollectionOfBtcTransaction"`
	} `json:"mongo"`
}

var theConfigurations *configurations

func getDefaultConfigurations() (c *configurations) {
	c = new(configurations)

	c.Eth.URL = "http://10.0.11.82:8545"
	c.Eth.Interval = 1

	c.Btc.URL = ""
	c.Btc.Interval = 1

	c.DataPersistence = "mongo"
	c.Mongo.URL = "10.0.11.82:27017"
	c.Mongo.Database = "BlockchainHooking"
	c.Mongo.CollectionOfEthBlock = "EthBlock"
	c.Mongo.CollectionOfEthTransaction = "EthTransaction"
	c.Mongo.CollectionOfBtcBlock = "BtcBlock"
	c.Mongo.CollectionOfBtcTransaction = "BtcTransaction"

	return
}
