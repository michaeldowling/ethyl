package ethyl

import (
    "log"
    "strconv"
    "github.com/michaeldowling/ethyl/util"
)

type EthAPI struct {
    Client *EthylClient
}

func CreateEthAPI(client *EthylClient) (EthAPI) {

    ethapi := EthAPI{Client:client};
    return (ethapi);

}

func (e *EthAPI) ProtocolVersion() (uint64, error) {

    var result StringResultEthereumNetworkResponse;

    e.Client.Call("eth_protocolVersion", "", &result);
    versionString := result.Result;

    log.Println("Version:  " + versionString);

    version, err := strconv.ParseUint(versionString, 10, 64);
    return version, err;

}

func (e *EthAPI) GasPrice() (int64, error) {

    var result StringResultEthereumNetworkResponse;

    e.Client.Call("eth_gasPrice", "", &result);
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
    e.Client.Call("eth_syncing", "", &result);

    if (result["result"] == false) {
        return false, "", "", "", nil;
    }

    return false, "", "", "", nil;

}



