package contracts

import (
    "encoding/json"
    "log"
    "github.com/michaeldowling/ethyl"
    "time"
)

type DeployConfig struct {
    Gas      int
    GasPrice int
    Value    int
}

type DeployResults struct {
    State           string
    TransactionHash string
    ContractAddress string
}

type Contract struct {

    Abi          ABI
    Address      string
    ContractCode EVMCode


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

func DefineContract(abi string, evmCode string) (Contract, error) {

    var contract Contract;
    var abiDefinition ABI;
    contractErr := json.Unmarshal([]byte(abi), &abiDefinition.InterfaceDefinitions);
    if (contractErr != nil) {
        log.Printf("Unable to define contract:  %v", contractErr);
        return contract, contractErr;
    }

    contract.Abi = abiDefinition;
    contract.ContractCode = EVMCode(evmCode);

    return contract, nil;

}

func At(client ethyl.EthylClient, abi string, address string) (Contract, error) {

    c, err := DefineContract(abi, "");
    c.Address = address;

    return c, err;

}

func (c *Contract) Deploy(client ethyl.EthylClient, config DeployConfig) (DeployResults, error) {

    instr := ethyl.TransactionInstructions{From:client.Accounts[0], Gas:int64(config.Gas), Value:int64(config.Value), Data:string(c.ContractCode)};
    txHash, err := client.Eth.SendTransaction(instr);

    if (err != nil) {
        log.Printf("Error while sending transaction to create contract:  %v", err);
        return DeployResults{State:"error"}, err;
    }

    deployResults := DeployResults{State:"pending", TransactionHash:txHash};
    return deployResults, nil;

}

func (c *Contract) MonitorDeploy(client ethyl.EthylClient, deployResults DeployResults) (chan string) {

    deploymentChannel := make(chan string);
    go func() {

        for {

            tx, err := client.Eth.GetTransactionByHash(deployResults.TransactionHash);
            if (err != nil) {
                deploymentChannel <- "";
                break;
            }
            if (tx.BlockNumber != 0) {
                deploymentChannel <- tx.Hash; // TODO - need the contract address and use getTxReceipt :(
                break
            }

            time.Sleep(1 * time.Second);

        }

    }();

    return deploymentChannel;

}

