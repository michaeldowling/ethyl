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

func TestAPINetVersion(t *testing.T) {

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);

    version, err := client.Net.Version();

    assert.NotNil(t, version);
    assert.Equal(t, "69", version);
    assert.Nil(t, err);


}
