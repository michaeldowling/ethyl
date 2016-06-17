package ethyl

import (
    "log"
    "strconv"
    "github.com/michaeldowling/ethyl/util"
    "errors"
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
};

type TransactionInstructions struct {
    From  string `json:"from"`
    To    string `json:"to"`
    Gas   int64 `json:"gas"`
    Data  string `json:"data"`
    Value int64 `json:"value"`

}


type EthAPI struct {
    Client *EthylClient
}

func CreateEthAPI(client *EthylClient) (EthAPI) {

    ethapi := EthAPI{Client:client};
    return (ethapi);

}

func (e *EthAPI) ProtocolVersion() (uint64, error) {

    var result StringResultEthereumNetworkResponse;

    e.Client.Call("eth_protocolVersion", nil, &result);
    versionString := result.Result;

    log.Println("Version:  " + versionString);

    version, err := strconv.ParseUint(versionString, 10, 64);
    return version, err;

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

    return result.Result, nil;

}



