package util

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEthereumHash(t *testing.T) {

    result, err := EthereumHash([]byte("baz(uint32,bool)"));
    assert.Nil(t, err);
    assert.NotNil(t, result);

    assert.Equal(t, "0xcdcd77c0992ec5bbfc459984220f8c45084cc24d9b6efed1fae540db8de801d2", result);

}
