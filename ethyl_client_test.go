package ethyl

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)


func TestCreateClient(t *testing.T) {

    testUrl := os.Getenv("TEST_RPC_URL")

    client, err := CreateClient(testUrl);
    assert.Nil(t, err);
    assert.NotNil(t, client);
    assert.Equal(t, testUrl, client.Url);

}


