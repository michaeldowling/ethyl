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

func TestEthAPI_Accounts(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    accounts, err := client.Eth.Accounts();
    assert.Nil(t, err);
    assert.NotNil(t, accounts);
    assert.True(t, len(accounts) > 0);

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

    // get account first
    assert.True(t, len(client.Accounts) > 0);

    instructions := TransactionInstructions{From:client.Accounts[0], To:"0xabd5d148b31f38a8d2aa9eb041c478d36dd51c35", Gas:200000, Value:10};
    txhash, err := client.Eth.SendTransaction(instructions);

    assert.Nil(t, err);
    assert.NotEmpty(t, txhash);


}

func TestEthAPI_NewFilter(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    options := FilterOptions{FromBlock:"latest"};
    filterId, filterErr := client.Eth.NewFilter(options);

    assert.Nil(t, filterErr);
    assert.NotNil(t, filterId);
    assert.NotEmpty(t, filterId);

    log.Println("Filter ID:  " + filterId);


}


