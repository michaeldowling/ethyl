package contracts

import (
    "encoding/json"
    "log"
)

type Contract struct {
    Abi     ABI
    Address string
}

type ABI struct {
    InterfaceDefinitions []ContractInterfaceDefinition
}

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

func DefineContract(abi string) (Contract, error) {

    var contract Contract;
    var abiDefinition ABI;
    contractErr := json.Unmarshal([]byte(abi), &abiDefinition.InterfaceDefinitions);
    if (contractErr != nil) {
        log.Printf("Unable to define contract:  %v", contractErr);
        return contract, contractErr;
    }

    contract.Abi = abiDefinition;
    return contract, nil;

}

func At(abi string, address string) (Contract, error) {

    c, err := DefineContract(abi);
    c.Address = address;
    return c, err;

}

func (c *Contract) New() (error) {

    return nil;

}

