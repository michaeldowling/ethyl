package ethyl

import (
    "log"
    "strconv"
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




