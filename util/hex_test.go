package util

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConvertHexToExpectedInteger(t *testing.T) {

    hex := "0x7DE";
    num, err := ToInt64(hex);

    assert.Nil(t, err);
    assert.Equal(t, int64(2014), num);


}
