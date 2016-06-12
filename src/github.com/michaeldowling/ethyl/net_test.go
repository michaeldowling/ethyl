package ethyl

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "log"
    "strconv"
)

func TestAPINetVersion(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    version, err := client.Net.Version();

    assert.NotNil(t, version);
    assert.Equal(t, "100", version);
    assert.Nil(t, err);

}

func TestAPINetIsListening(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    isListening, err := client.Net.IsListening();
    log.Println("Is lIstening:  " + strconv.FormatBool(isListening));

    assert.NotNil(t, isListening);
    assert.True(t, isListening);
    assert.Nil(t, err);


}

func TestAPINetPeerCount(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    peerCount, _ := client.Net.PeerCount();
    log.Printf("Peer Count:  %v", peerCount);

    assert.True(t, peerCount >= 0);

}
