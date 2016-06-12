package ethyl

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {

    client, err := CreateClient("localhost", 8545);
    assert.Nil(t, err);
    assert.NotNil(t, client);
    assert.Equal(t, "localhost", client.Host);
    assert.Equal(t, 8545, client.Port);

}


