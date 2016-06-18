package contracts

import (
    "encoding/json"
    "log"
    "github.com/michaeldowling/ethyl"
)

type Contract struct {

    Abi          ABI
    Address      string
    ContractCode EVMCode

    Client       ethyl.EthylClient

}

type ABI struct {
    InterfaceDefinitions []ContractInterfaceDefinition
}

type EVMCode string;

type ContractInterfaceDefinition struct {
    Constant bool `json:"constant"`
    Name     string `json:"name"`
    Type     string `json:"type"`
    Inputs   []InputOutputDefinition `json:"inputs"`
    Outputs  []InputOutputDefinition `json:"outputs"`
}

type InputOutputDefinition struct {
    Name         string `json:"name"`
    EthereumType string `json:"type"`
}

func DefineContract(client ethyl.EthylClient, abi string) (Contract, error) {

    var contract Contract;
    var abiDefinition ABI;
    contractErr := json.Unmarshal([]byte(abi), &abiDefinition.InterfaceDefinitions);
    if (contractErr != nil) {
        log.Printf("Unable to define contract:  %v", contractErr);
        return contract, contractErr;
    }

    contract.Abi = abiDefinition;
    contract.Client = client;

    return contract, nil;

}

func At(client ethyl.EthylClient, abi string, address string) (Contract, error) {

    c, err := DefineContract(client, abi);
    c.Address = address;

    return c, err;

}

func (c *Contract) New() (error) {

    return nil;

}

