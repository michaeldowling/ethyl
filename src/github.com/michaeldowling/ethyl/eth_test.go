package ethyl

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "log"
)

func TestAPI_EthProtocolVersion(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    protocolVersion, _ := client.Eth.ProtocolVersion();
    assert.True(t, protocolVersion > 0, "Protocol Version should be greater than 0");

}

func TestAPI_EthGasPrice(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    gasPrice, _ := client.Eth.GasPrice();
    assert.True(t, gasPrice > 0);

    log.Printf("Gas price:  %v", gasPrice);

}

func TestAPI_EthSyncing(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    isSyncing, _, _, _, err := client.Eth.Syncing();
    assert.Nil(t, err);
    assert.NotNil(t, isSyncing);

}

func TestEthAPI_GetTransactionByHash(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    tx, err := client.Eth.GetTransactionByHash("0xe564bd605dcf1e70646e6e0a13294199cdb6f026fd74a11357f07d9863c8b989");
    assert.Nil(t, err);
    assert.NotNil(t, tx);


}

func TestEthAPI_SendTransaction_EtherTransfer(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    instructions := TransactionInstructions{From:"0x4cb40800a4965f544f8edb9b17b6859773b5908e", To:"0xabd5d148b31f38a8d2aa9eb041c478d36dd51c35", Gas:200000, Value:10000};
    txhash, err := client.Eth.SendTransaction(instructions);

    assert.Nil(t, err);
    assert.NotEmpty(t, txhash);


}

