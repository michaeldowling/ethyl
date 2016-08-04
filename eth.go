/*
 * (C) Copyright ${year} Michael Dowling (http://github.com/michaeldowling) and others.
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

package ethyl

import (
    "log"
    "github.com/michaeldowling/ethyl/util"
    "errors"
    "time"
)

type Transaction struct {
    Hash             string
    Nonce            string
    BlockHash        string
    BlockNumber      int64
    TransactionIndex int64
    From             string
    To               string
    Value            string
    Gas              int64
    GasPrice         int64
    Input            string
}

type TransactionInstructions struct {
    From  string `json:"from"`
    To    string `json:"to,omitempty"`
    Gas   int64 `json:"gas,omitempty"`
    Data  string `json:"data,omitempty"`
    Value int64 `json:"value,omitempty"`
}

type TransactionReceipt struct {
    TransactionHash   string `json:"transactionHash"`
    TransactionIndex  int64 `json:"transactionIndex"`
    BlockHash         string `json:"blockHash"`
    BlockNumber       int64 `json:"blockNumber"`
    CumulativeGasUsed int64 `json:"cumulativeGasUsed"`
    GasUsed           int64 `json:"gasUsed"`
    ContractAddress   string `json:"contractAddress"`
    Error             string
    // Logs []string - TODO
}

type LogEntry struct {
    Type             string `json:"type"`
    LogIndex         int64 `json:"logIndex"`
    TransactionIndex int64 `json:"transactionIndex"`
    TransactionHash  string `json:"transactionHash"`
    BlockHash        string `json:"blockHash"`
    BlockNumber      int64 `json:"blockNumber"`
    Address          string `json:"address"`
    Data             string `json:"data"`
    Topics           []string `json:"topics"`

}


type EthAPI struct {
    Client *EthylClient
}

func CreateEthAPI(client *EthylClient) (EthAPI) {

    ethapi := EthAPI{Client:client};
    return (ethapi);

}

func (e *EthAPI) ProtocolVersion() (int64, error) {

    var result StringResultEthereumNetworkResponse;

    e.Client.Call("eth_protocolVersion", nil, &result);
    versionString := result.Result;
    version, err := util.ToInt64(versionString);

    if (err != nil) {
        return 0, err;
    }

    log.Printf("Version:  %d \n", version);
    return version, nil;

}

func (e *EthAPI) Accounts() ([]string, error) {

    var result GenericSliceResultEthereumNetworkResponse;

    e.Client.Call("eth_accounts", nil, &result);
    accounts := make([]string, len(result.Result));
    for idx, item := range result.Result {
        accounts[idx] = item.(string);
    }

    return accounts, nil;

}

func (e *EthAPI) GasPrice() (int64, error) {

    var result StringResultEthereumNetworkResponse;

    e.Client.Call("eth_gasPrice", nil, &result);
    gasPriceHex := result.Result;

    gasPrice, err := util.ToInt64(gasPriceHex);
    if (err != nil) {
        log.Println("Error parsing Gas Price:  " + err.Error());
        return 0, err;
    }

    return gasPrice, err;

}

func (e *EthAPI) Syncing() (isSyncing bool, startingBlock string, currentBlock string, highestBlock string, err error) {

    // var result StringResultEthereumNetworkResponse;
    var result map[string]interface{};
    e.Client.Call("eth_syncing", nil, &result);

    if (result["result"] == false) {
        return false, "", "", "", nil;
    }

    return false, "", "", "", nil;

}

func (e *EthAPI) GetTransactionByHash(transactionHash string) (Transaction, error) {

    var result TransactionObjectResultEthereumNetworkResponse;
    txid := []string{transactionHash};

    e.Client.Call("eth_getTransactionByHash", txid, &result);
    tx := result.Result;

    // Do some conversions
    blockNumber, _ := util.ToInt64(tx.BlockNumber);
    transactionIndex, _ := util.ToInt64(tx.TransactionIndex);
    gas, _ := util.ToInt64(tx.Gas);
    gasPrice, _ := util.ToInt64(tx.GasPrice);

    transaction := Transaction{
        Hash:tx.Hash,
        Nonce:tx.Nonce,
        BlockHash:tx.BlockHash,
        BlockNumber:blockNumber,
        TransactionIndex:transactionIndex,
        From:tx.From,
        To:tx.To,
        Value:tx.Value,
        Gas:gas,
        GasPrice:gasPrice,
        Input:tx.Input,
    };

    return transaction, nil;

}

func (e *EthAPI) GetTransactionReceipt(transactionHash string) (TransactionReceipt, error) {

    var result TransactionReceiptObjectResultEthereumNetworkResponse;
    txid := []string{transactionHash};

    e.Client.Call("eth_getTransactionReceipt", txid, &result);
    tx := result.Result;

    // Do some conversions
    transactionIndex, _ := util.ToInt64(tx.TransactionIndex);
    blockNumber, _ := util.ToInt64(tx.BlockNumber);
    cumulativeGasUsed, _ := util.ToInt64(tx.CumulativeGasUsed);
    gasUsed, _ := util.ToInt64(tx.GasUsed);

    transactionReceipt := TransactionReceipt{
        TransactionHash:tx.TransactionHash,
        TransactionIndex:transactionIndex,
        BlockHash:tx.BlockHash,
        BlockNumber:blockNumber,
        CumulativeGasUsed:cumulativeGasUsed,
        GasUsed:gasUsed,
        ContractAddress:tx.ContractAddress,
    };

    return transactionReceipt, nil;


}

func (e *EthAPI) SendTransaction(instructions TransactionInstructions, transactionArguments ... interface{}) (string, error) {

    var result StringResultEthereumNetworkResponse;

    jsonArgs := make([]string, len(transactionArguments));
    for idx, argument := range transactionArguments {
        jsonArgs[idx] = argument.(string);
    }

    txErr := e.Client.CallWithTransaction("eth_sendTransaction", instructions, jsonArgs, &result);

    if (txErr != nil) {
        return "", errors.New("Unable to send transaction to network:  " + txErr.Error());
    }

    if (result.Error.Code != 0) {
        return "", errors.New("Unable to send transaction to network:  " + result.Error.Message);
    }

    return result.Result, nil;

}

func (e *EthAPI) GetTransactionMonitor(txhash string) (chan TransactionReceipt, error) {

    deploymentChannel := make(chan TransactionReceipt);
    go func() {

        for {

            tx, err := e.GetTransactionReceipt(txhash);
            if (err != nil) {
                deploymentChannel <- TransactionReceipt{Error:err.Error()};
                break;
            }
            if (tx.BlockNumber != 0) {
                deploymentChannel <- tx;
                break
            }

            time.Sleep(1 * time.Second);

        }

    }();

    return deploymentChannel, nil;


}

func (e *EthAPI) NewFilter(options FilterOptions) (string, error) {

    var result StringResultEthereumNetworkResponse;
    txErr := e.Client.CallWithFilterOptions("eth_newFilter", options, &result);

    if (txErr != nil) {
        log.Printf("Problem with calling with filter options:  %s \n", txErr);
        return "", txErr;
    }

    if (result.Error.Code != 0) {
        return "", errors.New("Unable to create a filter with options given:  " + result.Error.Message);
    }

    return result.Result, nil;


}

func (e *EthAPI) GetCode(address string) (string, error) {

    var result StringResultEthereumNetworkResponse;
    options := []string{address, "latest"};

    err := e.Client.Call("eth_getCode", options, &result);
    if(err != nil) {
        return "", err;
    }

    if(result.Result == "") {
        return "", errors.New("Code field was empty");
    }

    // remove the Ox.  I mean, really now.


    return result.Result, nil;


}

func (e *EthAPI) GetFilterChanges(filterId string) ([]LogEntry, error) {

    var result FilterLogObjectResultEthereumNetworkResponse;
    options := []string{filterId};

    err := e.Client.Call("eth_getFilterChanges", options, &result);
    if(err != nil) {
        return nil, err;
    }

    entries := result.Result;
    logs := make([]LogEntry, len(entries));
    for idx, entry := range entries {

        logIndex, _ := util.ToInt64(entry.LogIndex);
        txIndex, _ := util.ToInt64(entry.TransactionIndex);
        blockNumber, _ := util.ToInt64(entry.BlockNumber);

        logEntry := LogEntry{
            Type: entry.Type,
            LogIndex:logIndex,
            TransactionIndex:txIndex,
            TransactionHash:entry.TransactionHash,
            BlockHash:entry.BlockHash,
            BlockNumber:blockNumber,
            Address:entry.Address,
            Data:entry.Data,
            Topics:entry.Topics,
        };

        logs[idx] = logEntry;

    }

    return logs, nil;

}

