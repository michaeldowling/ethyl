package ethyl

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAPIEthProtocolVersion(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    protocolVersion, _ := client.Eth.ProtocolVersion();
    assert.True(t, protocolVersion > 0, "Protocol Version should be greater than 0");


}


