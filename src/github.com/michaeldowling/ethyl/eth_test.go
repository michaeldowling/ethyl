package ethyl

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "log"
)

func TestAPIEthProtocolVersion(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    protocolVersion, _ := client.Eth.ProtocolVersion();
    assert.True(t, protocolVersion > 0, "Protocol Version should be greater than 0");

}

func TestAPIEthGasPrice(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    gasPrice, _ := client.Eth.GasPrice();
    assert.True(t, gasPrice > 0);

    log.Printf("Gas price:  %v", gasPrice);

}

func TestAPIEthSyncing(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    isSyncing, _, _, _, err := client.Eth.Syncing();
    assert.Nil(t, err);
    assert.NotNil(t, isSyncing);

}

