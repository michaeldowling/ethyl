package contracts

import (
    "testing"
    "io/ioutil"
    "github.com/stretchr/testify/assert"
)

func TestContract_DefineContract(t *testing.T) {

    // Load ABI
    contents, fileErr := ioutil.ReadFile("../../../../../test_contracts/ballot_test_abi.json");
    assert.Nil(t, fileErr);
    assert.NotNil(t, contents);

    contract, contractErr := DefineContract(string(contents));
    assert.Nil(t, contractErr);
    assert.NotNil(t, contract);

    assert.Equal(t, "delegate", contract.Abi.InterfaceDefinitions[0].Name);


}


