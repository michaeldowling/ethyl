/*
 * (C) Copyright 2016 Michael Dowling (http://github.com/michaeldowling) and others.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package contracts

import (
    "testing"
    "io/ioutil"
    "github.com/stretchr/testify/assert"
    "encoding/json"
    ethcrypto "github.com/ethereum/go-ethereum/crypto"
    "encoding/hex"
)

func TestEncodeFunctionSelector(t *testing.T) {

    // Load ABI
    contents, fileErr := ioutil.ReadFile("../test_contracts/foo_test_abi.json");
    // contractBytecode := "606060405261018f806100126000396000f360606040526000357c010000000000000000000000000000000000000000000000000000000090048063a5643bf21461004f578063ab55044d146100f2578063cdcd77c01461012f5761004d565b005b6100f06004808035906020019082018035906020019191908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509090919080359060200190919080359060200190820180359060200191919080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050909091905050610166565b005b61012d60048080604001906002806020026040519081016040528092919082600260200280828437820191505050505090909190505061016c565b005b61014e6004808035906020019091908035906020019091905050610170565b60405180821515815260200191505060405180910390f35b5b505050565b5b50565b600060208363ffffffff1611806101845750815b905080505b9291505056";

    assert.Nil(t, fileErr);
    assert.NotNil(t, contents);

    var abiDefinition ABI;
    contractErr := json.Unmarshal([]byte(contents), &abiDefinition.InterfaceDefinitions);
    assert.Nil(t, contractErr);

    // get hash of function name
    hash := ethcrypto.Keccak256([]byte("baz(uint32,bool)"));
    assert.NotNil(t, hash);
    assert.NotEmpty(t, hash);

    firstFour := hash[0:4];
    assert.Equal(t, "0xcdcd77c0", "0x" + hex.EncodeToString(firstFour));




}
