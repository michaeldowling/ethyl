package ethyl

import (
    "log"
)

type NetAPI struct {
    Client *EthylClient
}

func CreateNetAPI(client *EthylClient) (NetAPI) {

    netapi := NetAPI{Client:client};
    return (netapi);

}

func (n *NetAPI) Version() (string, error) {

    var result EthereumNetworkResponse;

    n.Client.Call("net_version", "", &result);
    version := result.Result;


    log.Println("Version: " + version);
    return version, nil;


}
